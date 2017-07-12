// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsDataSChan interface { // bidirectional channel
	FsDataSROnlyChan // aka "<-chan" - receive only
	FsDataSSOnlyChan // aka "chan<-" - send only
}

type FsDataSROnlyChan interface { // receive-only channel
	RequestFsDataS() (dat fs.FsDataS)        // the receive function - aka "some-new-FsDataS-var := <-MyKind"
	TryFsDataS() (dat fs.FsDataS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsDataS-var, ok := <-MyKind"
}

type FsDataSSOnlyChan interface { // send-only channel
	ProvideFsDataS(dat fs.FsDataS) // the send function - aka "MyKind <- some FsDataS"
}