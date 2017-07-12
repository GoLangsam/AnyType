// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type CloserChan interface { // bidirectional channel
	CloserROnlyChan // aka "<-chan" - receive only
	CloserSOnlyChan // aka "chan<-" - send only
}

type CloserROnlyChan interface { // receive-only channel
	RequestCloser() (dat io.Closer)        // the receive function - aka "some-new-Closer-var := <-MyKind"
	TryCloser() (dat io.Closer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Closer-var, ok := <-MyKind"
}

type CloserSOnlyChan interface { // send-only channel
	ProvideCloser(dat io.Closer) // the send function - aka "MyKind <- some Closer"
}

type SChCloser struct { // supply channel
	dat chan io.Closer
	// req chan struct{}
}

func MakeSupplyCloserChan() *SChCloser {
	d := new(SChCloser)
	d.dat = make(chan io.Closer)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyCloserBuff(cap int) *SChCloser {
	d := new(SChCloser)
	d.dat = make(chan io.Closer, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideCloser is the send function - aka "MyKind <- some Closer"
func (c *SChCloser) ProvideCloser(dat io.Closer) {
	// .req
	c.dat <- dat
}

// RequestCloser is the receive function - aka "some Closer <- MyKind"
func (c *SChCloser) RequestCloser() (dat io.Closer) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryCloser is the comma-ok multi-valued form of RequestCloser and
// reports whether a received value was sent before the Closer channel was closed.
func (c *SChCloser) TryCloser() (dat io.Closer, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
