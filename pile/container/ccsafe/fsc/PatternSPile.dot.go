// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// PatternSPile is a structure for
// a lazily populated sequence (= slice)
// of items (of type `fs.PatternS`)
// which are cached in a growing-only list.
// Next() traverses the PatternSPile.
// Reset() allows a new transversal from the beginning.
//
// Yoe may either
// traverse the PatternSPile lazily -following its (buffered) growth that is-
// or
// await the signal from Wait() before starting traversal.
//
// Note: Pile() may be used concurrently,
// Next() (and Reset) should be confinded to a single routine (thread),
// as the iteration is not concurrency safe.
type PatternSPile struct {
	pile   chan fs.PatternS // channel to receive further items
	list   []fs.PatternS    // list of known items
	offset int              // index for Next()
}

// NewS returns a (pointer to a) fresh PatternSPile
// of items (of type `fs.PatternS`)
// with size as initial capacity
// and
// with buff non-blocking Add's before respective Next's
func PatternSNew(size, buff int) *PatternSPile {
	pile := new(PatternSPile)
	pile.list = make([]fs.PatternS, 0, size)
	pile.pile = make(chan fs.PatternS, buff)
	return pile
}

// Reset puts the pile iterator `Next()` back at the beginning.
func (d *PatternSPile) Reset() {
	d.offset = 0
}

// Next returns the next item,
// or false iff the pile is exhausted.
// Next may block, awaiting another Pile(),
// iff the PatternSPile is not Closed().
func (d *PatternSPile) Next() (item fs.PatternS, ok bool) {
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
// an item (of type `fs.PatternS`)
// to the PatternSPile.
//
// Note: Pile() may block, iff buff is exceeded and no corresponding Next()'s were called.
func (d *PatternSPile) Pile(item fs.PatternS) {
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
func (d *PatternSPile) Close() {
	close(d.pile)
}

// Wait returns a done channel which emits the size (=length) of the PatternSPile once it's been closed.
// Users of Wait() must not iterate (via Next() and Reset()) before the returned done-channel is closed!
//
// Wait is a convenience - useful iff You do not like/need to start any traversal before the PatternSPile is fully populated.
// (Or iff You just like to know the size to be traversed, i.e. in order to allocate some traversal-related structure.)
//
// Note: Upon close of the done-channel, the PatternSPile is Reset() so You may start traversing it (via Next) right away.
func (d *PatternSPile) Wait() (done <-chan int) {
	cha := make(chan int)
	go func(cha chan<- int, d *PatternSPile) {
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
