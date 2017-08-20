// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// FileChan represents a
// bidirectional
// channel
type FileChan interface {
	FileROnlyChan // aka "<-chan" - receive only
	FileSOnlyChan // aka "chan<-" - send only
}

// FileROnlyChan represents a
// receive-only
// channel
type FileROnlyChan interface {
	RequestFile() (dat *os.File)        // the receive function - aka "MyFile := <-MyFileROnlyChan"
	TryFile() (dat *os.File, open bool) // the multi-valued comma-ok receive function - aka "MyFile, ok := <-MyFileROnlyChan"
}

// FileSOnlyChan represents a
// send-only
// channel
type FileSOnlyChan interface {
	ProvideFile(dat *os.File) // the send function - aka "MyKind <- some File"
}
