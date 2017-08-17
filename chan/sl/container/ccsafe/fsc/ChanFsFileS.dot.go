// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsFileSChan interface { // bidirectional channel
	FsFileSROnlyChan // aka "<-chan" - receive only
	FsFileSSOnlyChan // aka "chan<-" - send only
}

type FsFileSROnlyChan interface { // receive-only channel
	RequestFsFileS() (dat fs.FsFileS)        // the receive function - aka "some-new-FsFileS-var := <-MyKind"
	TryFsFileS() (dat fs.FsFileS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFileS-var, ok := <-MyKind"
}

type FsFileSSOnlyChan interface { // send-only channel
	ProvideFsFileS(dat fs.FsFileS) // the send function - aka "MyKind <- some FsFileS"
}
