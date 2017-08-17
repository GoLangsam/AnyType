// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsPathSChan interface { // bidirectional channel
	FsPathSROnlyChan // aka "<-chan" - receive only
	FsPathSSOnlyChan // aka "chan<-" - send only
}

type FsPathSROnlyChan interface { // receive-only channel
	RequestFsPathS() (dat fs.FsPathS)        // the receive function - aka "some-new-FsPathS-var := <-MyKind"
	TryFsPathS() (dat fs.FsPathS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsPathS-var, ok := <-MyKind"
}

type FsPathSSOnlyChan interface { // send-only channel
	ProvideFsPathS(dat fs.FsPathS) // the send function - aka "MyKind <- some FsPathS"
}
