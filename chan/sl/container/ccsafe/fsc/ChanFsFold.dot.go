// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

type FsFoldChan interface { // bidirectional channel
	FsFoldROnlyChan // aka "<-chan" - receive only
	FsFoldSOnlyChan // aka "chan<-" - send only
}

type FsFoldROnlyChan interface { // receive-only channel
	RequestFsFold() (dat *fs.FsFold)        // the receive function - aka "some-new-FsFold-var := <-MyKind"
	TryFsFold() (dat *fs.FsFold, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFold-var, ok := <-MyKind"
}

type FsFoldSOnlyChan interface { // send-only channel
	ProvideFsFold(dat *fs.FsFold) // the send function - aka "MyKind <- some FsFold"
}
