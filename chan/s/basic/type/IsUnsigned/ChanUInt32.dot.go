// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt32Chan interface { // bidirectional channel
	UInt32ROnlyChan // aka "<-chan" - receive only
	UInt32SOnlyChan // aka "chan<-" - send only
}

type UInt32ROnlyChan interface { // receive-only channel
	RequestUInt32() (dat uint32)        // the receive function - aka "some-new-UInt32-var := <-MyKind"
	TryUInt32() (dat uint32, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt32-var, ok := <-MyKind"
}

type UInt32SOnlyChan interface { // send-only channel
	ProvideUInt32(dat uint32) // the send function - aka "MyKind <- some UInt32"
}

type SChUInt32 struct { // supply channel
	dat chan uint32
	// req chan struct{}
}

func MakeSupplyUInt32Chan() *SChUInt32 {
	d := new(SChUInt32)
	d.dat = make(chan uint32)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyUInt32Buff(cap int) *SChUInt32 {
	d := new(SChUInt32)
	d.dat = make(chan uint32, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt32 is the send function - aka "MyKind <- some UInt32"
func (c *SChUInt32) ProvideUInt32(dat uint32) {
	// .req
	c.dat <- dat
}

// RequestUInt32 is the receive function - aka "some UInt32 <- MyKind"
func (c *SChUInt32) RequestUInt32() (dat uint32) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt32 is the comma-ok multi-valued form of RequestUInt32 and
// reports whether a received value was sent before the UInt32 channel was closed.
func (c *SChUInt32) TryUInt32() (dat uint32, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeUInt32 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt32channel is returned.
func MergeUInt32(inps ...<-chan uint32) (out <-chan uint32) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint32)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt322(inps[0], MergeUInt32(inps[1:]...))
	}
}

// mergeUInt322 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt322(i1, i2 <-chan uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go func(out chan<- uint32, i1, i2 <-chan uint32) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint32 // what we've read
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
