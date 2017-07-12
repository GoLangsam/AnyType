// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

type FileInfoChan interface { // bidirectional channel
	FileInfoROnlyChan // aka "<-chan" - receive only
	FileInfoSOnlyChan // aka "chan<-" - send only
}

type FileInfoROnlyChan interface { // receive-only channel
	RequestFileInfo() (dat os.FileInfo)        // the receive function - aka "some-new-FileInfo-var := <-MyKind"
	TryFileInfo() (dat os.FileInfo, open bool) // the multi-valued comma-ok receive function - aka "some-new-FileInfo-var, ok := <-MyKind"
}

type FileInfoSOnlyChan interface { // send-only channel
	ProvideFileInfo(dat os.FileInfo) // the send function - aka "MyKind <- some FileInfo"
}
