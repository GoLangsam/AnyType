// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFileSChan represents a
// bidirectional
// channel
type FsFileSChan interface {
	FsFileSROnlyChan // aka "<-chan" - receive only
	FsFileSSOnlyChan // aka "chan<-" - send only
}

// FsFileSROnlyChan represents a
// receive-only
// channel
type FsFileSROnlyChan interface {
	RequestFsFileS() (dat fs.FsFileS)        // the receive function - aka "MyFsFileS := <-MyFsFileSROnlyChan"
	TryFsFileS() (dat fs.FsFileS, open bool) // the multi-valued comma-ok receive function - aka "MyFsFileS, ok := <-MyFsFileSROnlyChan"
}

// FsFileSSOnlyChan represents a
// send-only
// channel
type FsFileSSOnlyChan interface {
	ProvideFsFileS(dat fs.FsFileS) // the send function - aka "MyKind <- some FsFileS"
}
