// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/dotpath"
)

type DotPathChan interface { // bidirectional channel
	DotPathROnlyChan // aka "<-chan" - receive only
	DotPathSOnlyChan // aka "chan<-" - send only
}

type DotPathROnlyChan interface { // receive-only channel
	RequestDotPath() (dat dotpath.DotPath)        // the receive function - aka "some-new-DotPath-var := <-MyKind"
	TryDotPath() (dat dotpath.DotPath, open bool) // the multi-valued comma-ok receive function - aka "some-new-DotPath-var, ok := <-MyKind"
}

type DotPathSOnlyChan interface { // send-only channel
	ProvideDotPath(dat dotpath.DotPath) // the send function - aka "MyKind <- some DotPath"
}

type SChDotPath struct { // supply channel
	dat chan dotpath.DotPath
	// req chan struct{}
}

func MakeSupplyDotPathChan() *SChDotPath {
	d := new(SChDotPath)
	d.dat = make(chan dotpath.DotPath)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyDotPathBuff(cap int) *SChDotPath {
	d := new(SChDotPath)
	d.dat = make(chan dotpath.DotPath, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideDotPath is the send function - aka "MyKind <- some DotPath"
func (c *SChDotPath) ProvideDotPath(dat dotpath.DotPath) {
	// .req
	c.dat <- dat
}

// RequestDotPath is the receive function - aka "some DotPath <- MyKind"
func (c *SChDotPath) RequestDotPath() (dat dotpath.DotPath) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryDotPath is the comma-ok multi-valued form of RequestDotPath and
// reports whether a received value was sent before the DotPath channel was closed.
func (c *SChDotPath) TryDotPath() (dat dotpath.DotPath, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
