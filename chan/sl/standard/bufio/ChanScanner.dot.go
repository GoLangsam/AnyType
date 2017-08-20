// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// ScannerChan represents a
// bidirectional
// channel
type ScannerChan interface {
	ScannerROnlyChan // aka "<-chan" - receive only
	ScannerSOnlyChan // aka "chan<-" - send only
}

// ScannerROnlyChan represents a
// receive-only
// channel
type ScannerROnlyChan interface {
	RequestScanner() (dat *bufio.Scanner)        // the receive function - aka "MyScanner := <-MyScannerROnlyChan"
	TryScanner() (dat *bufio.Scanner, open bool) // the multi-valued comma-ok receive function - aka "MyScanner, ok := <-MyScannerROnlyChan"
}

// ScannerSOnlyChan represents a
// send-only
// channel
type ScannerSOnlyChan interface {
	ProvideScanner(dat *bufio.Scanner) // the send function - aka "MyKind <- some Scanner"
}
