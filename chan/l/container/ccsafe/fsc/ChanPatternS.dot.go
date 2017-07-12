// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type PatternSChan interface { // bidirectional channel
	PatternSROnlyChan // aka "<-chan" - receive only
	PatternSSOnlyChan // aka "chan<-" - send only
}

type PatternSROnlyChan interface { // receive-only channel
	RequestPatternS() (dat fs.PatternS)        // the receive function - aka "some-new-PatternS-var := <-MyKind"
	TryPatternS() (dat fs.PatternS, open bool) // the multi-valued comma-ok receive function - aka "some-new-PatternS-var, ok := <-MyKind"
}

type PatternSSOnlyChan interface { // send-only channel
	ProvidePatternS(dat fs.PatternS) // the send function - aka "MyKind <- some PatternS"
}

type DChPatternS struct { // demand channel
	dat chan fs.PatternS
	req chan struct{}
}

func MakeDemandPatternSChan() *DChPatternS {
	d := new(DChPatternS)
	d.dat = make(chan fs.PatternS)
	d.req = make(chan struct{})
	return d
}

func MakeDemandPatternSBuff(cap int) *DChPatternS {
	d := new(DChPatternS)
	d.dat = make(chan fs.PatternS, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePatternS is the send function - aka "MyKind <- some PatternS"
func (c *DChPatternS) ProvidePatternS(dat fs.PatternS) {
	<-c.req
	c.dat <- dat
}

// RequestPatternS is the receive function - aka "some PatternS <- MyKind"
func (c *DChPatternS) RequestPatternS() (dat fs.PatternS) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPatternS is the comma-ok multi-valued form of RequestPatternS and
// reports whether a received value was sent before the PatternS channel was closed.
func (c *DChPatternS) TryPatternS() (dat fs.PatternS, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
