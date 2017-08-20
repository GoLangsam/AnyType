// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// ElementChan represents a
// bidirectional
// channel
type ElementChan interface {
	ElementROnlyChan // aka "<-chan" - receive only
	ElementSOnlyChan // aka "chan<-" - send only
}

// ElementROnlyChan represents a
// receive-only
// channel
type ElementROnlyChan interface {
	RequestElement() (dat list.Element)        // the receive function - aka "MyElement := <-MyElementROnlyChan"
	TryElement() (dat list.Element, open bool) // the multi-valued comma-ok receive function - aka "MyElement, ok := <-MyElementROnlyChan"
}

// ElementSOnlyChan represents a
// send-only
// channel
type ElementSOnlyChan interface {
	ProvideElement(dat list.Element) // the send function - aka "MyKind <- some Element"
}
