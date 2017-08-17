// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

type SignalChan interface { // bidirectional channel
	SignalROnlyChan // aka "<-chan" - receive only
	SignalSOnlyChan // aka "chan<-" - send only
}

type SignalROnlyChan interface { // receive-only channel
	RequestSignal() (dat os.Signal)        // the receive function - aka "some-new-Signal-var := <-MyKind"
	TrySignal() (dat os.Signal, open bool) // the multi-valued comma-ok receive function - aka "some-new-Signal-var, ok := <-MyKind"
}

type SignalSOnlyChan interface { // send-only channel
	ProvideSignal(dat os.Signal) // the send function - aka "MyKind <- some Signal"
}
