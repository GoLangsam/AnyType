// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UIntChan interface { // bidirectional channel
	UIntROnlyChan // aka "<-chan" - receive only
	UIntSOnlyChan // aka "chan<-" - send only
}

type UIntROnlyChan interface { // receive-only channel
	RequestUInt() (dat uint)        // the receive function - aka "some-new-UInt-var := <-MyKind"
	TryUInt() (dat uint, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt-var, ok := <-MyKind"
}

type UIntSOnlyChan interface { // send-only channel
	ProvideUInt(dat uint) // the send function - aka "MyKind <- some UInt"
}

type DChUInt struct { // demand channel
	dat chan uint
	req chan struct{}
}

func MakeDemandUIntChan() *DChUInt {
	d := new(DChUInt)
	d.dat = make(chan uint)
	d.req = make(chan struct{})
	return d
}

func MakeDemandUIntBuff(cap int) *DChUInt {
	d := new(DChUInt)
	d.dat = make(chan uint, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt is the send function - aka "MyKind <- some UInt"
func (c *DChUInt) ProvideUInt(dat uint) {
	<-c.req
	c.dat <- dat
}

// RequestUInt is the receive function - aka "some UInt <- MyKind"
func (c *DChUInt) RequestUInt() (dat uint) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt is the comma-ok multi-valued form of RequestUInt and
// reports whether a received value was sent before the UInt channel was closed.
func (c *DChUInt) TryUInt() (dat uint, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeUInt returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UIntchannel is returned.
func MergeUInt(inps ...<-chan uint) (out <-chan uint) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt2(inps[0], MergeUInt(inps[1:]...))
	}
}

// mergeUInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt2(i1, i2 <-chan uint) (out <-chan uint) {
	cha := make(chan uint)
	go func(out chan<- uint, i1, i2 <-chan uint) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 uint // what we've read
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
