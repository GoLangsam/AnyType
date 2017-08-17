// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadWriterChan interface { // bidirectional channel
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

type ReadWriterROnlyChan interface { // receive-only channel
	RequestReadWriter() (dat io.ReadWriter)        // the receive function - aka "some-new-ReadWriter-var := <-MyKind"
	TryReadWriter() (dat io.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriter-var, ok := <-MyKind"
}

type ReadWriterSOnlyChan interface { // send-only channel
	ProvideReadWriter(dat io.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}

type SChReadWriter struct { // supply channel
	dat chan io.ReadWriter
	// req chan struct{}
}

func MakeSupplyReadWriterChan() *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan io.ReadWriter)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyReadWriterBuff(cap int) *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan io.ReadWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *SChReadWriter) ProvideReadWriter(dat io.ReadWriter) {
	// .req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *SChReadWriter) RequestReadWriter() (dat io.ReadWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *SChReadWriter) TryReadWriter() (dat io.ReadWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
