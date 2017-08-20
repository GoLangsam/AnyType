// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFileSChan represents a
// bidirectional
// channel
type FsFileSChan interface {
	FsFileSROnlyChan // aka "<-chan" - receive only
	FsFileSSOnlyChan // aka "chan<-" - send only
}

// FsFileSROnlyChan represents a
// receive-only
// channel
type FsFileSROnlyChan interface {
	RequestFsFileS() (dat fs.FsFileS)        // the receive function - aka "MyFsFileS := <-MyFsFileSROnlyChan"
	TryFsFileS() (dat fs.FsFileS, open bool) // the multi-valued comma-ok receive function - aka "MyFsFileS, ok := <-MyFsFileSROnlyChan"
}

// FsFileSSOnlyChan represents a
// send-only
// channel
type FsFileSSOnlyChan interface {
	ProvideFsFileS(dat fs.FsFileS) // the send function - aka "MyKind <- some FsFileS"
}

// DChFsFileS is a demand channel
type DChFsFileS struct {
	dat chan fs.FsFileS
	req chan struct{}
}

// MakeDemandFsFileSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsFileSChan() *DChFsFileS {
	d := new(DChFsFileS)
	d.dat = make(chan fs.FsFileS)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsFileSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsFileSBuff(cap int) *DChFsFileS {
	d := new(DChFsFileS)
	d.dat = make(chan fs.FsFileS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsFileS is the send function - aka "MyKind <- some FsFileS"
func (c *DChFsFileS) ProvideFsFileS(dat fs.FsFileS) {
	<-c.req
	c.dat <- dat
}

// RequestFsFileS is the receive function - aka "some FsFileS <- MyKind"
func (c *DChFsFileS) RequestFsFileS() (dat fs.FsFileS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsFileS is the comma-ok multi-valued form of RequestFsFileS and
// reports whether a received value was sent before the FsFileS channel was closed.
func (c *DChFsFileS) TryFsFileS() (dat fs.FsFileS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChFsFileS is a supply channel
type SChFsFileS struct {
	dat chan fs.FsFileS
	// req chan struct{}
}

// MakeSupplyFsFileSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsFileSChan() *SChFsFileS {
	d := new(SChFsFileS)
	d.dat = make(chan fs.FsFileS)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsFileSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsFileSBuff(cap int) *SChFsFileS {
	d := new(SChFsFileS)
	d.dat = make(chan fs.FsFileS, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsFileS is the send function - aka "MyKind <- some FsFileS"
func (c *SChFsFileS) ProvideFsFileS(dat fs.FsFileS) {
	// .req
	c.dat <- dat
}

// RequestFsFileS is the receive function - aka "some FsFileS <- MyKind"
func (c *SChFsFileS) RequestFsFileS() (dat fs.FsFileS) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsFileS is the comma-ok multi-valued form of RequestFsFileS and
// reports whether a received value was sent before the FsFileS channel was closed.
func (c *SChFsFileS) TryFsFileS() (dat fs.FsFileS, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
