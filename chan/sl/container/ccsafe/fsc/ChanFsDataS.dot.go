// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsDataSChan represents a
// bidirectional
// channel
type FsDataSChan interface {
	FsDataSROnlyChan // aka "<-chan" - receive only
	FsDataSSOnlyChan // aka "chan<-" - send only
}

// FsDataSROnlyChan represents a
// receive-only
// channel
type FsDataSROnlyChan interface {
	RequestFsDataS() (dat fs.FsDataS)        // the receive function - aka "MyFsDataS := <-MyFsDataSROnlyChan"
	TryFsDataS() (dat fs.FsDataS, open bool) // the multi-valued comma-ok receive function - aka "MyFsDataS, ok := <-MyFsDataSROnlyChan"
}

// FsDataSSOnlyChan represents a
// send-only
// channel
type FsDataSSOnlyChan interface {
	ProvideFsDataS(dat fs.FsDataS) // the send function - aka "MyKind <- some FsDataS"
}
