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
