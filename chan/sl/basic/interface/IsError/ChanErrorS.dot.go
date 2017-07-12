// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type ErrorSChan interface { // bidirectional channel
	ErrorSROnlyChan // aka "<-chan" - receive only
	ErrorSSOnlyChan // aka "chan<-" - send only
}

type ErrorSROnlyChan interface { // receive-only channel
	RequestErrorS() (dat []error)        // the receive function - aka "some-new-ErrorS-var := <-MyKind"
	TryErrorS() (dat []error, open bool) // the multi-valued comma-ok receive function - aka "some-new-ErrorS-var, ok := <-MyKind"
}

type ErrorSSOnlyChan interface { // send-only channel
	ProvideErrorS(dat []error) // the send function - aka "MyKind <- some ErrorS"
}
