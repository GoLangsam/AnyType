// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsPathSChan interface { // bidirectional channel
	FsPathSROnlyChan // aka "<-chan" - receive only
	FsPathSSOnlyChan // aka "chan<-" - send only
}

type FsPathSROnlyChan interface { // receive-only channel
	RequestFsPathS() (dat fs.FsPathS)        // the receive function - aka "some-new-FsPathS-var := <-MyKind"
	TryFsPathS() (dat fs.FsPathS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsPathS-var, ok := <-MyKind"
}

type FsPathSSOnlyChan interface { // send-only channel
	ProvideFsPathS(dat fs.FsPathS) // the send function - aka "MyKind <- some FsPathS"
}

type DChFsPathS struct { // demand channel
	dat chan fs.FsPathS
	req chan struct{}
}

func MakeDemandFsPathSChan() *DChFsPathS {
	d := new(DChFsPathS)
	d.dat = make(chan fs.FsPathS)
	d.req = make(chan struct{})
	return d
}

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
