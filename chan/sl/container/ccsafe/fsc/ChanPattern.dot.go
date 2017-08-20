// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// PatternChan represents a
// bidirectional
// channel
type PatternChan interface {
	PatternROnlyChan // aka "<-chan" - receive only
	PatternSOnlyChan // aka "chan<-" - send only
}

// PatternROnlyChan represents a
// receive-only
// channel
type PatternROnlyChan interface {
	RequestPattern() (dat *fs.Pattern)        // the receive function - aka "MyPattern := <-MyPatternROnlyChan"
	TryPattern() (dat *fs.Pattern, open bool) // the multi-valued comma-ok receive function - aka "MyPattern, ok := <-MyPatternROnlyChan"
}

// PatternSOnlyChan represents a
// send-only
// channel
type PatternSOnlyChan interface {
	ProvidePattern(dat *fs.Pattern) // the send function - aka "MyKind <- some Pattern"
}

// DChPattern is a demand channel
type DChPattern struct {
	dat chan *fs.Pattern
	req chan struct{}
}

// MakeDemandPatternChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandPatternChan() *DChPattern {
	d := new(DChPattern)
	d.dat = make(chan *fs.Pattern)
	d.req = make(chan struct{})
	return d
}

// MakeDemandPatternBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
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

// DChPattern is a supply channel
type SChPattern struct {
	dat chan *fs.Pattern
	// req chan struct{}
}

// MakeSupplyPatternChan() returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyPatternChan() *SChPattern {
	d := new(SChPattern)
	d.dat = make(chan *fs.Pattern)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyPatternBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
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
