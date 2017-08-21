// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt64Chan represents a
// bidirectional
// channel
type UInt64Chan interface {
	UInt64ROnlyChan // aka "<-chan" - receive only
	UInt64SOnlyChan // aka "chan<-" - send only
}

// UInt64ROnlyChan represents a
// receive-only
// channel
type UInt64ROnlyChan interface {
	RequestUInt64() (dat uint64)        // the receive function - aka "MyUInt64 := <-MyUInt64ROnlyChan"
	TryUInt64() (dat uint64, open bool) // the multi-valued comma-ok receive function - aka "MyUInt64, ok := <-MyUInt64ROnlyChan"
}

// UInt64SOnlyChan represents a
// send-only
// channel
type UInt64SOnlyChan interface {
	ProvideUInt64(dat uint64) // the send function - aka "MyKind <- some UInt64"
}

// DChUInt64 is a demand channel
type DChUInt64 struct {
	dat chan uint64
	req chan struct{}
}

// MakeDemandUInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUInt64Chan() *DChUInt64 {
	d := new(DChUInt64)
	d.dat = make(chan uint64)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUInt64Buff(cap int) *DChUInt64 {
	d := new(DChUInt64)
	d.dat = make(chan uint64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt64 is the send function - aka "MyKind <- some UInt64"
func (c *DChUInt64) ProvideUInt64(dat uint64) {
	<-c.req
	c.dat <- dat
}

// RequestUInt64 is the receive function - aka "some UInt64 <- MyKind"
func (c *DChUInt64) RequestUInt64() (dat uint64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt64 is the comma-ok multi-valued form of RequestUInt64 and
// reports whether a received value was sent before the UInt64 channel was closed.
func (c *DChUInt64) TryUInt64() (dat uint64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// SChUInt64 is a supply channel
type SChUInt64 struct {
	dat chan uint64
	// req chan struct{}
}

// MakeSupplyUInt64Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt64Chan() *SChUInt64 {
	d := new(SChUInt64)
	d.dat = make(chan uint64)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt64Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt64Buff(cap int) *SChUInt64 {
	d := new(SChUInt64)
	d.dat = make(chan uint64, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt64 is the send function - aka "MyKind <- some UInt64"
func (c *SChUInt64) ProvideUInt64(dat uint64) {
	// .req
	c.dat <- dat
}

// RequestUInt64 is the receive function - aka "some UInt64 <- MyKind"
func (c *SChUInt64) RequestUInt64() (dat uint64) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt64 is the comma-ok multi-valued form of RequestUInt64 and
// reports whether a received value was sent before the UInt64 channel was closed.
func (c *SChUInt64) TryUInt64() (dat uint64, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeUInt64 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt64channel is returned.
func MergeUInt64(inps ...<-chan uint64) (out <-chan uint64) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint64)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt642(inps[0], MergeUInt64(inps[1:]...))
	}
}

// mergeUInt642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt642(i1, i2 <-chan uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	go func(out chan<- uint64, i1, i2 <-chan uint64) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint64 // what we've read
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
