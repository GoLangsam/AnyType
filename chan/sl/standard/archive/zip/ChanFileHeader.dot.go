// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

// FileHeaderChan represents a
// bidirectional
// channel
type FileHeaderChan interface {
	FileHeaderROnlyChan // aka "<-chan" - receive only
	FileHeaderSOnlyChan // aka "chan<-" - send only
}

// FileHeaderROnlyChan represents a
// receive-only
// channel
type FileHeaderROnlyChan interface {
	RequestFileHeader() (dat zip.FileHeader)        // the receive function - aka "MyFileHeader := <-MyFileHeaderROnlyChan"
	TryFileHeader() (dat zip.FileHeader, open bool) // the multi-valued comma-ok receive function - aka "MyFileHeader, ok := <-MyFileHeaderROnlyChan"
}

// FileHeaderSOnlyChan represents a
// send-only
// channel
type FileHeaderSOnlyChan interface {
	ProvideFileHeader(dat zip.FileHeader) // the send function - aka "MyKind <- some FileHeader"
}
