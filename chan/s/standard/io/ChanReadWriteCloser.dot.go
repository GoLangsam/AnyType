// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriteCloserChan interface { // bidirectional channel
	ReadWriteCloserROnlyChan // aka "<-chan" - receive only
	ReadWriteCloserSOnlyChan // aka "chan<-" - send only
}

type ReadWriteCloserROnlyChan interface { // receive-only channel
	RequestReadWriteCloser() (dat io.ReadWriteCloser)        // the receive function - aka "some-new-ReadWriteCloser-var := <-MyKind"
	TryReadWriteCloser() (dat io.ReadWriteCloser, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriteCloser-var, ok := <-MyKind"
}

type ReadWriteCloserSOnlyChan interface { // send-only channel
	ProvideReadWriteCloser(dat io.ReadWriteCloser) // the send function - aka "MyKind <- some ReadWriteCloser"
}

type SChReadWriteCloser struct { // supply channel
	dat chan io.ReadWriteCloser
	// req chan struct{}
}

func MakeSupplyReadWriteCloserChan() *SChReadWriteCloser {
	d := new(SChReadWriteCloser)
	d.dat = make(chan io.ReadWriteCloser)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyReadWriteCloserBuff(cap int) *SChReadWriteCloser {
	d := new(SChReadWriteCloser)
	d.dat = make(chan io.ReadWriteCloser, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadWriteCloser is the send function - aka "MyKind <- some ReadWriteCloser"
func (c *SChReadWriteCloser) ProvideReadWriteCloser(dat io.ReadWriteCloser) {
	// .req
	c.dat <- dat
}

// RequestReadWriteCloser is the receive function - aka "some ReadWriteCloser <- MyKind"
func (c *SChReadWriteCloser) RequestReadWriteCloser() (dat io.ReadWriteCloser) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadWriteCloser is the comma-ok multi-valued form of RequestReadWriteCloser and
// reports whether a received value was sent before the ReadWriteCloser channel was closed.
func (c *SChReadWriteCloser) TryReadWriteCloser() (dat io.ReadWriteCloser, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
