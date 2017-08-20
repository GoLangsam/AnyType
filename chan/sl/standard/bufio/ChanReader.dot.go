// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// ReaderChan represents a
// bidirectional
// channel
type ReaderChan interface {
	ReaderROnlyChan // aka "<-chan" - receive only
	ReaderSOnlyChan // aka "chan<-" - send only
}

// ReaderROnlyChan represents a
// receive-only
// channel
type ReaderROnlyChan interface {
	RequestReader() (dat *bufio.Reader)        // the receive function - aka "MyReader := <-MyReaderROnlyChan"
	TryReader() (dat *bufio.Reader, open bool) // the multi-valued comma-ok receive function - aka "MyReader, ok := <-MyReaderROnlyChan"
}

// ReaderSOnlyChan represents a
// send-only
// channel
type ReaderSOnlyChan interface {
	ProvideReader(dat *bufio.Reader) // the send function - aka "MyKind <- some Reader"
}
