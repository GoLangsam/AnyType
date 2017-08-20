// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/dotpath"
)

// DotPathChan represents a
// bidirectional
// channel
type DotPathChan interface {
	DotPathROnlyChan // aka "<-chan" - receive only
	DotPathSOnlyChan // aka "chan<-" - send only
}

// DotPathROnlyChan represents a
// receive-only
// channel
type DotPathROnlyChan interface {
	RequestDotPath() (dat dotpath.DotPath)        // the receive function - aka "MyDotPath := <-MyDotPathROnlyChan"
	TryDotPath() (dat dotpath.DotPath, open bool) // the multi-valued comma-ok receive function - aka "MyDotPath, ok := <-MyDotPathROnlyChan"
}

// DotPathSOnlyChan represents a
// send-only
// channel
type DotPathSOnlyChan interface {
	ProvideDotPath(dat dotpath.DotPath) // the send function - aka "MyKind <- some DotPath"
}
