// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsDataChan represents a
// bidirectional
// channel
type FsDataChan interface {
	FsDataROnlyChan // aka "<-chan" - receive only
	FsDataSOnlyChan // aka "chan<-" - send only
}

// FsDataROnlyChan represents a
// receive-only
// channel
type FsDataROnlyChan interface {
	RequestFsData() (dat *fs.FsData)        // the receive function - aka "MyFsData := <-MyFsDataROnlyChan"
	TryFsData() (dat *fs.FsData, open bool) // the multi-valued comma-ok receive function - aka "MyFsData, ok := <-MyFsDataROnlyChan"
}

// FsDataSOnlyChan represents a
// send-only
// channel
type FsDataSOnlyChan interface {
	ProvideFsData(dat *fs.FsData) // the send function - aka "MyKind <- some FsData"
}
