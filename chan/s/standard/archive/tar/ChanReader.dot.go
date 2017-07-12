// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

type ReaderChan interface { // bidirectional channel
	ReaderROnlyChan // aka "<-chan" - receive only
	ReaderSOnlyChan // aka "chan<-" - send only
}

type ReaderROnlyChan interface { // receive-only channel
	RequestReader() (dat *tar.Reader)        // the receive function - aka "some-new-Reader-var := <-MyKind"
	TryReader() (dat *tar.Reader, open bool) // the multi-valued comma-ok receive function - aka "some-new-Reader-var, ok := <-MyKind"
}

type ReaderSOnlyChan interface { // send-only channel
	ProvideReader(dat *tar.Reader) // the send function - aka "MyKind <- some Reader"
}

type SChReader struct { // supply channel
	dat chan *tar.Reader
	// req chan struct{}
}

func MakeSupplyReaderChan() *SChReader {
	d := new(SChReader)
	d.dat = make(chan *tar.Reader)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyReaderBuff(cap int) *SChReader {
	d := new(SChReader)
	d.dat = make(chan *tar.Reader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReader is the send function - aka "MyKind <- some Reader"
func (c *SChReader) ProvideReader(dat *tar.Reader) {
	// .req
	c.dat <- dat
}

// RequestReader is the receive function - aka "some Reader <- MyKind"
func (c *SChReader) RequestReader() (dat *tar.Reader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReader is the comma-ok multi-valued form of RequestReader and
// reports whether a received value was sent before the Reader channel was closed.
func (c *SChReader) TryReader() (dat *tar.Reader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len