// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// ErrorSChan represents a
// bidirectional
// channel
type ErrorSChan interface {
	ErrorSROnlyChan // aka "<-chan" - receive only
	ErrorSSOnlyChan // aka "chan<-" - send only
}

// ErrorSROnlyChan represents a
// receive-only
// channel
type ErrorSROnlyChan interface {
	RequestErrorS() (dat []error)        // the receive function - aka "MyErrorS := <-MyErrorSROnlyChan"
	TryErrorS() (dat []error, open bool) // the multi-valued comma-ok receive function - aka "MyErrorS, ok := <-MyErrorSROnlyChan"
}

// ErrorSSOnlyChan represents a
// send-only
// channel
type ErrorSSOnlyChan interface {
	ProvideErrorS(dat []error) // the send function - aka "MyKind <- some ErrorS"
}
