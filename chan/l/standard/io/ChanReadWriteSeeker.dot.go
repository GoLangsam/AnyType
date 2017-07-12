// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriteSeekerChan interface { // bidirectional channel
	ReadWriteSeekerROnlyChan // aka "<-chan" - receive only
	ReadWriteSeekerSOnlyChan // aka "chan<-" - send only
}

type ReadWriteSeekerROnlyChan interface { // receive-only channel
	RequestReadWriteSeeker() (dat io.ReadWriteSeeker)        // the receive function - aka "some-new-ReadWriteSeeker-var := <-MyKind"
	TryReadWriteSeeker() (dat io.ReadWriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriteSeeker-var, ok := <-MyKind"
}

type ReadWriteSeekerSOnlyChan interface { // send-only channel
	ProvideReadWriteSeeker(dat io.ReadWriteSeeker) // the send function - aka "MyKind <- some ReadWriteSeeker"
}

type DChReadWriteSeeker struct { // demand channel
	dat chan io.ReadWriteSeeker
	req chan struct{}
}

func MakeDemandReadWriteSeekerChan() *DChReadWriteSeeker {
	d := new(DChReadWriteSeeker)
	d.dat = make(chan io.ReadWriteSeeker)
	d.req = make(chan struct{})
	return d
}

func MakeDemandReadWriteSeekerBuff(cap int) *DChReadWriteSeeker {
	d := new(DChReadWriteSeeker)
	d.dat = make(chan io.ReadWriteSeeker, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadWriteSeeker is the send function - aka "MyKind <- some ReadWriteSeeker"
func (c *DChReadWriteSeeker) ProvideReadWriteSeeker(dat io.ReadWriteSeeker) {
	<-c.req
	c.dat <- dat
}

// RequestReadWriteSeeker is the receive function - aka "some ReadWriteSeeker <- MyKind"
func (c *DChReadWriteSeeker) RequestReadWriteSeeker() (dat io.ReadWriteSeeker) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadWriteSeeker is the comma-ok multi-valued form of RequestReadWriteSeeker and
// reports whether a received value was sent before the ReadWriteSeeker channel was closed.
func (c *DChReadWriteSeeker) TryReadWriteSeeker() (dat io.ReadWriteSeeker, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
