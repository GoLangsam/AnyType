// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsDataSChan interface { // bidirectional channel
	FsDataSROnlyChan // aka "<-chan" - receive only
	FsDataSSOnlyChan // aka "chan<-" - send only
}

type FsDataSROnlyChan interface { // receive-only channel
	RequestFsDataS() (dat fs.FsDataS)        // the receive function - aka "some-new-FsDataS-var := <-MyKind"
	TryFsDataS() (dat fs.FsDataS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsDataS-var, ok := <-MyKind"
}

type FsDataSSOnlyChan interface { // send-only channel
	ProvideFsDataS(dat fs.FsDataS) // the send function - aka "MyKind <- some FsDataS"
}

type DChFsDataS struct { // demand channel
	dat chan fs.FsDataS
	req chan struct{}
}

func MakeDemandFsDataSChan() *DChFsDataS {
	d := new(DChFsDataS)
	d.dat = make(chan fs.FsDataS)
	d.req = make(chan struct{})
	return d
}

func MakeDemandFsDataSBuff(cap int) *DChFsDataS {
	d := new(DChFsDataS)
	d.dat = make(chan fs.FsDataS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsDataS is the send function - aka "MyKind <- some FsDataS"
func (c *DChFsDataS) ProvideFsDataS(dat fs.FsDataS) {
	<-c.req
	c.dat <- dat
}

// RequestFsDataS is the receive function - aka "some FsDataS <- MyKind"
func (c *DChFsDataS) RequestFsDataS() (dat fs.FsDataS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsDataS is the comma-ok multi-valued form of RequestFsDataS and
// reports whether a received value was sent before the FsDataS channel was closed.
func (c *DChFsDataS) TryFsDataS() (dat fs.FsDataS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
