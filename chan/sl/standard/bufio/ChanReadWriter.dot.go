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
