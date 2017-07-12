// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int8Chan interface { // bidirectional channel
	Int8ROnlyChan // aka "<-chan" - receive only
	Int8SOnlyChan // aka "chan<-" - send only
}

type Int8ROnlyChan interface { // receive-only channel
	RequestInt8() (dat int8)        // the receive function - aka "some-new-Int8-var := <-MyKind"
	TryInt8() (dat int8, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int8-var, ok := <-MyKind"
}

type Int8SOnlyChan interface { // send-only channel
	ProvideInt8(dat int8) // the send function - aka "MyKind <- some Int8"
}
