// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsInfoSChan interface { // bidirectional channel
	FsInfoSROnlyChan // aka "<-chan" - receive only
	FsInfoSSOnlyChan // aka "chan<-" - send only
}

type FsInfoSROnlyChan interface { // receive-only channel
	RequestFsInfoS() (dat fs.FsInfoS)        // the receive function - aka "some-new-FsInfoS-var := <-MyKind"
	TryFsInfoS() (dat fs.FsInfoS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsInfoS-var, ok := <-MyKind"
}

type FsInfoSSOnlyChan interface { // send-only channel
	ProvideFsInfoS(dat fs.FsInfoS) // the send function - aka "MyKind <- some FsInfoS"
}
