// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// RuneSChan represents a
// bidirectional
// channel
type RuneSChan interface {
	RuneSROnlyChan // aka "<-chan" - receive only
	RuneSSOnlyChan // aka "chan<-" - send only
}

// RuneSROnlyChan represents a
// receive-only
// channel
type RuneSROnlyChan interface {
	RequestRuneS() (dat []rune)        // the receive function - aka "MyRuneS := <-MyRuneSROnlyChan"
	TryRuneS() (dat []rune, open bool) // the multi-valued comma-ok receive function - aka "MyRuneS, ok := <-MyRuneSROnlyChan"
}

// RuneSSOnlyChan represents a
// send-only
// channel
type RuneSSOnlyChan interface {
	ProvideRuneS(dat []rune) // the send function - aka "MyKind <- some RuneS"
}
