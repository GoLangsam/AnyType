// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// ReplacerChan represents a
// bidirectional
// channel
type ReplacerChan interface {
	ReplacerROnlyChan // aka "<-chan" - receive only
	ReplacerSOnlyChan // aka "chan<-" - send only
}

// ReplacerROnlyChan represents a
// receive-only
// channel
type ReplacerROnlyChan interface {
	RequestReplacer() (dat *strings.Replacer)        // the receive function - aka "MyReplacer := <-MyReplacerROnlyChan"
	TryReplacer() (dat *strings.Replacer, open bool) // the multi-valued comma-ok receive function - aka "MyReplacer, ok := <-MyReplacerROnlyChan"
}

// ReplacerSOnlyChan represents a
// send-only
// channel
type ReplacerSOnlyChan interface {
	ProvideReplacer(dat *strings.Replacer) // the send function - aka "MyKind <- some Replacer"
}
