// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

type HeaderChan interface { // bidirectional channel
	HeaderROnlyChan // aka "<-chan" - receive only
	HeaderSOnlyChan // aka "chan<-" - send only
}

type HeaderROnlyChan interface { // receive-only channel
	RequestHeader() (dat *tar.Header)        // the receive function - aka "some-new-Header-var := <-MyKind"
	TryHeader() (dat *tar.Header, open bool) // the multi-valued comma-ok receive function - aka "some-new-Header-var, ok := <-MyKind"
}

type HeaderSOnlyChan interface { // send-only channel
	ProvideHeader(dat *tar.Header) // the send function - aka "MyKind <- some Header"
}

type DChHeader struct { // demand channel
	dat chan *tar.Header
	req chan struct{}
}

func MakeDemandHeaderChan() *DChHeader {
	d := new(DChHeader)
	d.dat = make(chan *tar.Header)
	d.req = make(chan struct{})
	return d
}

func MakeDemandHeaderBuff(cap int) *DChHeader {
	d := new(DChHeader)
	d.dat = make(chan *tar.Header, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideHeader is the send function - aka "MyKind <- some Header"
func (c *DChHeader) ProvideHeader(dat *tar.Header) {
	<-c.req
	c.dat <- dat
}

// RequestHeader is the receive function - aka "some Header <- MyKind"
func (c *DChHeader) RequestHeader() (dat *tar.Header) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryHeader is the comma-ok multi-valued form of RequestHeader and
// reports whether a received value was sent before the Header channel was closed.
func (c *DChHeader) TryHeader() (dat *tar.Header, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
