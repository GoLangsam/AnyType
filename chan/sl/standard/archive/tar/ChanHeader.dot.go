// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

type HeaderChan interface { // bidirectional channel
	HeaderROnlyChan // aka "<-chan" - receive only
	HeaderSOnlyChan // aka "chan<-" - send only
}

type HeaderROnlyChan interface { // receive-only channel
	RequestHeader() (dat *tar.Header)        // the receive function - aka "some-new-Header-var := <-MyKind"
	TryHeader() (dat *tar.Header, open bool) // the multi-valued comma-ok receive function - aka "some-new-Header-var, ok := <-MyKind"
}

type HeaderSOnlyChan interface { // send-only channel
	ProvideHeader(dat *tar.Header) // the send function - aka "MyKind <- some Header"
}
