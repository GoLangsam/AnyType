// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsInfoSChan interface { // bidirectional channel
	FsInfoSROnlyChan // aka "<-chan" - receive only
	FsInfoSSOnlyChan // aka "chan<-" - send only
}

type FsInfoSROnlyChan interface { // receive-only channel
	RequestFsInfoS() (dat fs.FsInfoS)        // the receive function - aka "some-new-FsInfoS-var := <-MyKind"
	TryFsInfoS() (dat fs.FsInfoS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsInfoS-var, ok := <-MyKind"
}

type FsInfoSSOnlyChan interface { // send-only channel
	ProvideFsInfoS(dat fs.FsInfoS) // the send function - aka "MyKind <- some FsInfoS"
}

type DChFsInfoS struct { // demand channel
	dat chan fs.FsInfoS
	req chan struct{}
}

func MakeDemandFsInfoSChan() *DChFsInfoS {
	d := new(DChFsInfoS)
	d.dat = make(chan fs.FsInfoS)
	d.req = make(chan struct{})
	return d
}

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
