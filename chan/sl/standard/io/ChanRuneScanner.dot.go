// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type RuneScannerChan interface { // bidirectional channel
	RuneScannerROnlyChan // aka "<-chan" - receive only
	RuneScannerSOnlyChan // aka "chan<-" - send only
}

type RuneScannerROnlyChan interface { // receive-only channel
	RequestRuneScanner() (dat io.RuneScanner)        // the receive function - aka "some-new-RuneScanner-var := <-MyKind"
	TryRuneScanner() (dat io.RuneScanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-RuneScanner-var, ok := <-MyKind"
}

type RuneScannerSOnlyChan interface { // send-only channel
	ProvideRuneScanner(dat io.RuneScanner) // the send function - aka "MyKind <- some RuneScanner"
}
