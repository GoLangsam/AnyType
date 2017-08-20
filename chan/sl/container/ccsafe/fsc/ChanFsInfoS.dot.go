// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsInfoSChan represents a
// bidirectional
// channel
type FsInfoSChan interface {
	FsInfoSROnlyChan // aka "<-chan" - receive only
	FsInfoSSOnlyChan // aka "chan<-" - send only
}

// FsInfoSROnlyChan represents a
// receive-only
// channel
type FsInfoSROnlyChan interface {
	RequestFsInfoS() (dat fs.FsInfoS)        // the receive function - aka "MyFsInfoS := <-MyFsInfoSROnlyChan"
	TryFsInfoS() (dat fs.FsInfoS, open bool) // the multi-valued comma-ok receive function - aka "MyFsInfoS, ok := <-MyFsInfoSROnlyChan"
}

// FsInfoSSOnlyChan represents a
// send-only
// channel
type FsInfoSSOnlyChan interface {
	ProvideFsInfoS(dat fs.FsInfoS) // the send function - aka "MyKind <- some FsInfoS"
}

// DChFsInfoS is a demand channel
type DChFsInfoS struct {
	dat chan fs.FsInfoS
	req chan struct{}
}

// MakeDemandFsInfoSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsInfoSChan() *DChFsInfoS {
	d := new(DChFsInfoS)
	d.dat = make(chan fs.FsInfoS)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsInfoSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsInfoSBuff(cap int) *DChFsInfoS {
	d := new(DChFsInfoS)
	d.dat = make(chan fs.FsInfoS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsInfoS is the send function - aka "MyKind <- some FsInfoS"
func (c *DChFsInfoS) ProvideFsInfoS(dat fs.FsInfoS) {
	<-c.req
	c.dat <- dat
}

// RequestFsInfoS is the receive function - aka "some FsInfoS <- MyKind"
func (c *DChFsInfoS) RequestFsInfoS() (dat fs.FsInfoS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsInfoS is the comma-ok multi-valued form of RequestFsInfoS and
// reports whether a received value was sent before the FsInfoS channel was closed.
func (c *DChFsInfoS) TryFsInfoS() (dat fs.FsInfoS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChFsInfoS is a supply channel
type SChFsInfoS struct {
	dat chan fs.FsInfoS
	// req chan struct{}
}

// MakeSupplyFsInfoSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsInfoSChan() *SChFsInfoS {
	d := new(SChFsInfoS)
	d.dat = make(chan fs.FsInfoS)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsInfoSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsInfoSBuff(cap int) *SChFsInfoS {
	d := new(SChFsInfoS)
	d.dat = make(chan fs.FsInfoS, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsInfoS is the send function - aka "MyKind <- some FsInfoS"
func (c *SChFsInfoS) ProvideFsInfoS(dat fs.FsInfoS) {
	// .req
	c.dat <- dat
}

// RequestFsInfoS is the receive function - aka "some FsInfoS <- MyKind"
func (c *SChFsInfoS) RequestFsInfoS() (dat fs.FsInfoS) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsInfoS is the comma-ok multi-valued form of RequestFsInfoS and
// reports whether a received value was sent before the FsInfoS channel was closed.
func (c *SChFsInfoS) TryFsInfoS() (dat fs.FsInfoS, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
