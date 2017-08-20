// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsPathSChan represents a
// bidirectional
// channel
type FsPathSChan interface {
	FsPathSROnlyChan // aka "<-chan" - receive only
	FsPathSSOnlyChan // aka "chan<-" - send only
}

// FsPathSROnlyChan represents a
// receive-only
// channel
type FsPathSROnlyChan interface {
	RequestFsPathS() (dat fs.FsPathS)        // the receive function - aka "MyFsPathS := <-MyFsPathSROnlyChan"
	TryFsPathS() (dat fs.FsPathS, open bool) // the multi-valued comma-ok receive function - aka "MyFsPathS, ok := <-MyFsPathSROnlyChan"
}

// FsPathSSOnlyChan represents a
// send-only
// channel
type FsPathSSOnlyChan interface {
	ProvideFsPathS(dat fs.FsPathS) // the send function - aka "MyKind <- some FsPathS"
}

// DChFsPathS is a demand channel
type DChFsPathS struct {
	dat chan fs.FsPathS
	req chan struct{}
}

// MakeDemandFsPathSChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsPathSChan() *DChFsPathS {
	d := new(DChFsPathS)
	d.dat = make(chan fs.FsPathS)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsPathSBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsPathSBuff(cap int) *DChFsPathS {
	d := new(DChFsPathS)
	d.dat = make(chan fs.FsPathS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsPathS is the send function - aka "MyKind <- some FsPathS"
func (c *DChFsPathS) ProvideFsPathS(dat fs.FsPathS) {
	<-c.req
	c.dat <- dat
}

// RequestFsPathS is the receive function - aka "some FsPathS <- MyKind"
func (c *DChFsPathS) RequestFsPathS() (dat fs.FsPathS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsPathS is the comma-ok multi-valued form of RequestFsPathS and
// reports whether a received value was sent before the FsPathS channel was closed.
func (c *DChFsPathS) TryFsPathS() (dat fs.FsPathS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
