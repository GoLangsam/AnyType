// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt8Chan represents a
// bidirectional
// channel
type UInt8Chan interface {
	UInt8ROnlyChan // aka "<-chan" - receive only
	UInt8SOnlyChan // aka "chan<-" - send only
}

// UInt8ROnlyChan represents a
// receive-only
// channel
type UInt8ROnlyChan interface {
	RequestUInt8() (dat uint8)        // the receive function - aka "MyUInt8 := <-MyUInt8ROnlyChan"
	TryUInt8() (dat uint8, open bool) // the multi-valued comma-ok receive function - aka "MyUInt8, ok := <-MyUInt8ROnlyChan"
}

// UInt8SOnlyChan represents a
// send-only
// channel
type UInt8SOnlyChan interface {
	ProvideUInt8(dat uint8) // the send function - aka "MyKind <- some UInt8"
}

// SChUInt8 is a supply channel
type SChUInt8 struct {
	dat chan uint8
	// req chan struct{}
}

// MakeSupplyUInt8Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt8Chan() *SChUInt8 {
	d := new(SChUInt8)
	d.dat = make(chan uint8)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt8Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt8Buff(cap int) *SChUInt8 {
	d := new(SChUInt8)
	d.dat = make(chan uint8, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt8 is the send function - aka "MyKind <- some UInt8"
func (c *SChUInt8) ProvideUInt8(dat uint8) {
	// .req
	c.dat <- dat
}

// RequestUInt8 is the receive function - aka "some UInt8 <- MyKind"
func (c *SChUInt8) RequestUInt8() (dat uint8) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt8 is the comma-ok multi-valued form of RequestUInt8 and
// reports whether a received value was sent before the UInt8 channel was closed.
func (c *SChUInt8) TryUInt8() (dat uint8, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeUInt8 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt8channel is returned.
func MergeUInt8(inps ...<-chan uint8) (out <-chan uint8) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint8)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt82(inps[0], MergeUInt8(inps[1:]...))
	}
}

// mergeUInt82 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt82(i1, i2 <-chan uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, i1, i2 <-chan uint8) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 uint8 // what we've read
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
