// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteReaderChan interface { // bidirectional channel
	ByteReaderROnlyChan // aka "<-chan" - receive only
	ByteReaderSOnlyChan // aka "chan<-" - send only
}

type ByteReaderROnlyChan interface { // receive-only channel
	RequestByteReader() (dat io.ByteReader)        // the receive function - aka "some-new-ByteReader-var := <-MyKind"
	TryByteReader() (dat io.ByteReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteReader-var, ok := <-MyKind"
}

type ByteReaderSOnlyChan interface { // send-only channel
	ProvideByteReader(dat io.ByteReader) // the send function - aka "MyKind <- some ByteReader"
}

type DChByteReader struct { // demand channel
	dat chan io.ByteReader
	req chan struct{}
}

func MakeDemandByteReaderChan() *DChByteReader {
	d := new(DChByteReader)
	d.dat = make(chan io.ByteReader)
	d.req = make(chan struct{})
	return d
}

func MakeDemandByteReaderBuff(cap int) *DChByteReader {
	d := new(DChByteReader)
	d.dat = make(chan io.ByteReader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideByteReader is the send function - aka "MyKind <- some ByteReader"
func (c *DChByteReader) ProvideByteReader(dat io.ByteReader) {
	<-c.req
	c.dat <- dat
}

// RequestByteReader is the receive function - aka "some ByteReader <- MyKind"
func (c *DChByteReader) RequestByteReader() (dat io.ByteReader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryByteReader is the comma-ok multi-valued form of RequestByteReader and
// reports whether a received value was sent before the ByteReader channel was closed.
func (c *DChByteReader) TryByteReader() (dat io.ByteReader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
