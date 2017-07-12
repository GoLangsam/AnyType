// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteWriterChan interface { // bidirectional channel
	ByteWriterROnlyChan // aka "<-chan" - receive only
	ByteWriterSOnlyChan // aka "chan<-" - send only
}

type ByteWriterROnlyChan interface { // receive-only channel
	RequestByteWriter() (dat io.ByteWriter)        // the receive function - aka "some-new-ByteWriter-var := <-MyKind"
	TryByteWriter() (dat io.ByteWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteWriter-var, ok := <-MyKind"
}

type ByteWriterSOnlyChan interface { // send-only channel
	ProvideByteWriter(dat io.ByteWriter) // the send function - aka "MyKind <- some ByteWriter"
}

type SChByteWriter struct { // supply channel
	dat chan io.ByteWriter
	// req chan struct{}
}

func MakeSupplyByteWriterChan() *SChByteWriter {
	d := new(SChByteWriter)
	d.dat = make(chan io.ByteWriter)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyByteWriterBuff(cap int) *SChByteWriter {
	d := new(SChByteWriter)
	d.dat = make(chan io.ByteWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByteWriter is the send function - aka "MyKind <- some ByteWriter"
func (c *SChByteWriter) ProvideByteWriter(dat io.ByteWriter) {
	// .req
	c.dat <- dat
}

// RequestByteWriter is the receive function - aka "some ByteWriter <- MyKind"
func (c *SChByteWriter) RequestByteWriter() (dat io.ByteWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByteWriter is the comma-ok multi-valued form of RequestByteWriter and
// reports whether a received value was sent before the ByteWriter channel was closed.
func (c *SChByteWriter) TryByteWriter() (dat io.ByteWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
