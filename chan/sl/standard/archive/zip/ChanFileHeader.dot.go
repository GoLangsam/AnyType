// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

type FileHeaderChan interface { // bidirectional channel
	FileHeaderROnlyChan // aka "<-chan" - receive only
	FileHeaderSOnlyChan // aka "chan<-" - send only
}

type FileHeaderROnlyChan interface { // receive-only channel
	RequestFileHeader() (dat zip.FileHeader)        // the receive function - aka "some-new-FileHeader-var := <-MyKind"
	TryFileHeader() (dat zip.FileHeader, open bool) // the multi-valued comma-ok receive function - aka "some-new-FileHeader-var, ok := <-MyKind"
}

type FileHeaderSOnlyChan interface { // send-only channel
	ProvideFileHeader(dat zip.FileHeader) // the send function - aka "MyKind <- some FileHeader"
}
