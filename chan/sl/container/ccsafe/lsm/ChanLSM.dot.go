// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/lsm"
)

// LSMChan represents a
// bidirectional
// channel
type LSMChan interface {
	LSMROnlyChan // aka "<-chan" - receive only
	LSMSOnlyChan // aka "chan<-" - send only
}

// LSMROnlyChan represents a
// receive-only
// channel
type LSMROnlyChan interface {
	RequestLSM() (dat lsm.LazyStringerMap)        // the receive function - aka "MyLSM := <-MyLSMROnlyChan"
	TryLSM() (dat lsm.LazyStringerMap, open bool) // the multi-valued comma-ok receive function - aka "MyLSM, ok := <-MyLSMROnlyChan"
}

// LSMSOnlyChan represents a
// send-only
// channel
type LSMSOnlyChan interface {
	ProvideLSM(dat lsm.LazyStringerMap) // the send function - aka "MyKind <- some LSM"
}
