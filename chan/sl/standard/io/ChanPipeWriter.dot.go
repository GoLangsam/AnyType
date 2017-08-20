// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// PipeWriterChan represents a
// bidirectional
// channel
type PipeWriterChan interface {
	PipeWriterROnlyChan // aka "<-chan" - receive only
	PipeWriterSOnlyChan // aka "chan<-" - send only
}

// PipeWriterROnlyChan represents a
// receive-only
// channel
type PipeWriterROnlyChan interface {
	RequestPipeWriter() (dat *io.PipeWriter)        // the receive function - aka "MyPipeWriter := <-MyPipeWriterROnlyChan"
	TryPipeWriter() (dat *io.PipeWriter, open bool) // the multi-valued comma-ok receive function - aka "MyPipeWriter, ok := <-MyPipeWriterROnlyChan"
}

// PipeWriterSOnlyChan represents a
// send-only
// channel
type PipeWriterSOnlyChan interface {
	ProvidePipeWriter(dat *io.PipeWriter) // the send function - aka "MyKind <- some PipeWriter"
}
