// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsDataSChan represents a
// bidirectional
// channel
type FsDataSChan interface {
	FsDataSROnlyChan // aka "<-chan" - receive only
	FsDataSSOnlyChan // aka "chan<-" - send only
}

// FsDataSROnlyChan represents a
// receive-only
// channel
type FsDataSROnlyChan interface {
	RequestFsDataS() (dat fs.FsDataS)        // the receive function - aka "MyFsDataS := <-MyFsDataSROnlyChan"
	TryFsDataS() (dat fs.FsDataS, open bool) // the multi-valued comma-ok receive function - aka "MyFsDataS, ok := <-MyFsDataSROnlyChan"
}

// FsDataSSOnlyChan represents a
// send-only
// channel
type FsDataSSOnlyChan interface {
	ProvideFsDataS(dat fs.FsDataS) // the send function - aka "MyKind <- some FsDataS"
}

// DChFsDataS is a demand channel
type DChFsDataS struct {
	dat chan fs.FsDataS
	req chan struct{}
}

// MakeDemandFsDataSChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsDataSChan() *DChFsDataS {
	d := new(DChFsDataS)
	d.dat = make(chan fs.FsDataS)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsDataSBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
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
