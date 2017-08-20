// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int32Chan interface { // bidirectional channel
	Int32ROnlyChan // aka "<-chan" - receive only
	Int32SOnlyChan // aka "chan<-" - send only
}

type Int32ROnlyChan interface { // receive-only channel
	RequestInt32() (dat int32)        // the receive function - aka "some-new-Int32-var := <-MyKind"
	TryInt32() (dat int32, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int32-var, ok := <-MyKind"
}

type Int32SOnlyChan interface { // send-only channel
	ProvideInt32(dat int32) // the send function - aka "MyKind <- some Int32"
}

type SChInt32 struct { // supply channel
	dat chan int32
	// req chan struct{}
}

func MakeSupplyInt32Chan() *SChInt32 {
	d := new(SChInt32)
	d.dat = make(chan int32)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyInt32Buff(cap int) *SChInt32 {
	d := new(SChInt32)
	d.dat = make(chan int32, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt32 is the send function - aka "MyKind <- some Int32"
func (c *SChInt32) ProvideInt32(dat int32) {
	// .req
	c.dat <- dat
}

// RequestInt32 is the receive function - aka "some Int32 <- MyKind"
func (c *SChInt32) RequestInt32() (dat int32) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt32 is the comma-ok multi-valued form of RequestInt32 and
// reports whether a received value was sent before the Int32 channel was closed.
func (c *SChInt32) TryInt32() (dat int32, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeInt32 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int32channel is returned.
func MergeInt32(inps ...<-chan int32) (out <-chan int32) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int32)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt322(inps[0], MergeInt32(inps[1:]...))
	}
}

// mergeInt322 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt322(i1, i2 <-chan int32) (out <-chan int32) {
	cha := make(chan int32)
	go func(out chan<- int32, i1, i2 <-chan int32) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int32 // what we've read
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
