// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseSChan represents a
// bidirectional
// channel
type FsBaseSChan interface {
	FsBaseSROnlyChan // aka "<-chan" - receive only
	FsBaseSSOnlyChan // aka "chan<-" - send only
}

// FsBaseSROnlyChan represents a
// receive-only
// channel
type FsBaseSROnlyChan interface {
	RequestFsBaseS() (dat fs.FsBaseS)        // the receive function - aka "MyFsBaseS := <-MyFsBaseSROnlyChan"
	TryFsBaseS() (dat fs.FsBaseS, open bool) // the multi-valued comma-ok receive function - aka "MyFsBaseS, ok := <-MyFsBaseSROnlyChan"
}

// FsBaseSSOnlyChan represents a
// send-only
// channel
type FsBaseSSOnlyChan interface {
	ProvideFsBaseS(dat fs.FsBaseS) // the send function - aka "MyKind <- some FsBaseS"
}
