// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseChan represents a
// bidirectional
// channel
type FsBaseChan interface {
	FsBaseROnlyChan // aka "<-chan" - receive only
	FsBaseSOnlyChan // aka "chan<-" - send only
}

// FsBaseROnlyChan represents a
// receive-only
// channel
type FsBaseROnlyChan interface {
	RequestFsBase() (dat *fs.FsBase)        // the receive function - aka "MyFsBase := <-MyFsBaseROnlyChan"
	TryFsBase() (dat *fs.FsBase, open bool) // the multi-valued comma-ok receive function - aka "MyFsBase, ok := <-MyFsBaseROnlyChan"
}

// FsBaseSOnlyChan represents a
// send-only
// channel
type FsBaseSOnlyChan interface {
	ProvideFsBase(dat *fs.FsBase) // the send function - aka "MyKind <- some FsBase"
}

// SChFsBase is a supply channel
type SChFsBase struct {
	dat chan *fs.FsBase
	// req chan struct{}
}

// MakeSupplyFsBaseChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsBaseChan() *SChFsBase {
	d := new(SChFsBase)
	d.dat = make(chan *fs.FsBase)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsBaseBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsBaseBuff(cap int) *SChFsBase {
	d := new(SChFsBase)
	d.dat = make(chan *fs.FsBase, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsBase is the send function - aka "MyKind <- some FsBase"
func (c *SChFsBase) ProvideFsBase(dat *fs.FsBase) {
	// .req
	c.dat <- dat
}

// RequestFsBase is the receive function - aka "some FsBase <- MyKind"
func (c *SChFsBase) RequestFsBase() (dat *fs.FsBase) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsBase is the comma-ok multi-valued form of RequestFsBase and
// reports whether a received value was sent before the FsBase channel was closed.
func (c *SChFsBase) TryFsBase() (dat *fs.FsBase, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
