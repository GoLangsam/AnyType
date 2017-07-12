// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Complex128Chan interface { // bidirectional channel
	Complex128ROnlyChan // aka "<-chan" - receive only
	Complex128SOnlyChan // aka "chan<-" - send only
}

type Complex128ROnlyChan interface { // receive-only channel
	RequestComplex128() (dat complex128)        // the receive function - aka "some-new-Complex128-var := <-MyKind"
	TryComplex128() (dat complex128, open bool) // the multi-valued comma-ok receive function - aka "some-new-Complex128-var, ok := <-MyKind"
}

type Complex128SOnlyChan interface { // send-only channel
	ProvideComplex128(dat complex128) // the send function - aka "MyKind <- some Complex128"
}

type DChComplex128 struct { // demand channel
	dat chan complex128
	req chan struct{}
}

func MakeDemandComplex128Chan() *DChComplex128 {
	d := new(DChComplex128)
	d.dat = make(chan complex128)
	d.req = make(chan struct{})
	return d
}

func MakeDemandComplex128Buff(cap int) *DChComplex128 {
	d := new(DChComplex128)
	d.dat = make(chan complex128, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideComplex128 is the send function - aka "MyKind <- some Complex128"
func (c *DChComplex128) ProvideComplex128(dat complex128) {
	<-c.req
	c.dat <- dat
}

// RequestComplex128 is the receive function - aka "some Complex128 <- MyKind"
func (c *DChComplex128) RequestComplex128() (dat complex128) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryComplex128 is the comma-ok multi-valued form of RequestComplex128 and
// reports whether a received value was sent before the Complex128 channel was closed.
func (c *DChComplex128) TryComplex128() (dat complex128, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
