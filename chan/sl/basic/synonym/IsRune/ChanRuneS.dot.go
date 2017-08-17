// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type RuneSChan interface { // bidirectional channel
	RuneSROnlyChan // aka "<-chan" - receive only
	RuneSSOnlyChan // aka "chan<-" - send only
}

type RuneSROnlyChan interface { // receive-only channel
	RequestRuneS() (dat []rune)        // the receive function - aka "some-new-RuneS-var := <-MyKind"
	TryRuneS() (dat []rune, open bool) // the multi-valued comma-ok receive function - aka "some-new-RuneS-var, ok := <-MyKind"
}

type RuneSSOnlyChan interface { // send-only channel
	ProvideRuneS(dat []rune) // the send function - aka "MyKind <- some RuneS"
}
