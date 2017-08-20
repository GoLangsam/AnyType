// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/tag"
)

// TagChan represents a
// bidirectional
// channel
type TagChan interface {
	TagROnlyChan // aka "<-chan" - receive only
	TagSOnlyChan // aka "chan<-" - send only
}

// TagROnlyChan represents a
// receive-only
// channel
type TagROnlyChan interface {
	RequestTag() (dat tag.TagAny)        // the receive function - aka "MyTag := <-MyTagROnlyChan"
	TryTag() (dat tag.TagAny, open bool) // the multi-valued comma-ok receive function - aka "MyTag, ok := <-MyTagROnlyChan"
}

// TagSOnlyChan represents a
// send-only
// channel
type TagSOnlyChan interface {
	ProvideTag(dat tag.TagAny) // the send function - aka "MyKind <- some Tag"
}
