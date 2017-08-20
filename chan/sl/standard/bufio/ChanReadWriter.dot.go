// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// ReadWriterChan represents a
// bidirectional
// channel
type ReadWriterChan interface {
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

// ReadWriterROnlyChan represents a
// receive-only
// channel
type ReadWriterROnlyChan interface {
	RequestReadWriter() (dat *bufio.ReadWriter)        // the receive function - aka "MyReadWriter := <-MyReadWriterROnlyChan"
	TryReadWriter() (dat *bufio.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "MyReadWriter, ok := <-MyReadWriterROnlyChan"
}

// ReadWriterSOnlyChan represents a
// send-only
// channel
type ReadWriterSOnlyChan interface {
	ProvideReadWriter(dat *bufio.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}
