// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsBaseChan interface { // bidirectional channel
	FsBaseROnlyChan // aka "<-chan" - receive only
	FsBaseSOnlyChan // aka "chan<-" - send only
}

type FsBaseROnlyChan interface { // receive-only channel
	RequestFsBase() (dat *fs.FsBase)        // the receive function - aka "some-new-FsBase-var := <-MyKind"
	TryFsBase() (dat *fs.FsBase, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsBase-var, ok := <-MyKind"
}

type FsBaseSOnlyChan interface { // send-only channel
	ProvideFsBase(dat *fs.FsBase) // the send function - aka "MyKind <- some FsBase"
}
