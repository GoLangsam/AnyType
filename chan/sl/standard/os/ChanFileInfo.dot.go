// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// FileInfoChan represents a
// bidirectional
// channel
type FileInfoChan interface {
	FileInfoROnlyChan // aka "<-chan" - receive only
	FileInfoSOnlyChan // aka "chan<-" - send only
}

// FileInfoROnlyChan represents a
// receive-only
// channel
type FileInfoROnlyChan interface {
	RequestFileInfo() (dat os.FileInfo)        // the receive function - aka "MyFileInfo := <-MyFileInfoROnlyChan"
	TryFileInfo() (dat os.FileInfo, open bool) // the multi-valued comma-ok receive function - aka "MyFileInfo, ok := <-MyFileInfoROnlyChan"
}

// FileInfoSOnlyChan represents a
// send-only
// channel
type FileInfoSOnlyChan interface {
	ProvideFileInfo(dat os.FileInfo) // the send function - aka "MyKind <- some FileInfo"
}
