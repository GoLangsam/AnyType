// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ElementChan interface { // bidirectional channel
	ElementROnlyChan // aka "<-chan" - receive only
	ElementSOnlyChan // aka "chan<-" - send only
}

type ElementROnlyChan interface { // receive-only channel
	RequestElement() (dat list.Element)        // the receive function - aka "some-new-Element-var := <-MyKind"
	TryElement() (dat list.Element, open bool) // the multi-valued comma-ok receive function - aka "some-new-Element-var, ok := <-MyKind"
}

type ElementSOnlyChan interface { // send-only channel
	ProvideElement(dat list.Element) // the send function - aka "MyKind <- some Element"
}
