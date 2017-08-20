// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// SplitFuncChan represents a
// bidirectional
// channel
type SplitFuncChan interface {
	SplitFuncROnlyChan // aka "<-chan" - receive only
	SplitFuncSOnlyChan // aka "chan<-" - send only
}

// SplitFuncROnlyChan represents a
// receive-only
// channel
type SplitFuncROnlyChan interface {
	RequestSplitFunc() (dat bufio.SplitFunc)        // the receive function - aka "MySplitFunc := <-MySplitFuncROnlyChan"
	TrySplitFunc() (dat bufio.SplitFunc, open bool) // the multi-valued comma-ok receive function - aka "MySplitFunc, ok := <-MySplitFuncROnlyChan"
}

// SplitFuncSOnlyChan represents a
// send-only
// channel
type SplitFuncSOnlyChan interface {
	ProvideSplitFunc(dat bufio.SplitFunc) // the send function - aka "MyKind <- some SplitFunc"
}
