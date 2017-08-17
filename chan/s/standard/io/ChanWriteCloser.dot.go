// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriteCloserChan interface { // bidirectional channel
	WriteCloserROnlyChan // aka "<-chan" - receive only
	WriteCloserSOnlyChan // aka "chan<-" - send only
}

type WriteCloserROnlyChan interface { // receive-only channel
	RequestWriteCloser() (dat io.WriteCloser)        // the receive function - aka "some-new-WriteCloser-var := <-MyKind"
	TryWriteCloser() (dat io.WriteCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriteCloser-var, ok := <-MyKind"
}

type WriteCloserSOnlyChan interface { // send-only channel
	ProvideWriteCloser(dat io.WriteCloser) // the send function - aka "MyKind <- some WriteCloser"
}

type SChWriteCloser struct { // supply channel
	dat chan io.WriteCloser
	// req chan struct{}
}

func MakeSupplyWriteCloserChan() *SChWriteCloser {
	d := new(SChWriteCloser)
	d.dat = make(chan io.WriteCloser)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyWriteCloserBuff(cap int) *SChWriteCloser {
	d := new(SChWriteCloser)
	d.dat = make(chan io.WriteCloser, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriteCloser is the send function - aka "MyKind <- some WriteCloser"
func (c *SChWriteCloser) ProvideWriteCloser(dat io.WriteCloser) {
	// .req
	c.dat <- dat
}

// RequestWriteCloser is the receive function - aka "some WriteCloser <- MyKind"
func (c *SChWriteCloser) RequestWriteCloser() (dat io.WriteCloser) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriteCloser is the comma-ok multi-valued form of RequestWriteCloser and
// reports whether a received value was sent before the WriteCloser channel was closed.
func (c *SChWriteCloser) TryWriteCloser() (dat io.WriteCloser, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
