// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type SplitFuncChan interface { // bidirectional channel
	SplitFuncROnlyChan // aka "<-chan" - receive only
	SplitFuncSOnlyChan // aka "chan<-" - send only
}

type SplitFuncROnlyChan interface { // receive-only channel
	RequestSplitFunc() (dat bufio.SplitFunc)        // the receive function - aka "some-new-SplitFunc-var := <-MyKind"
	TrySplitFunc() (dat bufio.SplitFunc, open bool) // the multi-valued comma-ok receive function - aka "some-new-SplitFunc-var, ok := <-MyKind"
}

type SplitFuncSOnlyChan interface { // send-only channel
	ProvideSplitFunc(dat bufio.SplitFunc) // the send function - aka "MyKind <- some SplitFunc"
}
