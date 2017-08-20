// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsPathSChan represents a
// bidirectional
// channel
type FsPathSChan interface {
	FsPathSROnlyChan // aka "<-chan" - receive only
	FsPathSSOnlyChan // aka "chan<-" - send only
}

// FsPathSROnlyChan represents a
// receive-only
// channel
type FsPathSROnlyChan interface {
	RequestFsPathS() (dat fs.FsPathS)        // the receive function - aka "MyFsPathS := <-MyFsPathSROnlyChan"
	TryFsPathS() (dat fs.FsPathS, open bool) // the multi-valued comma-ok receive function - aka "MyFsPathS, ok := <-MyFsPathSROnlyChan"
}

// FsPathSSOnlyChan represents a
// send-only
// channel
type FsPathSSOnlyChan interface {
	ProvideFsPathS(dat fs.FsPathS) // the send function - aka "MyKind <- some FsPathS"
}
