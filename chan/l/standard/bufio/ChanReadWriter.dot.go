// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type ReadWriterChan interface { // bidirectional channel
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

type ReadWriterROnlyChan interface { // receive-only channel
	RequestReadWriter() (dat *bufio.ReadWriter)        // the receive function - aka "some-new-ReadWriter-var := <-MyKind"
	TryReadWriter() (dat *bufio.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadWriter-var, ok := <-MyKind"
}

type ReadWriterSOnlyChan interface { // send-only channel
	ProvideReadWriter(dat *bufio.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}

type DChReadWriter struct { // demand channel
	dat chan *bufio.ReadWriter
	req chan struct{}
}

func MakeDemandReadWriterChan() *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan *bufio.ReadWriter)
	d.req = make(chan struct{})
	return d
}

func MakeDemandReadWriterBuff(cap int) *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan *bufio.ReadWriter, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *DChReadWriter) ProvideReadWriter(dat *bufio.ReadWriter) {
	<-c.req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *DChReadWriter) RequestReadWriter() (dat *bufio.ReadWriter) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *DChReadWriter) TryReadWriter() (dat *bufio.ReadWriter, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
