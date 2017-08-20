// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseChan represents a
// bidirectional
// channel
type FsBaseChan interface {
	FsBaseROnlyChan // aka "<-chan" - receive only
	FsBaseSOnlyChan // aka "chan<-" - send only
}

// FsBaseROnlyChan represents a
// receive-only
// channel
type FsBaseROnlyChan interface {
	RequestFsBase() (dat *fs.FsBase)        // the receive function - aka "MyFsBase := <-MyFsBaseROnlyChan"
	TryFsBase() (dat *fs.FsBase, open bool) // the multi-valued comma-ok receive function - aka "MyFsBase, ok := <-MyFsBaseROnlyChan"
}

// FsBaseSOnlyChan represents a
// send-only
// channel
type FsBaseSOnlyChan interface {
	ProvideFsBase(dat *fs.FsBase) // the send function - aka "MyKind <- some FsBase"
}
