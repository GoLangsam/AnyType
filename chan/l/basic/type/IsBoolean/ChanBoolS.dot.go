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

type DChBoolS struct { // demand channel
	dat chan []bool
	req chan struct{}
}

func MakeDemandBoolSChan() *DChBoolS {
	d := new(DChBoolS)
	d.dat = make(chan []bool)
	d.req = make(chan struct{})
	return d
}

func MakeDemandBoolSBuff(cap int) *DChBoolS {
	d := new(DChBoolS)
	d.dat = make(chan []bool, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideBoolS is the send function - aka "MyKind <- some BoolS"
func (c *DChBoolS) ProvideBoolS(dat []bool) {
	<-c.req
	c.dat <- dat
}

// RequestBoolS is the receive function - aka "some BoolS <- MyKind"
func (c *DChBoolS) RequestBoolS() (dat []bool) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryBoolS is the comma-ok multi-valued form of RequestBoolS and
// reports whether a received value was sent before the BoolS channel was closed.
func (c *DChBoolS) TryBoolS() (dat []bool, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
