// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Float64Chan represents a
// bidirectional
// channel
type Float64Chan interface {
	Float64ROnlyChan // aka "<-chan" - receive only
	Float64SOnlyChan // aka "chan<-" - send only
}

// Float64ROnlyChan represents a
// receive-only
// channel
type Float64ROnlyChan interface {
	RequestFloat64() (dat float64)        // the receive function - aka "MyFloat64 := <-MyFloat64ROnlyChan"
	TryFloat64() (dat float64, open bool) // the multi-valued comma-ok receive function - aka "MyFloat64, ok := <-MyFloat64ROnlyChan"
}

// Float64SOnlyChan represents a
// send-only
// channel
type Float64SOnlyChan interface {
	ProvideFloat64(dat float64) // the send function - aka "MyKind <- some Float64"
}
