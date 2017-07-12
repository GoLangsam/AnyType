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

type SChFsData struct { // supply channel
	dat chan *fs.FsData
	// req chan struct{}
}

func MakeSupplyFsDataChan() *SChFsData {
	d := new(SChFsData)
	d.dat = make(chan *fs.FsData)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyFsDataBuff(cap int) *SChFsData {
	d := new(SChFsData)
	d.dat = make(chan *fs.FsData, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsData is the send function - aka "MyKind <- some FsData"
func (c *SChFsData) ProvideFsData(dat *fs.FsData) {
	// .req
	c.dat <- dat
}

// RequestFsData is the receive function - aka "some FsData <- MyKind"
func (c *SChFsData) RequestFsData() (dat *fs.FsData) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsData is the comma-ok multi-valued form of RequestFsData and
// reports whether a received value was sent before the FsData channel was closed.
func (c *SChFsData) TryFsData() (dat *fs.FsData, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
