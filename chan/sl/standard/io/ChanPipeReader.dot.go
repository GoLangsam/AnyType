// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// PipeReaderChan represents a
// bidirectional
// channel
type PipeReaderChan interface {
	PipeReaderROnlyChan // aka "<-chan" - receive only
	PipeReaderSOnlyChan // aka "chan<-" - send only
}

// PipeReaderROnlyChan represents a
// receive-only
// channel
type PipeReaderROnlyChan interface {
	RequestPipeReader() (dat *io.PipeReader)        // the receive function - aka "MyPipeReader := <-MyPipeReaderROnlyChan"
	TryPipeReader() (dat *io.PipeReader, open bool) // the multi-valued comma-ok receive function - aka "MyPipeReader, ok := <-MyPipeReaderROnlyChan"
}

// PipeReaderSOnlyChan represents a
// send-only
// channel
type PipeReaderSOnlyChan interface {
	ProvidePipeReader(dat *io.PipeReader) // the send function - aka "MyKind <- some PipeReader"
}
