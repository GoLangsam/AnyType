// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// Int64Pile is a hybrid container for
// a lazily and concurrently populated growing-only slice
// of items (of type `int64`)
// which may be traversed in parallel to it's growth.
//
// Usage for a pile `p`:
//  p := MakeInt64Pile(128, 32)
//
// Have it grow concurrently using multiple:
//  var item int64 = something
//  p.Pile(item)
// in as many go routines as You may seem fit.
//
// In parallel, You may either
// traverse `p` in parallel right away:
//  for item, ok := p.Iter(); ok; item, ok = p.Next() { ... do sth with item ... }
// Here p.Iter() starts a new transversal with the first item (if any), and
// p.Next() keeps traverses the Int64Pile.
//
// or traverse blocking / awaiting close first:
//  for item := range <-p.Done() { ... do sth with item ... }
//
// or use the result when available:
//  r, p := <-p.Done(), nil
// Hint: here we get the result in `r` and at the same time discard / deallocate / forget the pile `p` itself.
//
// Note: The traversal is *not* intended to be concurrency safe!
// Thus: You may call `Pile` concurrently to Your traversal, but use of
// either `Done` or `Iter` and `Next` *must* be confined to a single go routine (thread).
//
type Int64Pile struct {
	pile   chan int64 // channel to receive further items
	list   []int64    // list of known items
	offset int        // index for Next()
}

// MakeInt64Pile returns a (pointer to a) fresh pile
// of items (of type `int64`)
// with size as initial capacity
// and
// with buff as initial leeway, allowing as many Pile's to execute non-blocking before respective Done or Next's.
func MakeInt64Pile(size, buff int) *Int64Pile {
	pile := new(Int64Pile)
	pile.list = make([]int64, 0, size)
	pile.pile = make(chan int64, buff)
	return pile
}

// Pile appends an `int64` item to the Int64Pile.
//
// Note: Pile will block iff buff is exceeded and no Done() or Next()'s are used.
func (d *Int64Pile) Pile(item int64) {
	d.pile <- item
}

// Close - call once when everything has been piled.
//
// Close intentionally implements io.Closer
//
// Note: After Close(),
// any Close(...) will panic
// and
// any Pile(...) will panic
// and
// any Done() or Next() will return immediately: no eventual blocking, that is.
func (d *Int64Pile) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				panic(r)
			}
		}
	}()
	close(d.pile)
	return nil
}

// Iter puts the pile iterator back to the beginning
// and returns the first `Next()`, iff any.
// Usage for a pile `p`:
//  for item, ok := p.Iter(); ok; item, ok = p.Next() { ... do sth with item ... }
func (d *Int64Pile) Iter() (item int64, ok bool) {
	d.offset = 0
	return d.Next()
}

// Next returns the next item,
// or false iff the pile is exhausted.
//
// Note: Iff the pile is not closed yet,
// Next may block, awaiting some Pile().
func (d *Int64Pile) Next() (item int64, ok bool) {
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

// Done returns a channel which emits the result (as slice of Int64) once the pile is closed.
//
// Users of Done() *must not* iterate (via Iter() Next()...) before the done-channel is closed!
//
// Done is a convenience - useful iff You do not like/need to start any traversal before the pile is fully populated.
// Once the pile is closed, Done() will signal in constant time.
//
// Note: Upon signalling, the pile is reset to it's tip,
// so You may traverse it (via Next) right away.
// Usage for a pile `p`:
// Traverse blocking / awaiting close first:
//  for item := range <-p.Done() { ... do sth with item ... }
// or use the result when available
//  r, p := <-p.Done(), nil
// while discaring the pile itself.
func (d *Int64Pile) Done() (done <-chan []int64) {
	cha := make(chan []int64)
	go func(cha chan<- []int64, d *Int64Pile) {
		defer close(cha)
		d.offset = 0
		if len(d.list) > d.offset {
			// skip what's already known
			d.offset = len(d.list) - 1
		}
		for _, ok := d.Next(); ok; _, ok = d.Next() {
			// keep draining
		}
		d.offset = 0  // reset
		cha <- d.list // signal the result, and terminate
	}(cha, d)
	return cha
}
