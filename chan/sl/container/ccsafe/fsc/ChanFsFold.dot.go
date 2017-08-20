// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// FsFoldChan represents a
// bidirectional
// channel
type FsFoldChan interface {
	FsFoldROnlyChan // aka "<-chan" - receive only
	FsFoldSOnlyChan // aka "chan<-" - send only
}

// FsFoldROnlyChan represents a
// receive-only
// channel
type FsFoldROnlyChan interface {
	RequestFsFold() (dat *fs.FsFold)        // the receive function - aka "MyFsFold := <-MyFsFoldROnlyChan"
	TryFsFold() (dat *fs.FsFold, open bool) // the multi-valued comma-ok receive function - aka "MyFsFold, ok := <-MyFsFoldROnlyChan"
}

// FsFoldSOnlyChan represents a
// send-only
// channel
type FsFoldSOnlyChan interface {
	ProvideFsFold(dat *fs.FsFold) // the send function - aka "MyKind <- some FsFold"
}
