// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseSChan represents a
// bidirectional
// channel
type FsBaseSChan interface {
	FsBaseSROnlyChan // aka "<-chan" - receive only
	FsBaseSSOnlyChan // aka "chan<-" - send only
}

// FsBaseSROnlyChan represents a
// receive-only
// channel
type FsBaseSROnlyChan interface {
	RequestFsBaseS() (dat fs.FsBaseS)        // the receive function - aka "MyFsBaseS := <-MyFsBaseSROnlyChan"
	TryFsBaseS() (dat fs.FsBaseS, open bool) // the multi-valued comma-ok receive function - aka "MyFsBaseS, ok := <-MyFsBaseSROnlyChan"
}

// FsBaseSSOnlyChan represents a
// send-only
// channel
type FsBaseSSOnlyChan interface {
	ProvideFsBaseS(dat fs.FsBaseS) // the send function - aka "MyKind <- some FsBaseS"
}

// DChFsBaseS is a demand channel
type DChFsBaseS struct {
	dat chan fs.FsBaseS
	req chan struct{}
}

// MakeDemandFsBaseSChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsBaseSChan() *DChFsBaseS {
	d := new(DChFsBaseS)
	d.dat = make(chan fs.FsBaseS)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsBaseSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsBaseSBuff(cap int) *DChFsBaseS {
	d := new(DChFsBaseS)
	d.dat = make(chan fs.FsBaseS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsBaseS is the send function - aka "MyKind <- some FsBaseS"
func (c *DChFsBaseS) ProvideFsBaseS(dat fs.FsBaseS) {
	<-c.req
	c.dat <- dat
}

// RequestFsBaseS is the receive function - aka "some FsBaseS <- MyKind"
func (c *DChFsBaseS) RequestFsBaseS() (dat fs.FsBaseS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsBaseS is the comma-ok multi-valued form of RequestFsBaseS and
// reports whether a received value was sent before the FsBaseS channel was closed.
func (c *DChFsBaseS) TryFsBaseS() (dat fs.FsBaseS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
