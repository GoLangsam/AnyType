// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/dotpath"
)

type DotPathChan interface { // bidirectional channel
	DotPathROnlyChan // aka "<-chan" - receive only
	DotPathSOnlyChan // aka "chan<-" - send only
}

type DotPathROnlyChan interface { // receive-only channel
	RequestDotPath() (dat dotpath.DotPath)        // the receive function - aka "some-new-DotPath-var := <-MyKind"
	TryDotPath() (dat dotpath.DotPath, open bool) // the multi-valued comma-ok receive function - aka "some-new-DotPath-var, ok := <-MyKind"
}

type DotPathSOnlyChan interface { // send-only channel
	ProvideDotPath(dat dotpath.DotPath) // the send function - aka "MyKind <- some DotPath"
}
