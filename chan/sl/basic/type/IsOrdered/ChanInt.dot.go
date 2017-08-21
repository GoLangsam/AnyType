// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// IntChan represents a
// bidirectional
// channel
type IntChan interface {
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

// IntROnlyChan represents a
// receive-only
// channel
type IntROnlyChan interface {
	RequestInt() (dat int)        // the receive function - aka "MyInt := <-MyIntROnlyChan"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "MyInt, ok := <-MyIntROnlyChan"
}

// IntSOnlyChan represents a
// send-only
// channel
type IntSOnlyChan interface {
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}

// DChInt is a demand channel
type DChInt struct {
	dat chan int
	req chan struct{}
}

// MakeDemandIntChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandIntChan() *DChInt {
	d := new(DChInt)
	d.dat = make(chan int)
	d.req = make(chan struct{})
	return d
}

// MakeDemandIntBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandIntBuff(cap int) *DChInt {
	d := new(DChInt)
	d.dat = make(chan int, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt is the send function - aka "MyKind <- some Int"
func (c *DChInt) ProvideInt(dat int) {
	<-c.req
	c.dat <- dat
}

// RequestInt is the receive function - aka "some Int <- MyKind"
func (c *DChInt) RequestInt() (dat int) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt is the comma-ok multi-valued form of RequestInt and
// reports whether a received value was sent before the Int channel was closed.
func (c *DChInt) TryInt() (dat int, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChInt is a supply channel
type SChInt struct {
	dat chan int
	// req chan struct{}
}

// MakeSupplyIntChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyIntChan() *SChInt {
	d := new(SChInt)
	d.dat = make(chan int)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyIntBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyIntBuff(cap int) *SChInt {
	d := new(SChInt)
	d.dat = make(chan int, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt is the send function - aka "MyKind <- some Int"
func (c *SChInt) ProvideInt(dat int) {
	// .req
	c.dat <- dat
}

// RequestInt is the receive function - aka "some Int <- MyKind"
func (c *SChInt) RequestInt() (dat int) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt is the comma-ok multi-valued form of RequestInt and
// reports whether a received value was sent before the Int channel was closed.
func (c *SChInt) TryInt() (dat int, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeInt returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Intchannel is returned.
func MergeInt(inps ...<-chan int) (out <-chan int) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt2(inps[0], MergeInt(inps[1:]...))
	}
}

// mergeInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt2(i1, i2 <-chan int) (out <-chan int) {
	cha := make(chan int)
	go func(out chan<- int, i1, i2 <-chan int) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 int  // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
