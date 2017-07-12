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

type SChFsDataS struct { // supply channel
	dat chan fs.FsDataS
	// req chan struct{}
}

func MakeSupplyFsDataSChan() *SChFsDataS {
	d := new(SChFsDataS)
	d.dat = make(chan fs.FsDataS)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyFsDataSBuff(cap int) *SChFsDataS {
	d := new(SChFsDataS)
	d.dat = make(chan fs.FsDataS, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsDataS is the send function - aka "MyKind <- some FsDataS"
func (c *SChFsDataS) ProvideFsDataS(dat fs.FsDataS) {
	// .req
	c.dat <- dat
}

// RequestFsDataS is the receive function - aka "some FsDataS <- MyKind"
func (c *SChFsDataS) RequestFsDataS() (dat fs.FsDataS) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsDataS is the comma-ok multi-valued form of RequestFsDataS and
// reports whether a received value was sent before the FsDataS channel was closed.
func (c *SChFsDataS) TryFsDataS() (dat fs.FsDataS, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
