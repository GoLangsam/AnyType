// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type RuneReaderChan interface { // bidirectional channel
	RuneReaderROnlyChan // aka "<-chan" - receive only
	RuneReaderSOnlyChan // aka "chan<-" - send only
}

type RuneReaderROnlyChan interface { // receive-only channel
	RequestRuneReader() (dat io.RuneReader)        // the receive function - aka "some-new-RuneReader-var := <-MyKind"
	TryRuneReader() (dat io.RuneReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-RuneReader-var, ok := <-MyKind"
}

type RuneReaderSOnlyChan interface { // send-only channel
	ProvideRuneReader(dat io.RuneReader) // the send function - aka "MyKind <- some RuneReader"
}
