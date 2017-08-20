// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// ByteScannerChan represents a
// bidirectional
// channel
type ByteScannerChan interface {
	ByteScannerROnlyChan // aka "<-chan" - receive only
	ByteScannerSOnlyChan // aka "chan<-" - send only
}

// ByteScannerROnlyChan represents a
// receive-only
// channel
type ByteScannerROnlyChan interface {
	RequestByteScanner() (dat io.ByteScanner)        // the receive function - aka "MyByteScanner := <-MyByteScannerROnlyChan"
	TryByteScanner() (dat io.ByteScanner, open bool) // the multi-valued comma-ok receive function - aka "MyByteScanner, ok := <-MyByteScannerROnlyChan"
}

// ByteScannerSOnlyChan represents a
// send-only
// channel
type ByteScannerSOnlyChan interface {
	ProvideByteScanner(dat io.ByteScanner) // the send function - aka "MyKind <- some ByteScanner"
}
