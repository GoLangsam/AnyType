// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsFoldChan interface { // bidirectional channel
	FsFoldROnlyChan // aka "<-chan" - receive only
	FsFoldSOnlyChan // aka "chan<-" - send only
}

type FsFoldROnlyChan interface { // receive-only channel
	RequestFsFold() (dat *fs.FsFold)        // the receive function - aka "some-new-FsFold-var := <-MyKind"
	TryFsFold() (dat *fs.FsFold, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFold-var, ok := <-MyKind"
}

type FsFoldSOnlyChan interface { // send-only channel
	ProvideFsFold(dat *fs.FsFold) // the send function - aka "MyKind <- some FsFold"
}

type DChFsFold struct { // demand channel
	dat chan *fs.FsFold
	req chan struct{}
}

func MakeDemandFsFoldChan() *DChFsFold {
	d := new(DChFsFold)
	d.dat = make(chan *fs.FsFold)
	d.req = make(chan struct{})
	return d
}

func MakeDemandFsFoldBuff(cap int) *DChFsFold {
	d := new(DChFsFold)
	d.dat = make(chan *fs.FsFold, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFsFold is the send function - aka "MyKind <- some FsFold"
func (c *DChFsFold) ProvideFsFold(dat *fs.FsFold) {
	<-c.req
	c.dat <- dat
}

// RequestFsFold is the receive function - aka "some FsFold <- MyKind"
func (c *DChFsFold) RequestFsFold() (dat *fs.FsFold) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFsFold is the comma-ok multi-valued form of RequestFsFold and
// reports whether a received value was sent before the FsFold channel was closed.
func (c *DChFsFold) TryFsFold() (dat *fs.FsFold, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len