// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt32Chan interface { // bidirectional channel
	UInt32ROnlyChan // aka "<-chan" - receive only
	UInt32SOnlyChan // aka "chan<-" - send only
}

type UInt32ROnlyChan interface { // receive-only channel
	RequestUInt32() (dat uint32)        // the receive function - aka "some-new-UInt32-var := <-MyKind"
	TryUInt32() (dat uint32, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt32-var, ok := <-MyKind"
}

type UInt32SOnlyChan interface { // send-only channel
	ProvideUInt32(dat uint32) // the send function - aka "MyKind <- some UInt32"
}

type DChUInt32 struct { // demand channel
	dat chan uint32
	req chan struct{}
}

func MakeDemandUInt32Chan() *DChUInt32 {
	d := new(DChUInt32)
	d.dat = make(chan uint32)
	d.req = make(chan struct{})
	return d
}

func MakeDemandUInt32Buff(cap int) *DChUInt32 {
	d := new(DChUInt32)
	d.dat = make(chan uint32, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideUInt32 is the send function - aka "MyKind <- some UInt32"
func (c *DChUInt32) ProvideUInt32(dat uint32) {
	<-c.req
	c.dat <- dat
}

// RequestUInt32 is the receive function - aka "some UInt32 <- MyKind"
func (c *DChUInt32) RequestUInt32() (dat uint32) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryUInt32 is the comma-ok multi-valued form of RequestUInt32 and
// reports whether a received value was sent before the UInt32 channel was closed.
func (c *DChUInt32) TryUInt32() (dat uint32, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
