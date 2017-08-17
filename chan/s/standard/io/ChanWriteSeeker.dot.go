// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriteSeekerChan interface { // bidirectional channel
	WriteSeekerROnlyChan // aka "<-chan" - receive only
	WriteSeekerSOnlyChan // aka "chan<-" - send only
}

type WriteSeekerROnlyChan interface { // receive-only channel
	RequestWriteSeeker() (dat io.WriteSeeker)        // the receive function - aka "some-new-WriteSeeker-var := <-MyKind"
	TryWriteSeeker() (dat io.WriteSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriteSeeker-var, ok := <-MyKind"
}

type WriteSeekerSOnlyChan interface { // send-only channel
	ProvideWriteSeeker(dat io.WriteSeeker) // the send function - aka "MyKind <- some WriteSeeker"
}

type SChWriteSeeker struct { // supply channel
	dat chan io.WriteSeeker
	// req chan struct{}
}

func MakeSupplyWriteSeekerChan() *SChWriteSeeker {
	d := new(SChWriteSeeker)
	d.dat = make(chan io.WriteSeeker)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyWriteSeekerBuff(cap int) *SChWriteSeeker {
	d := new(SChWriteSeeker)
	d.dat = make(chan io.WriteSeeker, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriteSeeker is the send function - aka "MyKind <- some WriteSeeker"
func (c *SChWriteSeeker) ProvideWriteSeeker(dat io.WriteSeeker) {
	// .req
	c.dat <- dat
}

// RequestWriteSeeker is the receive function - aka "some WriteSeeker <- MyKind"
func (c *SChWriteSeeker) RequestWriteSeeker() (dat io.WriteSeeker) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriteSeeker is the comma-ok multi-valued form of RequestWriteSeeker and
// reports whether a received value was sent before the WriteSeeker channel was closed.
func (c *SChWriteSeeker) TryWriteSeeker() (dat io.WriteSeeker, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
