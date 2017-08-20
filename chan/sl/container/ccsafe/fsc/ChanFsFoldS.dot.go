// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFoldSChan represents a
// bidirectional
// channel
type FsFoldSChan interface {
	FsFoldSROnlyChan // aka "<-chan" - receive only
	FsFoldSSOnlyChan // aka "chan<-" - send only
}

// FsFoldSROnlyChan represents a
// receive-only
// channel
type FsFoldSROnlyChan interface {
	RequestFsFoldS() (dat fs.FsFoldS)        // the receive function - aka "MyFsFoldS := <-MyFsFoldSROnlyChan"
	TryFsFoldS() (dat fs.FsFoldS, open bool) // the multi-valued comma-ok receive function - aka "MyFsFoldS, ok := <-MyFsFoldSROnlyChan"
}

// FsFoldSSOnlyChan represents a
// send-only
// channel
type FsFoldSSOnlyChan interface {
	ProvideFsFoldS(dat fs.FsFoldS) // the send function - aka "MyKind <- some FsFoldS"
}
