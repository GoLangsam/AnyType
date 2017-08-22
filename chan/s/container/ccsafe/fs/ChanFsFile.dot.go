// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFileChan represents a
// bidirectional
// channel
type FsFileChan interface {
	FsFileROnlyChan // aka "<-chan" - receive only
	FsFileSOnlyChan // aka "chan<-" - send only
}

// FsFileROnlyChan represents a
// receive-only
// channel
type FsFileROnlyChan interface {
	RequestFsFile() (dat *fs.FsFile)        // the receive function - aka "MyFsFile := <-MyFsFileROnlyChan"
	TryFsFile() (dat *fs.FsFile, open bool) // the multi-valued comma-ok receive function - aka "MyFsFile, ok := <-MyFsFileROnlyChan"
}

// FsFileSOnlyChan represents a
// send-only
// channel
type FsFileSOnlyChan interface {
	ProvideFsFile(dat *fs.FsFile) // the send function - aka "MyKind <- some FsFile"
}

// SChFsFile is a supply channel
type SChFsFile struct {
	dat chan *fs.FsFile
	// req chan struct{}
}

// MakeSupplyFsFileChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyFsFileChan() *SChFsFile {
	d := new(SChFsFile)
	d.dat = make(chan *fs.FsFile)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyFsFileBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyFsFileBuff(cap int) *SChFsFile {
	d := new(SChFsFile)
	d.dat = make(chan *fs.FsFile, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFsFile is the send function - aka "MyKind <- some FsFile"
func (c *SChFsFile) ProvideFsFile(dat *fs.FsFile) {
	// .req
	c.dat <- dat
}

// RequestFsFile is the receive function - aka "some FsFile <- MyKind"
func (c *SChFsFile) RequestFsFile() (dat *fs.FsFile) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFsFile is the comma-ok multi-valued form of RequestFsFile and
// reports whether a received value was sent before the FsFile channel was closed.
func (c *SChFsFile) TryFsFile() (dat *fs.FsFile, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
