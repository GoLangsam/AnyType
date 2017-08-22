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

// SChFsData is a supply channel
type SChFsData struct {
	dat chan *fs.FsData
	// req chan struct{}
}

// MakeSupplyFsDataChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsDataChan() *SChFsData {
	d := new(SChFsData)
	d.dat = make(chan *fs.FsData)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsDataBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
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
