// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type PipeWriterChan interface { // bidirectional channel
	PipeWriterROnlyChan // aka "<-chan" - receive only
	PipeWriterSOnlyChan // aka "chan<-" - send only
}

type PipeWriterROnlyChan interface { // receive-only channel
	RequestPipeWriter() (dat *io.PipeWriter)        // the receive function - aka "some-new-PipeWriter-var := <-MyKind"
	TryPipeWriter() (dat *io.PipeWriter, open bool) // the multi-valued comma-ok receive function - aka "some-new-PipeWriter-var, ok := <-MyKind"
}

type PipeWriterSOnlyChan interface { // send-only channel
	ProvidePipeWriter(dat *io.PipeWriter) // the send function - aka "MyKind <- some PipeWriter"
}

type DChPipeWriter struct { // demand channel
	dat chan *io.PipeWriter
	req chan struct{}
}

func MakeDemandPipeWriterChan() *DChPipeWriter {
	d := new(DChPipeWriter)
	d.dat = make(chan *io.PipeWriter)
	d.req = make(chan struct{})
	return d
}

func MakeDemandPipeWriterBuff(cap int) *DChPipeWriter {
	d := new(DChPipeWriter)
	d.dat = make(chan *io.PipeWriter, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePipeWriter is the send function - aka "MyKind <- some PipeWriter"
func (c *DChPipeWriter) ProvidePipeWriter(dat *io.PipeWriter) {
	<-c.req
	c.dat <- dat
}

// RequestPipeWriter is the receive function - aka "some PipeWriter <- MyKind"
func (c *DChPipeWriter) RequestPipeWriter() (dat *io.PipeWriter) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPipeWriter is the comma-ok multi-valued form of RequestPipeWriter and
// reports whether a received value was sent before the PipeWriter channel was closed.
func (c *DChPipeWriter) TryPipeWriter() (dat *io.PipeWriter, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
