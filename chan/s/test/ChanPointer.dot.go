// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type PointerChan interface { // bidirectional channel
	PointerROnlyChan // aka "<-chan" - receive only
	PointerSOnlyChan // aka "chan<-" - send only
}

type PointerROnlyChan interface { // receive-only channel
	RequestPointer() (dat *SomeType)        // the receive function - aka "some-new-Pointer-var := <-MyKind"
	TryPointer() (dat *SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-Pointer-var, ok := <-MyKind"
}

type PointerSOnlyChan interface { // send-only channel
	ProvidePointer(dat *SomeType) // the send function - aka "MyKind <- some Pointer"
}

type SChPointer struct { // supply channel
	dat chan *SomeType
	// req chan struct{}
}

func MakeSupplyPointerChan() *SChPointer {
	d := new(SChPointer)
	d.dat = make(chan *SomeType)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyPointerBuff(cap int) *SChPointer {
	d := new(SChPointer)
	d.dat = make(chan *SomeType, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvidePointer is the send function - aka "MyKind <- some Pointer"
func (c *SChPointer) ProvidePointer(dat *SomeType) {
	// .req
	c.dat <- dat
}

// RequestPointer is the receive function - aka "some Pointer <- MyKind"
func (c *SChPointer) RequestPointer() (dat *SomeType) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryPointer is the comma-ok multi-valued form of RequestPointer and
// reports whether a received value was sent before the Pointer channel was closed.
func (c *SChPointer) TryPointer() (dat *SomeType, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
