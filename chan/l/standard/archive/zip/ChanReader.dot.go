// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

type ReaderChan interface { // bidirectional channel
	ReaderROnlyChan // aka "<-chan" - receive only
	ReaderSOnlyChan // aka "chan<-" - send only
}

type ReaderROnlyChan interface { // receive-only channel
	RequestReader() (dat zip.Reader)        // the receive function - aka "some-new-Reader-var := <-MyKind"
	TryReader() (dat zip.Reader, open bool) // the multi-valued comma-ok receive function - aka "some-new-Reader-var, ok := <-MyKind"
}

type ReaderSOnlyChan interface { // send-only channel
	ProvideReader(dat zip.Reader) // the send function - aka "MyKind <- some Reader"
}

type DChReader struct { // demand channel
	dat chan zip.Reader
	req chan struct{}
}

func MakeDemandReaderChan() *DChReader {
	d := new(DChReader)
	d.dat = make(chan zip.Reader)
	d.req = make(chan struct{})
	return d
}

func MakeDemandReaderBuff(cap int) *DChReader {
	d := new(DChReader)
	d.dat = make(chan zip.Reader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReader is the send function - aka "MyKind <- some Reader"
func (c *DChReader) ProvideReader(dat zip.Reader) {
	<-c.req
	c.dat <- dat
}

// RequestReader is the receive function - aka "some Reader <- MyKind"
func (c *DChReader) RequestReader() (dat zip.Reader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReader is the comma-ok multi-valued form of RequestReader and
// reports whether a received value was sent before the Reader channel was closed.
func (c *DChReader) TryReader() (dat zip.Reader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
