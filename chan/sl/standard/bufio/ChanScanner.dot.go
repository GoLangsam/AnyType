// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type ScannerChan interface { // bidirectional channel
	ScannerROnlyChan // aka "<-chan" - receive only
	ScannerSOnlyChan // aka "chan<-" - send only
}

type ScannerROnlyChan interface { // receive-only channel
	RequestScanner() (dat *bufio.Scanner)        // the receive function - aka "some-new-Scanner-var := <-MyKind"
	TryScanner() (dat *bufio.Scanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-Scanner-var, ok := <-MyKind"
}

type ScannerSOnlyChan interface { // send-only channel
	ProvideScanner(dat *bufio.Scanner) // the send function - aka "MyKind <- some Scanner"
}
