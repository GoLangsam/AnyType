// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

type FileInfoChan interface { // bidirectional channel
	FileInfoROnlyChan // aka "<-chan" - receive only
	FileInfoSOnlyChan // aka "chan<-" - send only
}

type FileInfoROnlyChan interface { // receive-only channel
	RequestFileInfo() (dat os.FileInfo)        // the receive function - aka "some-new-FileInfo-var := <-MyKind"
	TryFileInfo() (dat os.FileInfo, open bool) // the multi-valued comma-ok receive function - aka "some-new-FileInfo-var, ok := <-MyKind"
}

type FileInfoSOnlyChan interface { // send-only channel
	ProvideFileInfo(dat os.FileInfo) // the send function - aka "MyKind <- some FileInfo"
}

type DChFileInfo struct { // demand channel
	dat chan os.FileInfo
	req chan struct{}
}

func MakeDemandFileInfoChan() *DChFileInfo {
	d := new(DChFileInfo)
	d.dat = make(chan os.FileInfo)
	d.req = make(chan struct{})
	return d
}

func MakeDemandFileInfoBuff(cap int) *DChFileInfo {
	d := new(DChFileInfo)
	d.dat = make(chan os.FileInfo, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFileInfo is the send function - aka "MyKind <- some FileInfo"
func (c *DChFileInfo) ProvideFileInfo(dat os.FileInfo) {
	<-c.req
	c.dat <- dat
}

// RequestFileInfo is the receive function - aka "some FileInfo <- MyKind"
func (c *DChFileInfo) RequestFileInfo() (dat os.FileInfo) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFileInfo is the comma-ok multi-valued form of RequestFileInfo and
// reports whether a received value was sent before the FileInfo channel was closed.
func (c *DChFileInfo) TryFileInfo() (dat os.FileInfo, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len