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

type SChUInt struct { // supply channel
	dat chan uint
	// req chan struct{}
}

func MakeSupplyUIntChan() *SChUInt {
	d := new(SChUInt)
	d.dat = make(chan uint)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyUIntBuff(cap int) *SChUInt {
	d := new(SChUInt)
	d.dat = make(chan uint, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt is the send function - aka "MyKind <- some UInt"
func (c *SChUInt) ProvideUInt(dat uint) {
	// .req
	c.dat <- dat
}

// RequestUInt is the receive function - aka "some UInt <- MyKind"
func (c *SChUInt) RequestUInt() (dat uint) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt is the comma-ok multi-valued form of RequestUInt and
// reports whether a received value was sent before the UInt channel was closed.
func (c *SChUInt) TryUInt() (dat uint, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
// MergeUInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeUInt2(i1, i2 <-chan uint) (out <-chan uint) {
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