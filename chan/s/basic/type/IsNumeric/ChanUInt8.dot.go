// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt8Chan interface { // bidirectional channel
	UInt8ROnlyChan // aka "<-chan" - receive only
	UInt8SOnlyChan // aka "chan<-" - send only
}

type UInt8ROnlyChan interface { // receive-only channel
	RequestUInt8() (dat uint8)        // the receive function - aka "some-new-UInt8-var := <-MyKind"
	TryUInt8() (dat uint8, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt8-var, ok := <-MyKind"
}

type UInt8SOnlyChan interface { // send-only channel
	ProvideUInt8(dat uint8) // the send function - aka "MyKind <- some UInt8"
}

type SChUInt8 struct { // supply channel
	dat chan uint8
	// req chan struct{}
}

func MakeSupplyUInt8Chan() *SChUInt8 {
	d := new(SChUInt8)
	d.dat = make(chan uint8)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyUInt8Buff(cap int) *SChUInt8 {
	d := new(SChUInt8)
	d.dat = make(chan uint8, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideUInt8 is the send function - aka "MyKind <- some UInt8"
func (c *SChUInt8) ProvideUInt8(dat uint8) {
	// .req
	c.dat <- dat
}

// RequestUInt8 is the receive function - aka "some UInt8 <- MyKind"
func (c *SChUInt8) RequestUInt8() (dat uint8) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryUInt8 is the comma-ok multi-valued form of RequestUInt8 and
// reports whether a received value was sent before the UInt8 channel was closed.
func (c *SChUInt8) TryUInt8() (dat uint8, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
