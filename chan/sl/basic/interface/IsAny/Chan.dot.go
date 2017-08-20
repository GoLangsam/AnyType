// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsAny

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Chan represents a
// bidirectional
// channel
type Chan interface {
	ROnlyChan // aka "<-chan" - receive only
	SOnlyChan // aka "chan<-" - send only
}

// ROnlyChan represents a
// receive-only
// channel
type ROnlyChan interface {
	Request() (dat interface{})        // the receive function - aka "My := <-MyROnlyChan"
	Try() (dat interface{}, open bool) // the multi-valued comma-ok receive function - aka "My, ok := <-MyROnlyChan"
}

// SOnlyChan represents a
// send-only
// channel
type SOnlyChan interface {
	Provide(dat interface{}) // the send function - aka "MyKind <- some "
}
