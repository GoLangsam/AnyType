// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsBoolean

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type BoolSChan interface { // bidirectional channel
	BoolSROnlyChan // aka "<-chan" - receive only
	BoolSSOnlyChan // aka "chan<-" - send only
}

type BoolSROnlyChan interface { // receive-only channel
	RequestBoolS() (dat []bool)        // the receive function - aka "some-new-BoolS-var := <-MyKind"
	TryBoolS() (dat []bool, open bool) // the multi-valued comma-ok receive function - aka "some-new-BoolS-var, ok := <-MyKind"
}

type BoolSSOnlyChan interface { // send-only channel
	ProvideBoolS(dat []bool) // the send function - aka "MyKind <- some BoolS"
}
