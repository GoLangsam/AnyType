// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsBoolean

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type BoolSChan interface { // bidirectional channel
	BoolSROnlyChan // aka "<-chan" - receive only
	BoolSSOnlyChan // aka "chan<-" - send only
}

type BoolSROnlyChan interface { // receive-only channel
	RequestBoolS() (dat []bool)        // the receive function - aka "some-new-BoolS-var := <-MyKind"
	TryBoolS() (dat []bool, open bool) // the multi-valued comma-ok receive function - aka "some-new-BoolS-var, ok := <-MyKind"
}

type BoolSSOnlyChan interface { // send-only channel
	ProvideBoolS(dat []bool) // the send function - aka "MyKind <- some BoolS"
}

type SChBoolS struct { // supply channel
	dat chan []bool
	// req chan struct{}
}

func MakeSupplyBoolSChan() *SChBoolS {
	d := new(SChBoolS)
	d.dat = make(chan []bool)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyBoolSBuff(cap int) *SChBoolS {
	d := new(SChBoolS)
	d.dat = make(chan []bool, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideBoolS is the send function - aka "MyKind <- some BoolS"
func (c *SChBoolS) ProvideBoolS(dat []bool) {
	// .req
	c.dat <- dat
}

// RequestBoolS is the receive function - aka "some BoolS <- MyKind"
func (c *SChBoolS) RequestBoolS() (dat []bool) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryBoolS is the comma-ok multi-valued form of RequestBoolS and
// reports whether a received value was sent before the BoolS channel was closed.
func (c *SChBoolS) TryBoolS() (dat []bool, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
