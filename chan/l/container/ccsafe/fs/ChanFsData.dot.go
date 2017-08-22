// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsDataChan represents a
// bidirectional
// channel
type FsDataChan interface {
	FsDataROnlyChan // aka "<-chan" - receive only
	FsDataSOnlyChan // aka "chan<-" - send only
}

// FsDataROnlyChan represents a
// receive-only
// channel
type FsDataROnlyChan interface {
	RequestFsData() (dat *fs.FsData)        // the receive function - aka "MyFsData := <-MyFsDataROnlyChan"
	TryFsData() (dat *fs.FsData, open bool) // the multi-valued comma-ok receive function - aka "MyFsData, ok := <-MyFsDataROnlyChan"
}

// FsDataSOnlyChan represents a
// send-only
// channel
type FsDataSOnlyChan interface {
	ProvideFsData(dat *fs.FsData) // the send function - aka "MyKind <- some FsData"
}

// DChFsData is a demand channel
type DChFsData struct {
	dat chan *fs.FsData
	req chan struct{}
}

// MakeDemandFsDataChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsDataChan() *DChFsData {
	d := new(DChFsData)
	d.dat = make(chan *fs.FsData)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsDataBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsDataBuff(cap int) *DChFsData {
	d := new(DChFsData)
	d.dat = make(chan *fs.FsData, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsData is the send function - aka "MyKind <- some FsData"
func (c *DChFsData) ProvideFsData(dat *fs.FsData) {
	<-c.req
	c.dat <- dat
}

// RequestFsData is the receive function - aka "some FsData <- MyKind"
func (c *DChFsData) RequestFsData() (dat *fs.FsData) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsData is the comma-ok multi-valued form of RequestFsData and
// reports whether a received value was sent before the FsData channel was closed.
func (c *DChFsData) TryFsData() (dat *fs.FsData, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
