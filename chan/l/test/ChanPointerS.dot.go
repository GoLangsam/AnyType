// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type PointerSChan interface { // bidirectional channel
	PointerSROnlyChan // aka "<-chan" - receive only
	PointerSSOnlyChan // aka "chan<-" - send only
}

type PointerSROnlyChan interface { // receive-only channel
	RequestPointerS() (dat []*SomeType)        // the receive function - aka "some-new-PointerS-var := <-MyKind"
	TryPointerS() (dat []*SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-PointerS-var, ok := <-MyKind"
}

type PointerSSOnlyChan interface { // send-only channel
	ProvidePointerS(dat []*SomeType) // the send function - aka "MyKind <- some PointerS"
}

type DChPointerS struct { // demand channel
	dat chan []*SomeType
	req chan struct{}
}

func MakeDemandPointerSChan() *DChPointerS {
	d := new(DChPointerS)
	d.dat = make(chan []*SomeType)
	d.req = make(chan struct{})
	return d
}

func MakeDemandPointerSBuff(cap int) *DChPointerS {
	d := new(DChPointerS)
	d.dat = make(chan []*SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePointerS is the send function - aka "MyKind <- some PointerS"
func (c *DChPointerS) ProvidePointerS(dat []*SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestPointerS is the receive function - aka "some PointerS <- MyKind"
func (c *DChPointerS) RequestPointerS() (dat []*SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPointerS is the comma-ok multi-valued form of RequestPointerS and
// reports whether a received value was sent before the PointerS channel was closed.
func (c *DChPointerS) TryPointerS() (dat []*SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
