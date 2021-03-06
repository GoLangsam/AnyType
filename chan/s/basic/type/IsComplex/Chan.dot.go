// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsComplex

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Chan represents a
// bidirectional
// channel
type Chan interface {
	ROnlyChan // aka "<-chan" - receive only
	SOnlyChan // aka "chan<-" - send only
}

// ROnlyChan represents a
// receive-only
// channel
type ROnlyChan interface {
	Request() (dat complex128)        // the receive function - aka "My := <-MyROnlyChan"
	Try() (dat complex128, open bool) // the multi-valued comma-ok receive function - aka "My, ok := <-MyROnlyChan"
}

// SOnlyChan represents a
// send-only
// channel
type SOnlyChan interface {
	Provide(dat complex128) // the send function - aka "MyKind <- some "
}

// SCh is a supply channel
type SCh struct {
	dat chan complex128
	// req chan struct{}
}

// MakeSupplyChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyChan() *SCh {
	d := new(SCh)
	d.dat = make(chan complex128)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyBuff(cap int) *SCh {
	d := new(SCh)
	d.dat = make(chan complex128, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// Provide is the send function - aka "MyKind <- some "
func (c *SCh) Provide(dat complex128) {
	// .req
	c.dat <- dat
}

// Request is the receive function - aka "some  <- MyKind"
func (c *SCh) Request() (dat complex128) {
	// eq <- struct{}{}
	return <-c.dat
}

// Try is the comma-ok multi-valued form of Request and
// reports whether a received value was sent before the  channel was closed.
func (c *SCh) Try() (dat complex128, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
