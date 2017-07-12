// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Float32Chan interface { // bidirectional channel
	Float32ROnlyChan // aka "<-chan" - receive only
	Float32SOnlyChan // aka "chan<-" - send only
}

type Float32ROnlyChan interface { // receive-only channel
	RequestFloat32() (dat float32)        // the receive function - aka "some-new-Float32-var := <-MyKind"
	TryFloat32() (dat float32, open bool) // the multi-valued comma-ok receive function - aka "some-new-Float32-var, ok := <-MyKind"
}

type Float32SOnlyChan interface { // send-only channel
	ProvideFloat32(dat float32) // the send function - aka "MyKind <- some Float32"
}

type SChFloat32 struct { // supply channel
	dat chan float32
	// req chan struct{}
}

func MakeSupplyFloat32Chan() *SChFloat32 {
	d := new(SChFloat32)
	d.dat = make(chan float32)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyFloat32Buff(cap int) *SChFloat32 {
	d := new(SChFloat32)
	d.dat = make(chan float32, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFloat32 is the send function - aka "MyKind <- some Float32"
func (c *SChFloat32) ProvideFloat32(dat float32) {
	// .req
	c.dat <- dat
}

// RequestFloat32 is the receive function - aka "some Float32 <- MyKind"
func (c *SChFloat32) RequestFloat32() (dat float32) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFloat32 is the comma-ok multi-valued form of RequestFloat32 and
// reports whether a received value was sent before the Float32 channel was closed.
func (c *SChFloat32) TryFloat32() (dat float32, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
