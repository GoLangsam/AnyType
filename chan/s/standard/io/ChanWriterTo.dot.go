// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterToChan interface { // bidirectional channel
	WriterToROnlyChan // aka "<-chan" - receive only
	WriterToSOnlyChan // aka "chan<-" - send only
}

type WriterToROnlyChan interface { // receive-only channel
	RequestWriterTo() (dat io.WriterTo)        // the receive function - aka "some-new-WriterTo-var := <-MyKind"
	TryWriterTo() (dat io.WriterTo, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriterTo-var, ok := <-MyKind"
}

type WriterToSOnlyChan interface { // send-only channel
	ProvideWriterTo(dat io.WriterTo) // the send function - aka "MyKind <- some WriterTo"
}

type SChWriterTo struct { // supply channel
	dat chan io.WriterTo
	// req chan struct{}
}

func MakeSupplyWriterToChan() *SChWriterTo {
	d := new(SChWriterTo)
	d.dat = make(chan io.WriterTo)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyWriterToBuff(cap int) *SChWriterTo {
	d := new(SChWriterTo)
	d.dat = make(chan io.WriterTo, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriterTo is the send function - aka "MyKind <- some WriterTo"
func (c *SChWriterTo) ProvideWriterTo(dat io.WriterTo) {
	// .req
	c.dat <- dat
}

// RequestWriterTo is the receive function - aka "some WriterTo <- MyKind"
func (c *SChWriterTo) RequestWriterTo() (dat io.WriterTo) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriterTo is the comma-ok multi-valued form of RequestWriterTo and
// reports whether a received value was sent before the WriterTo channel was closed.
func (c *SChWriterTo) TryWriterTo() (dat io.WriterTo, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
