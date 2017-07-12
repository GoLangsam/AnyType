// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
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

type DChPattern struct { // demand channel
	dat chan *fs.Pattern
	req chan struct{}
}

func MakeDemandPatternChan() *DChPattern {
	d := new(DChPattern)
	d.dat = make(chan *fs.Pattern)
	d.req = make(chan struct{})
	return d
}

func MakeDemandPatternBuff(cap int) *DChPattern {
	d := new(DChPattern)
	d.dat = make(chan *fs.Pattern, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePattern is the send function - aka "MyKind <- some Pattern"
func (c *DChPattern) ProvidePattern(dat *fs.Pattern) {
	<-c.req
	c.dat <- dat
}

// RequestPattern is the receive function - aka "some Pattern <- MyKind"
func (c *DChPattern) RequestPattern() (dat *fs.Pattern) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPattern is the comma-ok multi-valued form of RequestPattern and
// reports whether a received value was sent before the Pattern channel was closed.
func (c *DChPattern) TryPattern() (dat *fs.Pattern, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
