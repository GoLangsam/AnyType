// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

type FileChan interface { // bidirectional channel
	FileROnlyChan // aka "<-chan" - receive only
	FileSOnlyChan // aka "chan<-" - send only
}

type FileROnlyChan interface { // receive-only channel
	RequestFile() (dat zip.File)        // the receive function - aka "some-new-File-var := <-MyKind"
	TryFile() (dat zip.File, open bool) // the multi-valued comma-ok receive function - aka "some-new-File-var, ok := <-MyKind"
}

type FileSOnlyChan interface { // send-only channel
	ProvideFile(dat zip.File) // the send function - aka "MyKind <- some File"
}
