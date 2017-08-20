// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

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

// DChFsFile is a demand channel
type DChFsFile struct {
	dat chan *fs.FsFile
	req chan struct{}
}

// MakeDemandFsFileChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFsFileChan() *DChFsFile {
	d := new(DChFsFile)
	d.dat = make(chan *fs.FsFile)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFsFileBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFsFileBuff(cap int) *DChFsFile {
	d := new(DChFsFile)
	d.dat = make(chan *fs.FsFile, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsFile is the send function - aka "MyKind <- some FsFile"
func (c *DChFsFile) ProvideFsFile(dat *fs.FsFile) {
	<-c.req
	c.dat <- dat
}

// RequestFsFile is the receive function - aka "some FsFile <- MyKind"
func (c *DChFsFile) RequestFsFile() (dat *fs.FsFile) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsFile is the comma-ok multi-valued form of RequestFsFile and
// reports whether a received value was sent before the FsFile channel was closed.
func (c *DChFsFile) TryFsFile() (dat *fs.FsFile, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
