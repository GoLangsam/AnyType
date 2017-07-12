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

type DChReadWriteCloser struct { // demand channel
	dat chan io.ReadWriteCloser
	req chan struct{}
}

func MakeDemandReadWriteCloserChan() *DChReadWriteCloser {
	d := new(DChReadWriteCloser)
	d.dat = make(chan io.ReadWriteCloser)
	d.req = make(chan struct{})
	return d
}

func MakeDemandReadWriteCloserBuff(cap int) *DChReadWriteCloser {
	d := new(DChReadWriteCloser)
	d.dat = make(chan io.ReadWriteCloser, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadWriteCloser is the send function - aka "MyKind <- some ReadWriteCloser"
func (c *DChReadWriteCloser) ProvideReadWriteCloser(dat io.ReadWriteCloser) {
	<-c.req
	c.dat <- dat
}

// RequestReadWriteCloser is the receive function - aka "some ReadWriteCloser <- MyKind"
func (c *DChReadWriteCloser) RequestReadWriteCloser() (dat io.ReadWriteCloser) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadWriteCloser is the comma-ok multi-valued form of RequestReadWriteCloser and
// reports whether a received value was sent before the ReadWriteCloser channel was closed.
func (c *DChReadWriteCloser) TryReadWriteCloser() (dat io.ReadWriteCloser, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
