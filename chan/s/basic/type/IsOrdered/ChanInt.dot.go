// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type IntChan interface { // bidirectional channel
	IntROnlyChan // aka "<-chan" - receive only
	IntSOnlyChan // aka "chan<-" - send only
}

type IntROnlyChan interface { // receive-only channel
	RequestInt() (dat int)        // the receive function - aka "some-new-Int-var := <-MyKind"
	TryInt() (dat int, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int-var, ok := <-MyKind"
}

type IntSOnlyChan interface { // send-only channel
	ProvideInt(dat int) // the send function - aka "MyKind <- some Int"
}

type SChInt struct { // supply channel
	dat chan int
	// req chan struct{}
}

func MakeSupplyIntChan() *SChInt {
	d := new(SChInt)
	d.dat = make(chan int)
	// d.req = make(chan struct{})
	return d
}

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
// MergeInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeInt2(i1, i2 <-chan int) (out <-chan int) {
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
