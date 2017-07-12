// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ElementChan interface { // bidirectional channel
	ElementROnlyChan // aka "<-chan" - receive only
	ElementSOnlyChan // aka "chan<-" - send only
}

type ElementROnlyChan interface { // receive-only channel
	RequestElement() (dat list.Element)        // the receive function - aka "some-new-Element-var := <-MyKind"
	TryElement() (dat list.Element, open bool) // the multi-valued comma-ok receive function - aka "some-new-Element-var, ok := <-MyKind"
}

type ElementSOnlyChan interface { // send-only channel
	ProvideElement(dat list.Element) // the send function - aka "MyKind <- some Element"
}

type SChElement struct { // supply channel
	dat chan list.Element
	// req chan struct{}
}

func MakeSupplyElementChan() *SChElement {
	d := new(SChElement)
	d.dat = make(chan list.Element)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyElementBuff(cap int) *SChElement {
	d := new(SChElement)
	d.dat = make(chan list.Element, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideElement is the send function - aka "MyKind <- some Element"
func (c *SChElement) ProvideElement(dat list.Element) {
	// .req
	c.dat <- dat
}

// RequestElement is the receive function - aka "some Element <- MyKind"
func (c *SChElement) RequestElement() (dat list.Element) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryElement is the comma-ok multi-valued form of RequestElement and
// reports whether a received value was sent before the Element channel was closed.
func (c *SChElement) TryElement() (dat list.Element, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len