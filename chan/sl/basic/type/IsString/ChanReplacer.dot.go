// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

type ReplacerChan interface { // bidirectional channel
	ReplacerROnlyChan // aka "<-chan" - receive only
	ReplacerSOnlyChan // aka "chan<-" - send only
}

type ReplacerROnlyChan interface { // receive-only channel
	RequestReplacer() (dat *strings.Replacer)        // the receive function - aka "some-new-Replacer-var := <-MyKind"
	TryReplacer() (dat *strings.Replacer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Replacer-var, ok := <-MyKind"
}

type ReplacerSOnlyChan interface { // send-only channel
	ProvideReplacer(dat *strings.Replacer) // the send function - aka "MyKind <- some Replacer"
}
