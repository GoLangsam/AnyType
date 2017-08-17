// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsBaseSChan interface { // bidirectional channel
	FsBaseSROnlyChan // aka "<-chan" - receive only
	FsBaseSSOnlyChan // aka "chan<-" - send only
}

type FsBaseSROnlyChan interface { // receive-only channel
	RequestFsBaseS() (dat fs.FsBaseS)        // the receive function - aka "some-new-FsBaseS-var := <-MyKind"
	TryFsBaseS() (dat fs.FsBaseS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsBaseS-var, ok := <-MyKind"
}

type FsBaseSSOnlyChan interface { // send-only channel
	ProvideFsBaseS(dat fs.FsBaseS) // the send function - aka "MyKind <- some FsBaseS"
}
