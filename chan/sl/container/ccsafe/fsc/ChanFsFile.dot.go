// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFileChan represents a
// bidirectional
// channel
type FsFileChan interface {
	FsFileROnlyChan // aka "<-chan" - receive only
	FsFileSOnlyChan // aka "chan<-" - send only
}

// FsFileROnlyChan represents a
// receive-only
// channel
type FsFileROnlyChan interface {
	RequestFsFile() (dat *fs.FsFile)        // the receive function - aka "MyFsFile := <-MyFsFileROnlyChan"
	TryFsFile() (dat *fs.FsFile, open bool) // the multi-valued comma-ok receive function - aka "MyFsFile, ok := <-MyFsFileROnlyChan"
}

// FsFileSOnlyChan represents a
// send-only
// channel
type FsFileSOnlyChan interface {
	ProvideFsFile(dat *fs.FsFile) // the send function - aka "MyKind <- some FsFile"
}
