// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type ByteChan interface { // bidirectional channel
	ByteROnlyChan // aka "<-chan" - receive only
	ByteSOnlyChan // aka "chan<-" - send only
}

type ByteROnlyChan interface { // receive-only channel
	RequestByte() (dat byte)        // the receive function - aka "some-new-Byte-var := <-MyKind"
	TryByte() (dat byte, open bool) // the multi-valued comma-ok receive function - aka "some-new-Byte-var, ok := <-MyKind"
}

type ByteSOnlyChan interface { // send-only channel
	ProvideByte(dat byte) // the send function - aka "MyKind <- some Byte"
}

type SChByte struct { // supply channel
	dat chan byte
	// req chan struct{}
}

func MakeSupplyByteChan() *SChByte {
	d := new(SChByte)
	d.dat = make(chan byte)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyByteBuff(cap int) *SChByte {
	d := new(SChByte)
	d.dat = make(chan byte, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByte is the send function - aka "MyKind <- some Byte"
func (c *SChByte) ProvideByte(dat byte) {
	// .req
	c.dat <- dat
}

// RequestByte is the receive function - aka "some Byte <- MyKind"
func (c *SChByte) RequestByte() (dat byte) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByte is the comma-ok multi-valued form of RequestByte and
// reports whether a received value was sent before the Byte channel was closed.
func (c *SChByte) TryByte() (dat byte, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
