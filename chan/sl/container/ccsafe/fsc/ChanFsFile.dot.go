// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsFileChan interface { // bidirectional channel
	FsFileROnlyChan // aka "<-chan" - receive only
	FsFileSOnlyChan // aka "chan<-" - send only
}

type FsFileROnlyChan interface { // receive-only channel
	RequestFsFile() (dat *fs.FsFile)        // the receive function - aka "some-new-FsFile-var := <-MyKind"
	TryFsFile() (dat *fs.FsFile, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFile-var, ok := <-MyKind"
}

type FsFileSOnlyChan interface { // send-only channel
	ProvideFsFile(dat *fs.FsFile) // the send function - aka "MyKind <- some FsFile"
}
