// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type LimitedReaderChan interface { // bidirectional channel
	LimitedReaderROnlyChan // aka "<-chan" - receive only
	LimitedReaderSOnlyChan // aka "chan<-" - send only
}

type LimitedReaderROnlyChan interface { // receive-only channel
	RequestLimitedReader() (dat *io.LimitedReader)        // the receive function - aka "some-new-LimitedReader-var := <-MyKind"
	TryLimitedReader() (dat *io.LimitedReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-LimitedReader-var, ok := <-MyKind"
}

type LimitedReaderSOnlyChan interface { // send-only channel
	ProvideLimitedReader(dat *io.LimitedReader) // the send function - aka "MyKind <- some LimitedReader"
}