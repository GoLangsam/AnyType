// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int16Chan interface { // bidirectional channel
	Int16ROnlyChan // aka "<-chan" - receive only
	Int16SOnlyChan // aka "chan<-" - send only
}

type Int16ROnlyChan interface { // receive-only channel
	RequestInt16() (dat int16)        // the receive function - aka "some-new-Int16-var := <-MyKind"
	TryInt16() (dat int16, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int16-var, ok := <-MyKind"
}

type Int16SOnlyChan interface { // send-only channel
	ProvideInt16(dat int16) // the send function - aka "MyKind <- some Int16"
}
