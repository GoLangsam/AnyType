// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadCloserChan interface { // bidirectional channel
	ReadCloserROnlyChan // aka "<-chan" - receive only
	ReadCloserSOnlyChan // aka "chan<-" - send only
}

type ReadCloserROnlyChan interface { // receive-only channel
	RequestReadCloser() (dat io.ReadCloser)        // the receive function - aka "some-new-ReadCloser-var := <-MyKind"
	TryReadCloser() (dat io.ReadCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadCloser-var, ok := <-MyKind"
}

type ReadCloserSOnlyChan interface { // send-only channel
	ProvideReadCloser(dat io.ReadCloser) // the send function - aka "MyKind <- some ReadCloser"
}

type SChReadCloser struct { // supply channel
	dat chan io.ReadCloser
	// req chan struct{}
}

func MakeSupplyReadCloserChan() *SChReadCloser {
	d := new(SChReadCloser)
	d.dat = make(chan io.ReadCloser)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyReadCloserBuff(cap int) *SChReadCloser {
	d := new(SChReadCloser)
	d.dat = make(chan io.ReadCloser, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadCloser is the send function - aka "MyKind <- some ReadCloser"
func (c *SChReadCloser) ProvideReadCloser(dat io.ReadCloser) {
	// .req
	c.dat <- dat
}

// RequestReadCloser is the receive function - aka "some ReadCloser <- MyKind"
func (c *SChReadCloser) RequestReadCloser() (dat io.ReadCloser) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadCloser is the comma-ok multi-valued form of RequestReadCloser and
// reports whether a received value was sent before the ReadCloser channel was closed.
func (c *SChReadCloser) TryReadCloser() (dat io.ReadCloser, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
