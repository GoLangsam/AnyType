// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

type FileChan interface { // bidirectional channel
	FileROnlyChan // aka "<-chan" - receive only
	FileSOnlyChan // aka "chan<-" - send only
}

type FileROnlyChan interface { // receive-only channel
	RequestFile() (dat *os.File)        // the receive function - aka "some-new-File-var := <-MyKind"
	TryFile() (dat *os.File, open bool) // the multi-valued comma-ok receive function - aka "some-new-File-var, ok := <-MyKind"
}

type FileSOnlyChan interface { // send-only channel
	ProvideFile(dat *os.File) // the send function - aka "MyKind <- some File"
}
