// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Float32Chan represents a
// bidirectional
// channel
type Float32Chan interface {
	Float32ROnlyChan // aka "<-chan" - receive only
	Float32SOnlyChan // aka "chan<-" - send only
}

// Float32ROnlyChan represents a
// receive-only
// channel
type Float32ROnlyChan interface {
	RequestFloat32() (dat float32)        // the receive function - aka "MyFloat32 := <-MyFloat32ROnlyChan"
	TryFloat32() (dat float32, open bool) // the multi-valued comma-ok receive function - aka "MyFloat32, ok := <-MyFloat32ROnlyChan"
}

// Float32SOnlyChan represents a
// send-only
// channel
type Float32SOnlyChan interface {
	ProvideFloat32(dat float32) // the send function - aka "MyKind <- some Float32"
}
