// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/dot"
)

type DotChan interface { // bidirectional channel
	DotROnlyChan // aka "<-chan" - receive only
	DotSOnlyChan // aka "chan<-" - send only
}

type DotROnlyChan interface { // receive-only channel
	RequestDot() (dat dot.Dot)        // the receive function - aka "some-new-Dot-var := <-MyKind"
	TryDot() (dat dot.Dot, open bool) // the multi-valued comma-ok receive function - aka "some-new-Dot-var, ok := <-MyKind"
}

type DotSOnlyChan interface { // send-only channel
	ProvideDot(dat dot.Dot) // the send function - aka "MyKind <- some Dot"
}

type DChDot struct { // demand channel
	dat chan dot.Dot
	req chan struct{}
}

func MakeDemandDotChan() *DChDot {
	d := new(DChDot)
	d.dat = make(chan dot.Dot)
	d.req = make(chan struct{})
	return d
}

func MakeDemandDotBuff(cap int) *DChDot {
	d := new(DChDot)
	d.dat = make(chan dot.Dot, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideDot is the send function - aka "MyKind <- some Dot"
func (c *DChDot) ProvideDot(dat dot.Dot) {
	<-c.req
	c.dat <- dat
}

// RequestDot is the receive function - aka "some Dot <- MyKind"
func (c *DChDot) RequestDot() (dat dot.Dot) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryDot is the comma-ok multi-valued form of RequestDot and
// reports whether a received value was sent before the Dot channel was closed.
func (c *DChDot) TryDot() (dat dot.Dot, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
