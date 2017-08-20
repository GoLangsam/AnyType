// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFoldSChan represents a
// bidirectional
// channel
type FsFoldSChan interface {
	FsFoldSROnlyChan // aka "<-chan" - receive only
	FsFoldSSOnlyChan // aka "chan<-" - send only
}

// FsFoldSROnlyChan represents a
// receive-only
// channel
type FsFoldSROnlyChan interface {
	RequestFsFoldS() (dat fs.FsFoldS)        // the receive function - aka "MyFsFoldS := <-MyFsFoldSROnlyChan"
	TryFsFoldS() (dat fs.FsFoldS, open bool) // the multi-valued comma-ok receive function - aka "MyFsFoldS, ok := <-MyFsFoldSROnlyChan"
}

// FsFoldSSOnlyChan represents a
// send-only
// channel
type FsFoldSSOnlyChan interface {
	ProvideFsFoldS(dat fs.FsFoldS) // the send function - aka "MyKind <- some FsFoldS"
}

// DChFsFoldS is a supply channel
type SChFsFoldS struct {
	dat chan fs.FsFoldS
	// req chan struct{}
}

// MakeSupplyFsFoldSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsFoldSChan() *SChFsFoldS {
	d := new(SChFsFoldS)
	d.dat = make(chan fs.FsFoldS)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsFoldSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsFoldSBuff(cap int) *SChFsFoldS {
	d := new(SChFsFoldS)
	d.dat = make(chan fs.FsFoldS, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsFoldS is the send function - aka "MyKind <- some FsFoldS"
func (c *SChFsFoldS) ProvideFsFoldS(dat fs.FsFoldS) {
	// .req
	c.dat <- dat
}

// RequestFsFoldS is the receive function - aka "some FsFoldS <- MyKind"
func (c *SChFsFoldS) RequestFsFoldS() (dat fs.FsFoldS) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsFoldS is the comma-ok multi-valued form of RequestFsFoldS and
// reports whether a received value was sent before the FsFoldS channel was closed.
func (c *SChFsFoldS) TryFsFoldS() (dat fs.FsFoldS, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
