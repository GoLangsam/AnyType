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

type DChByte struct { // demand channel
	dat chan byte
	req chan struct{}
}

func MakeDemandByteChan() *DChByte {
	d := new(DChByte)
	d.dat = make(chan byte)
	d.req = make(chan struct{})
	return d
}

func MakeDemandByteBuff(cap int) *DChByte {
	d := new(DChByte)
	d.dat = make(chan byte, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByte is the send function - aka "MyKind <- some Byte"
func (c *DChByte) ProvideByte(dat byte) {
	<-c.req
	c.dat <- dat
}

// RequestByte is the receive function - aka "some Byte <- MyKind"
func (c *DChByte) RequestByte() (dat byte) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByte is the comma-ok multi-valued form of RequestByte and
// reports whether a received value was sent before the Byte channel was closed.
func (c *DChByte) TryByte() (dat byte, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
