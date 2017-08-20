// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFoldChan represents a
// bidirectional
// channel
type FsFoldChan interface {
	FsFoldROnlyChan // aka "<-chan" - receive only
	FsFoldSOnlyChan // aka "chan<-" - send only
}

// FsFoldROnlyChan represents a
// receive-only
// channel
type FsFoldROnlyChan interface {
	RequestFsFold() (dat *fs.FsFold)        // the receive function - aka "MyFsFold := <-MyFsFoldROnlyChan"
	TryFsFold() (dat *fs.FsFold, open bool) // the multi-valued comma-ok receive function - aka "MyFsFold, ok := <-MyFsFoldROnlyChan"
}

// FsFoldSOnlyChan represents a
// send-only
// channel
type FsFoldSOnlyChan interface {
	ProvideFsFold(dat *fs.FsFold) // the send function - aka "MyKind <- some FsFold"
}

// DChFsFold is a demand channel
type DChFsFold struct {
	dat chan *fs.FsFold
	req chan struct{}
}

// MakeDemandFsFoldChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsFoldChan() *DChFsFold {
	d := new(DChFsFold)
	d.dat = make(chan *fs.FsFold)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsFoldBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsFoldBuff(cap int) *DChFsFold {
	d := new(DChFsFold)
	d.dat = make(chan *fs.FsFold, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsFold is the send function - aka "MyKind <- some FsFold"
func (c *DChFsFold) ProvideFsFold(dat *fs.FsFold) {
	<-c.req
	c.dat <- dat
}

// RequestFsFold is the receive function - aka "some FsFold <- MyKind"
func (c *DChFsFold) RequestFsFold() (dat *fs.FsFold) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsFold is the comma-ok multi-valued form of RequestFsFold and
// reports whether a received value was sent before the FsFold channel was closed.
func (c *DChFsFold) TryFsFold() (dat *fs.FsFold, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChFsFold is a supply channel
type SChFsFold struct {
	dat chan *fs.FsFold
	// req chan struct{}
}

// MakeSupplyFsFoldChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsFoldChan() *SChFsFold {
	d := new(SChFsFold)
	d.dat = make(chan *fs.FsFold)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsFoldBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsFoldBuff(cap int) *SChFsFold {
	d := new(SChFsFold)
	d.dat = make(chan *fs.FsFold, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsFold is the send function - aka "MyKind <- some FsFold"
func (c *SChFsFold) ProvideFsFold(dat *fs.FsFold) {
	// .req
	c.dat <- dat
}

// RequestFsFold is the receive function - aka "some FsFold <- MyKind"
func (c *SChFsFold) RequestFsFold() (dat *fs.FsFold) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsFold is the comma-ok multi-valued form of RequestFsFold and
// reports whether a received value was sent before the FsFold channel was closed.
func (c *SChFsFold) TryFsFold() (dat *fs.FsFold, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
