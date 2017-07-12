// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsDataChan interface { // bidirectional channel
	FsDataROnlyChan // aka "<-chan" - receive only
	FsDataSOnlyChan // aka "chan<-" - send only
}

type FsDataROnlyChan interface { // receive-only channel
	RequestFsData() (dat *fs.FsData)        // the receive function - aka "some-new-FsData-var := <-MyKind"
	TryFsData() (dat *fs.FsData, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsData-var, ok := <-MyKind"
}

type FsDataSOnlyChan interface { // send-only channel
	ProvideFsData(dat *fs.FsData) // the send function - aka "MyKind <- some FsData"
}

type DChFsData struct { // demand channel
	dat chan *fs.FsData
	req chan struct{}
}

func MakeDemandFsDataChan() *DChFsData {
	d := new(DChFsData)
	d.dat = make(chan *fs.FsData)
	d.req = make(chan struct{})
	return d
}

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
