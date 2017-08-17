// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/lsm"
)

type LSMChan interface { // bidirectional channel
	LSMROnlyChan // aka "<-chan" - receive only
	LSMSOnlyChan // aka "chan<-" - send only
}

type LSMROnlyChan interface { // receive-only channel
	RequestLSM() (dat lsm.LazyStringerMap)        // the receive function - aka "some-new-LSM-var := <-MyKind"
	TryLSM() (dat lsm.LazyStringerMap, open bool) // the multi-valued comma-ok receive function - aka "some-new-LSM-var, ok := <-MyKind"
}

type LSMSOnlyChan interface { // send-only channel
	ProvideLSM(dat lsm.LazyStringerMap) // the send function - aka "MyKind <- some LSM"
}
