// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ListSChan interface { // bidirectional channel
	ListSROnlyChan // aka "<-chan" - receive only
	ListSSOnlyChan // aka "chan<-" - send only
}

type ListSROnlyChan interface { // receive-only channel
	RequestListS() (dat []list.List)        // the receive function - aka "some-new-ListS-var := <-MyKind"
	TryListS() (dat []list.List, open bool) // the multi-valued comma-ok receive function - aka "some-new-ListS-var, ok := <-MyKind"
}

type ListSSOnlyChan interface { // send-only channel
	ProvideListS(dat []list.List) // the send function - aka "MyKind <- some ListS"
}

type DChListS struct { // demand channel
	dat chan []list.List
	req chan struct{}
}

func MakeDemandListSChan() *DChListS {
	d := new(DChListS)
	d.dat = make(chan []list.List)
	d.req = make(chan struct{})
	return d
}

func MakeDemandListSBuff(cap int) *DChListS {
	d := new(DChListS)
	d.dat = make(chan []list.List, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideListS is the send function - aka "MyKind <- some ListS"
func (c *DChListS) ProvideListS(dat []list.List) {
	<-c.req
	c.dat <- dat
}

// RequestListS is the receive function - aka "some ListS <- MyKind"
func (c *DChListS) RequestListS() (dat []list.List) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryListS is the comma-ok multi-valued form of RequestListS and
// reports whether a received value was sent before the ListS channel was closed.
func (c *DChListS) TryListS() (dat []list.List, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
