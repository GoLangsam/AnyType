// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsInfoSChan represents a
// bidirectional
// channel
type FsInfoSChan interface {
	FsInfoSROnlyChan // aka "<-chan" - receive only
	FsInfoSSOnlyChan // aka "chan<-" - send only
}

// FsInfoSROnlyChan represents a
// receive-only
// channel
type FsInfoSROnlyChan interface {
	RequestFsInfoS() (dat fs.FsInfoS)        // the receive function - aka "MyFsInfoS := <-MyFsInfoSROnlyChan"
	TryFsInfoS() (dat fs.FsInfoS, open bool) // the multi-valued comma-ok receive function - aka "MyFsInfoS, ok := <-MyFsInfoSROnlyChan"
}

// FsInfoSSOnlyChan represents a
// send-only
// channel
type FsInfoSSOnlyChan interface {
	ProvideFsInfoS(dat fs.FsInfoS) // the send function - aka "MyKind <- some FsInfoS"
}
