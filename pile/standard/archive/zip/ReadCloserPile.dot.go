// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// ReadCloserPile is a structure for
// a lazily populated sequence (= slice)
// of items (of type `zip.ReadCloser`)
// which are cached in a growing-only list.
// Next() traverses the ReadCloserPile.
// Reset() allows a new transversal from the beginning.
//
// Yoe may either
// traverse the ReadCloserPile lazily -following its (buffered) growth that is-
// or
// await the signal from Wait() before starting traversal.
//
// Note: Pile() may be used concurrently,
// Next() (and Reset) should be confinded to a single routine (thread),
// as the iteration is not concurrency safe.
type ReadCloserPile struct {
	pile   chan zip.ReadCloser // channel to receive further items
	list   []zip.ReadCloser    // list of known items
	offset int                 // index for Next()
}

// NewS returns a (pointer to a) fresh ReadCloserPile
// of items (of type `zip.ReadCloser`)
// with size as initial capacity
// and
// with buff non-blocking Add's before respective Next's
func ReadCloserNew(size, buff int) *ReadCloserPile {
	pile := new(ReadCloserPile)
	pile.list = make([]zip.ReadCloser, 0, size)
	pile.pile = make(chan zip.ReadCloser, buff)
	return pile
}

// Reset puts the pile iterator `Next()` back at the beginning.
func (d *ReadCloserPile) Reset() {
	d.offset = 0
}

// Next returns the next item,
// or false iff the pile is exhausted.
// Next may block, awaiting another Pile(),
// iff the ReadCloserPile is not Closed().
func (d *ReadCloserPile) Next() (item zip.ReadCloser, ok bool) {
	if d.offset < len(d.list) {
		ok = true
		item = d.list[d.offset]
		d.offset++
	} else if item, ok = <-d.pile; ok {
		d.list = append(d.list, item)
		d.offset++
	}
	return item, ok
}

// Pile adds
// an item (of type `zip.ReadCloser`)
// to the ReadCloserPile.
//
// Note: Pile() may block, iff buff is exceeded and no corresponding Next()'s were called.
func (d *ReadCloserPile) Pile(item zip.ReadCloser) {
	d.pile <- item
}

// Close - call once after everything has been piled.
//
// Note: After Close(),
// any Close(...) will panic
// and
// any Pile(...) will panic
// and
// any Next() will return immediately: no eventual blocking, that is.
func (d *ReadCloserPile) Close() {
	close(d.pile)
}

// Wait returns a done channel which emits the size (=length) of the ReadCloserPile once it's been closed.
// Users of Wait() must not iterate (via Next() and Reset()) before the returned done-channel is closed!
//
// Wait is a convenience - useful iff You do not like/need to start any traversal before the ReadCloserPile is fully populated.
// (Or iff You just like to know the size to be traversed, i.e. in order to allocate some traversal-related structure.)
//
// Note: Upon close of the done-channel, the ReadCloserPile is Reset() so You may start traversing it (via Next) right away.
func (d *ReadCloserPile) Wait() (done <-chan int) {
	cha := make(chan int)
	go func(cha chan<- int, d *ReadCloserPile) {
		defer close(cha)
		d.Reset()
		defer d.Reset()
		for {
			_, ok := d.Next() // keep draining
			if !ok {
				break
			}
		}
		cha <- len(d.list) // signal the lenght, and terminate
	}(cha, d)
	return cha
}
