// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// ListSChan represents a
// bidirectional
// channel
type ListSChan interface {
	ListSROnlyChan // aka "<-chan" - receive only
	ListSSOnlyChan // aka "chan<-" - send only
}

// ListSROnlyChan represents a
// receive-only
// channel
type ListSROnlyChan interface {
	RequestListS() (dat []list.List)        // the receive function - aka "MyListS := <-MyListSROnlyChan"
	TryListS() (dat []list.List, open bool) // the multi-valued comma-ok receive function - aka "MyListS, ok := <-MyListSROnlyChan"
}

// ListSSOnlyChan represents a
// send-only
// channel
type ListSSOnlyChan interface {
	ProvideListS(dat []list.List) // the send function - aka "MyKind <- some ListS"
}
