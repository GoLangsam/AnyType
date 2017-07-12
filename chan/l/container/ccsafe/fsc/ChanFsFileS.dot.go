// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsFileSChan interface { // bidirectional channel
	FsFileSROnlyChan // aka "<-chan" - receive only
	FsFileSSOnlyChan // aka "chan<-" - send only
}

type FsFileSROnlyChan interface { // receive-only channel
	RequestFsFileS() (dat fs.FsFileS)        // the receive function - aka "some-new-FsFileS-var := <-MyKind"
	TryFsFileS() (dat fs.FsFileS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFileS-var, ok := <-MyKind"
}

type FsFileSSOnlyChan interface { // send-only channel
	ProvideFsFileS(dat fs.FsFileS) // the send function - aka "MyKind <- some FsFileS"
}

type DChFsFileS struct { // demand channel
	dat chan fs.FsFileS
	req chan struct{}
}

func MakeDemandFsFileSChan() *DChFsFileS {
	d := new(DChFsFileS)
	d.dat = make(chan fs.FsFileS)
	d.req = make(chan struct{})
	return d
}

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
