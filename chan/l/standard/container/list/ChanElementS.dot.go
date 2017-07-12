// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ElementSChan interface { // bidirectional channel
	ElementSROnlyChan // aka "<-chan" - receive only
	ElementSSOnlyChan // aka "chan<-" - send only
}

type ElementSROnlyChan interface { // receive-only channel
	RequestElementS() (dat []list.Element)        // the receive function - aka "some-new-ElementS-var := <-MyKind"
	TryElementS() (dat []list.Element, open bool) // the multi-valued comma-ok receive function - aka "some-new-ElementS-var, ok := <-MyKind"
}

type ElementSSOnlyChan interface { // send-only channel
	ProvideElementS(dat []list.Element) // the send function - aka "MyKind <- some ElementS"
}

type DChElementS struct { // demand channel
	dat chan []list.Element
	req chan struct{}
}

func MakeDemandElementSChan() *DChElementS {
	d := new(DChElementS)
	d.dat = make(chan []list.Element)
	d.req = make(chan struct{})
	return d
}

func MakeDemandElementSBuff(cap int) *DChElementS {
	d := new(DChElementS)
	d.dat = make(chan []list.Element, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideElementS is the send function - aka "MyKind <- some ElementS"
func (c *DChElementS) ProvideElementS(dat []list.Element) {
	<-c.req
	c.dat <- dat
}

// RequestElementS is the receive function - aka "some ElementS <- MyKind"
func (c *DChElementS) RequestElementS() (dat []list.Element) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryElementS is the comma-ok multi-valued form of RequestElementS and
// reports whether a received value was sent before the ElementS channel was closed.
func (c *DChElementS) TryElementS() (dat []list.Element, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
