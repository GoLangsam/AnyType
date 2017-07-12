// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

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