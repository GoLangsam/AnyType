// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsBoolean

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// BoolSChan represents a
// bidirectional
// channel
type BoolSChan interface {
	BoolSROnlyChan // aka "<-chan" - receive only
	BoolSSOnlyChan // aka "chan<-" - send only
}

// BoolSROnlyChan represents a
// receive-only
// channel
type BoolSROnlyChan interface {
	RequestBoolS() (dat []bool)        // the receive function - aka "MyBoolS := <-MyBoolSROnlyChan"
	TryBoolS() (dat []bool, open bool) // the multi-valued comma-ok receive function - aka "MyBoolS, ok := <-MyBoolSROnlyChan"
}

// BoolSSOnlyChan represents a
// send-only
// channel
type BoolSSOnlyChan interface {
	ProvideBoolS(dat []bool) // the send function - aka "MyKind <- some BoolS"
}
