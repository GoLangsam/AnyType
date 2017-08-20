// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// SignalChan represents a
// bidirectional
// channel
type SignalChan interface {
	SignalROnlyChan // aka "<-chan" - receive only
	SignalSOnlyChan // aka "chan<-" - send only
}

// SignalROnlyChan represents a
// receive-only
// channel
type SignalROnlyChan interface {
	RequestSignal() (dat os.Signal)        // the receive function - aka "MySignal := <-MySignalROnlyChan"
	TrySignal() (dat os.Signal, open bool) // the multi-valued comma-ok receive function - aka "MySignal, ok := <-MySignalROnlyChan"
}

// SignalSOnlyChan represents a
// send-only
// channel
type SignalSOnlyChan interface {
	ProvideSignal(dat os.Signal) // the send function - aka "MyKind <- some Signal"
}
