// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadSeekerChan interface { // bidirectional channel
	ReadSeekerROnlyChan // aka "<-chan" - receive only
	ReadSeekerSOnlyChan // aka "chan<-" - send only
}

type ReadSeekerROnlyChan interface { // receive-only channel
	RequestReadSeeker() (dat io.ReadSeeker)        // the receive function - aka "some-new-ReadSeeker-var := <-MyKind"
	TryReadSeeker() (dat io.ReadSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadSeeker-var, ok := <-MyKind"
}

type ReadSeekerSOnlyChan interface { // send-only channel
	ProvideReadSeeker(dat io.ReadSeeker) // the send function - aka "MyKind <- some ReadSeeker"
}

type SChReadSeeker struct { // supply channel
	dat chan io.ReadSeeker
	// req chan struct{}
}

func MakeSupplyReadSeekerChan() *SChReadSeeker {
	d := new(SChReadSeeker)
	d.dat = make(chan io.ReadSeeker)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyReadSeekerBuff(cap int) *SChReadSeeker {
	d := new(SChReadSeeker)
	d.dat = make(chan io.ReadSeeker, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadSeeker is the send function - aka "MyKind <- some ReadSeeker"
func (c *SChReadSeeker) ProvideReadSeeker(dat io.ReadSeeker) {
	// .req
	c.dat <- dat
}

// RequestReadSeeker is the receive function - aka "some ReadSeeker <- MyKind"
func (c *SChReadSeeker) RequestReadSeeker() (dat io.ReadSeeker) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadSeeker is the comma-ok multi-valued form of RequestReadSeeker and
// reports whether a received value was sent before the ReadSeeker channel was closed.
func (c *SChReadSeeker) TryReadSeeker() (dat io.ReadSeeker, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
