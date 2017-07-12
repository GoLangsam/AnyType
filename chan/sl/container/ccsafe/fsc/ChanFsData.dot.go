// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsDataChan interface { // bidirectional channel
	FsDataROnlyChan // aka "<-chan" - receive only
	FsDataSOnlyChan // aka "chan<-" - send only
}

type FsDataROnlyChan interface { // receive-only channel
	RequestFsData() (dat *fs.FsData)        // the receive function - aka "some-new-FsData-var := <-MyKind"
	TryFsData() (dat *fs.FsData, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsData-var, ok := <-MyKind"
}

type FsDataSOnlyChan interface { // send-only channel
	ProvideFsData(dat *fs.FsData) // the send function - aka "MyKind <- some FsData"
}