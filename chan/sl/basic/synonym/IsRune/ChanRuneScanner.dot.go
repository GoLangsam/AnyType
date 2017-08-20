// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// RuneScannerChan represents a
// bidirectional
// channel
type RuneScannerChan interface {
	RuneScannerROnlyChan // aka "<-chan" - receive only
	RuneScannerSOnlyChan // aka "chan<-" - send only
}

// RuneScannerROnlyChan represents a
// receive-only
// channel
type RuneScannerROnlyChan interface {
	RequestRuneScanner() (dat io.RuneScanner)        // the receive function - aka "MyRuneScanner := <-MyRuneScannerROnlyChan"
	TryRuneScanner() (dat io.RuneScanner, open bool) // the multi-valued comma-ok receive function - aka "MyRuneScanner, ok := <-MyRuneScannerROnlyChan"
}

// RuneScannerSOnlyChan represents a
// send-only
// channel
type RuneScannerSOnlyChan interface {
	ProvideRuneScanner(dat io.RuneScanner) // the send function - aka "MyKind <- some RuneScanner"
}
