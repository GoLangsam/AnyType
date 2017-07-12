// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ListSChan interface { // bidirectional channel
	ListSROnlyChan // aka "<-chan" - receive only
	ListSSOnlyChan // aka "chan<-" - send only
}

type ListSROnlyChan interface { // receive-only channel
	RequestListS() (dat []list.List)        // the receive function - aka "some-new-ListS-var := <-MyKind"
	TryListS() (dat []list.List, open bool) // the multi-valued comma-ok receive function - aka "some-new-ListS-var, ok := <-MyKind"
}

type ListSSOnlyChan interface { // send-only channel
	ProvideListS(dat []list.List) // the send function - aka "MyKind <- some ListS"
}