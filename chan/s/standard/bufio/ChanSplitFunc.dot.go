// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type SplitFuncChan interface { // bidirectional channel
	SplitFuncROnlyChan // aka "<-chan" - receive only
	SplitFuncSOnlyChan // aka "chan<-" - send only
}

type SplitFuncROnlyChan interface { // receive-only channel
	RequestSplitFunc() (dat bufio.SplitFunc)        // the receive function - aka "some-new-SplitFunc-var := <-MyKind"
	TrySplitFunc() (dat bufio.SplitFunc, open bool) // the multi-valued comma-ok receive function - aka "some-new-SplitFunc-var, ok := <-MyKind"
}

type SplitFuncSOnlyChan interface { // send-only channel
	ProvideSplitFunc(dat bufio.SplitFunc) // the send function - aka "MyKind <- some SplitFunc"
}

type SChSplitFunc struct { // supply channel
	dat chan bufio.SplitFunc
	// req chan struct{}
}

func MakeSupplySplitFuncChan() *SChSplitFunc {
	d := new(SChSplitFunc)
	d.dat = make(chan bufio.SplitFunc)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplySplitFuncBuff(cap int) *SChSplitFunc {
	d := new(SChSplitFunc)
	d.dat = make(chan bufio.SplitFunc, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSplitFunc is the send function - aka "MyKind <- some SplitFunc"
func (c *SChSplitFunc) ProvideSplitFunc(dat bufio.SplitFunc) {
	// .req
	c.dat <- dat
}

// RequestSplitFunc is the receive function - aka "some SplitFunc <- MyKind"
func (c *SChSplitFunc) RequestSplitFunc() (dat bufio.SplitFunc) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySplitFunc is the comma-ok multi-valued form of RequestSplitFunc and
// reports whether a received value was sent before the SplitFunc channel was closed.
func (c *SChSplitFunc) TrySplitFunc() (dat bufio.SplitFunc, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
