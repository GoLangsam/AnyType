// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// UInt16Chan represents a
// bidirectional
// channel
type UInt16Chan interface {
	UInt16ROnlyChan // aka "<-chan" - receive only
	UInt16SOnlyChan // aka "chan<-" - send only
}

// UInt16ROnlyChan represents a
// receive-only
// channel
type UInt16ROnlyChan interface {
	RequestUInt16() (dat uint16)        // the receive function - aka "MyUInt16 := <-MyUInt16ROnlyChan"
	TryUInt16() (dat uint16, open bool) // the multi-valued comma-ok receive function - aka "MyUInt16, ok := <-MyUInt16ROnlyChan"
}

// UInt16SOnlyChan represents a
// send-only
// channel
type UInt16SOnlyChan interface {
	ProvideUInt16(dat uint16) // the send function - aka "MyKind <- some UInt16"
}

// DChUInt16 is a demand channel
type DChUInt16 struct {
	dat chan uint16
	req chan struct{}
}

// MakeDemandUInt16Chan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandUInt16Chan() *DChUInt16 {
	d := new(DChUInt16)
	d.dat = make(chan uint16)
	d.req = make(chan struct{})
	return d
}

// MakeDemandUInt16Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandUInt16Buff(cap int) *DChUInt16 {
	d := new(DChUInt16)
	d.dat = make(chan uint16, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt16 is the send function - aka "MyKind <- some UInt16"
func (c *DChUInt16) ProvideUInt16(dat uint16) {
	<-c.req
	c.dat <- dat
}

// RequestUInt16 is the receive function - aka "some UInt16 <- MyKind"
func (c *DChUInt16) RequestUInt16() (dat uint16) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt16 is the comma-ok multi-valued form of RequestUInt16 and
// reports whether a received value was sent before the UInt16 channel was closed.
func (c *DChUInt16) TryUInt16() (dat uint16, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChUInt16 is a supply channel
type SChUInt16 struct {
	dat chan uint16
	// req chan struct{}
}

// MakeSupplyUInt16Chan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyUInt16Chan() *SChUInt16 {
	d := new(SChUInt16)
	d.dat = make(chan uint16)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyUInt16Buff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyUInt16Buff(cap int) *SChUInt16 {
	d := new(SChUInt16)
	d.dat = make(chan uint16, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt16 is the send function - aka "MyKind <- some UInt16"
func (c *SChUInt16) ProvideUInt16(dat uint16) {
	// .req
	c.dat <- dat
}

// RequestUInt16 is the receive function - aka "some UInt16 <- MyKind"
func (c *SChUInt16) RequestUInt16() (dat uint16) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt16 is the comma-ok multi-valued form of RequestUInt16 and
// reports whether a received value was sent before the UInt16 channel was closed.
func (c *SChUInt16) TryUInt16() (dat uint16, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeUInt16 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed UInt16channel is returned.
func MergeUInt16(inps ...<-chan uint16) (out <-chan uint16) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan uint16)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeUInt162(inps[0], MergeUInt16(inps[1:]...))
	}
}

// mergeUInt162 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeUInt162(i1, i2 <-chan uint16) (out <-chan uint16) {
	cha := make(chan uint16)
	go func(out chan<- uint16, i1, i2 <-chan uint16) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint16 // what we've read
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
