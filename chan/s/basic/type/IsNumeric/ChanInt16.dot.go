// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int16Chan interface { // bidirectional channel
	Int16ROnlyChan // aka "<-chan" - receive only
	Int16SOnlyChan // aka "chan<-" - send only
}

type Int16ROnlyChan interface { // receive-only channel
	RequestInt16() (dat int16)        // the receive function - aka "some-new-Int16-var := <-MyKind"
	TryInt16() (dat int16, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int16-var, ok := <-MyKind"
}

type Int16SOnlyChan interface { // send-only channel
	ProvideInt16(dat int16) // the send function - aka "MyKind <- some Int16"
}

type SChInt16 struct { // supply channel
	dat chan int16
	// req chan struct{}
}

func MakeSupplyInt16Chan() *SChInt16 {
	d := new(SChInt16)
	d.dat = make(chan int16)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyInt16Buff(cap int) *SChInt16 {
	d := new(SChInt16)
	d.dat = make(chan int16, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt16 is the send function - aka "MyKind <- some Int16"
func (c *SChInt16) ProvideInt16(dat int16) {
	// .req
	c.dat <- dat
}

// RequestInt16 is the receive function - aka "some Int16 <- MyKind"
func (c *SChInt16) RequestInt16() (dat int16) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt16 is the comma-ok multi-valued form of RequestInt16 and
// reports whether a received value was sent before the Int16 channel was closed.
func (c *SChInt16) TryInt16() (dat int16, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
