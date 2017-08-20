// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// LimitedReaderChan represents a
// bidirectional
// channel
type LimitedReaderChan interface {
	LimitedReaderROnlyChan // aka "<-chan" - receive only
	LimitedReaderSOnlyChan // aka "chan<-" - send only
}

// LimitedReaderROnlyChan represents a
// receive-only
// channel
type LimitedReaderROnlyChan interface {
	RequestLimitedReader() (dat *io.LimitedReader)        // the receive function - aka "MyLimitedReader := <-MyLimitedReaderROnlyChan"
	TryLimitedReader() (dat *io.LimitedReader, open bool) // the multi-valued comma-ok receive function - aka "MyLimitedReader, ok := <-MyLimitedReaderROnlyChan"
}

// LimitedReaderSOnlyChan represents a
// send-only
// channel
type LimitedReaderSOnlyChan interface {
	ProvideLimitedReader(dat *io.LimitedReader) // the send function - aka "MyKind <- some LimitedReader"
}
