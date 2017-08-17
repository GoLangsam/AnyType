// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
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

type SChPatternS struct { // supply channel
	dat chan fs.PatternS
	// req chan struct{}
}

func MakeSupplyPatternSChan() *SChPatternS {
	d := new(SChPatternS)
	d.dat = make(chan fs.PatternS)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyPatternSBuff(cap int) *SChPatternS {
	d := new(SChPatternS)
	d.dat = make(chan fs.PatternS, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvidePatternS is the send function - aka "MyKind <- some PatternS"
func (c *SChPatternS) ProvidePatternS(dat fs.PatternS) {
	// .req
	c.dat <- dat
}

// RequestPatternS is the receive function - aka "some PatternS <- MyKind"
func (c *SChPatternS) RequestPatternS() (dat fs.PatternS) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryPatternS is the comma-ok multi-valued form of RequestPatternS and
// reports whether a received value was sent before the PatternS channel was closed.
func (c *SChPatternS) TryPatternS() (dat fs.PatternS, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
