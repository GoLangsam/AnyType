// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type ErrorSChan interface { // bidirectional channel
	ErrorSROnlyChan // aka "<-chan" - receive only
	ErrorSSOnlyChan // aka "chan<-" - send only
}

type ErrorSROnlyChan interface { // receive-only channel
	RequestErrorS() (dat []error)        // the receive function - aka "some-new-ErrorS-var := <-MyKind"
	TryErrorS() (dat []error, open bool) // the multi-valued comma-ok receive function - aka "some-new-ErrorS-var, ok := <-MyKind"
}

type ErrorSSOnlyChan interface { // send-only channel
	ProvideErrorS(dat []error) // the send function - aka "MyKind <- some ErrorS"
}

type SChErrorS struct { // supply channel
	dat chan []error
	// req chan struct{}
}

func MakeSupplyErrorSChan() *SChErrorS {
	d := new(SChErrorS)
	d.dat = make(chan []error)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyErrorSBuff(cap int) *SChErrorS {
	d := new(SChErrorS)
	d.dat = make(chan []error, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideErrorS is the send function - aka "MyKind <- some ErrorS"
func (c *SChErrorS) ProvideErrorS(dat []error) {
	// .req
	c.dat <- dat
}

// RequestErrorS is the receive function - aka "some ErrorS <- MyKind"
func (c *SChErrorS) RequestErrorS() (dat []error) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryErrorS is the comma-ok multi-valued form of RequestErrorS and
// reports whether a received value was sent before the ErrorS channel was closed.
func (c *SChErrorS) TryErrorS() (dat []error, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len