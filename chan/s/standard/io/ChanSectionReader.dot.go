// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type SectionReaderChan interface { // bidirectional channel
	SectionReaderROnlyChan // aka "<-chan" - receive only
	SectionReaderSOnlyChan // aka "chan<-" - send only
}

type SectionReaderROnlyChan interface { // receive-only channel
	RequestSectionReader() (dat *io.SectionReader)        // the receive function - aka "some-new-SectionReader-var := <-MyKind"
	TrySectionReader() (dat *io.SectionReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-SectionReader-var, ok := <-MyKind"
}

type SectionReaderSOnlyChan interface { // send-only channel
	ProvideSectionReader(dat *io.SectionReader) // the send function - aka "MyKind <- some SectionReader"
}

type SChSectionReader struct { // supply channel
	dat chan *io.SectionReader
	// req chan struct{}
}

func MakeSupplySectionReaderChan() *SChSectionReader {
	d := new(SChSectionReader)
	d.dat = make(chan *io.SectionReader)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplySectionReaderBuff(cap int) *SChSectionReader {
	d := new(SChSectionReader)
	d.dat = make(chan *io.SectionReader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSectionReader is the send function - aka "MyKind <- some SectionReader"
func (c *SChSectionReader) ProvideSectionReader(dat *io.SectionReader) {
	// .req
	c.dat <- dat
}

// RequestSectionReader is the receive function - aka "some SectionReader <- MyKind"
func (c *SChSectionReader) RequestSectionReader() (dat *io.SectionReader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySectionReader is the comma-ok multi-valued form of RequestSectionReader and
// reports whether a received value was sent before the SectionReader channel was closed.
func (c *SChSectionReader) TrySectionReader() (dat *io.SectionReader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
