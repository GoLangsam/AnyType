// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type PipeReaderChan interface { // bidirectional channel
	PipeReaderROnlyChan // aka "<-chan" - receive only
	PipeReaderSOnlyChan // aka "chan<-" - send only
}

type PipeReaderROnlyChan interface { // receive-only channel
	RequestPipeReader() (dat *io.PipeReader)        // the receive function - aka "some-new-PipeReader-var := <-MyKind"
	TryPipeReader() (dat *io.PipeReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-PipeReader-var, ok := <-MyKind"
}

type PipeReaderSOnlyChan interface { // send-only channel
	ProvidePipeReader(dat *io.PipeReader) // the send function - aka "MyKind <- some PipeReader"
}
