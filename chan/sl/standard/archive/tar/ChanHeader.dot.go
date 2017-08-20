// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

// HeaderChan represents a
// bidirectional
// channel
type HeaderChan interface {
	HeaderROnlyChan // aka "<-chan" - receive only
	HeaderSOnlyChan // aka "chan<-" - send only
}

// HeaderROnlyChan represents a
// receive-only
// channel
type HeaderROnlyChan interface {
	RequestHeader() (dat *tar.Header)        // the receive function - aka "MyHeader := <-MyHeaderROnlyChan"
	TryHeader() (dat *tar.Header, open bool) // the multi-valued comma-ok receive function - aka "MyHeader, ok := <-MyHeaderROnlyChan"
}

// HeaderSOnlyChan represents a
// send-only
// channel
type HeaderSOnlyChan interface {
	ProvideHeader(dat *tar.Header) // the send function - aka "MyKind <- some Header"
}
