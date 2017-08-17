// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

type FsFoldSChan interface { // bidirectional channel
	FsFoldSROnlyChan // aka "<-chan" - receive only
	FsFoldSSOnlyChan // aka "chan<-" - send only
}

type FsFoldSROnlyChan interface { // receive-only channel
	RequestFsFoldS() (dat fs.FsFoldS)        // the receive function - aka "some-new-FsFoldS-var := <-MyKind"
	TryFsFoldS() (dat fs.FsFoldS, open bool) // the multi-valued comma-ok receive function - aka "some-new-FsFoldS-var, ok := <-MyKind"
}

type FsFoldSSOnlyChan interface { // send-only channel
	ProvideFsFoldS(dat fs.FsFoldS) // the send function - aka "MyKind <- some FsFoldS"
}
