// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type PatternChan interface { // bidirectional channel
	PatternROnlyChan // aka "<-chan" - receive only
	PatternSOnlyChan // aka "chan<-" - send only
}

type PatternROnlyChan interface { // receive-only channel
	RequestPattern() (dat *fs.Pattern)        // the receive function - aka "some-new-Pattern-var := <-MyKind"
	TryPattern() (dat *fs.Pattern, open bool) // the multi-valued comma-ok receive function - aka "some-new-Pattern-var, ok := <-MyKind"
}

type PatternSOnlyChan interface { // send-only channel
	ProvidePattern(dat *fs.Pattern) // the send function - aka "MyKind <- some Pattern"
}

type SChPattern struct { // supply channel
	dat chan *fs.Pattern
	// req chan struct{}
}

func MakeSupplyPatternChan() *SChPattern {
	d := new(SChPattern)
	d.dat = make(chan *fs.Pattern)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyPatternBuff(cap int) *SChPattern {
	d := new(SChPattern)
	d.dat = make(chan *fs.Pattern, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvidePattern is the send function - aka "MyKind <- some Pattern"
func (c *SChPattern) ProvidePattern(dat *fs.Pattern) {
	// .req
	c.dat <- dat
}

// RequestPattern is the receive function - aka "some Pattern <- MyKind"
func (c *SChPattern) RequestPattern() (dat *fs.Pattern) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryPattern is the comma-ok multi-valued form of RequestPattern and
// reports whether a received value was sent before the Pattern channel was closed.
func (c *SChPattern) TryPattern() (dat *fs.Pattern, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
