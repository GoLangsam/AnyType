// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterChan interface { // bidirectional channel
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

type WriterROnlyChan interface { // receive-only channel
	RequestWriter() (dat io.Writer)        // the receive function - aka "some-new-Writer-var := <-MyKind"
	TryWriter() (dat io.Writer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Writer-var, ok := <-MyKind"
}

type WriterSOnlyChan interface { // send-only channel
	ProvideWriter(dat io.Writer) // the send function - aka "MyKind <- some Writer"
}

type SChWriter struct { // supply channel
	dat chan io.Writer
	// req chan struct{}
}

func MakeSupplyWriterChan() *SChWriter {
	d := new(SChWriter)
	d.dat = make(chan io.Writer)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyWriterBuff(cap int) *SChWriter {
	d := new(SChWriter)
	d.dat = make(chan io.Writer, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriter is the send function - aka "MyKind <- some Writer"
func (c *SChWriter) ProvideWriter(dat io.Writer) {
	// .req
	c.dat <- dat
}

// RequestWriter is the receive function - aka "some Writer <- MyKind"
func (c *SChWriter) RequestWriter() (dat io.Writer) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriter is the comma-ok multi-valued form of RequestWriter and
// reports whether a received value was sent before the Writer channel was closed.
func (c *SChWriter) TryWriter() (dat io.Writer, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
