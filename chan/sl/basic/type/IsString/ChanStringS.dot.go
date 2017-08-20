// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// StringSChan represents a
// bidirectional
// channel
type StringSChan interface {
	StringSROnlyChan // aka "<-chan" - receive only
	StringSSOnlyChan // aka "chan<-" - send only
}

// StringSROnlyChan represents a
// receive-only
// channel
type StringSROnlyChan interface {
	RequestStringS() (dat []string)        // the receive function - aka "MyStringS := <-MyStringSROnlyChan"
	TryStringS() (dat []string, open bool) // the multi-valued comma-ok receive function - aka "MyStringS, ok := <-MyStringSROnlyChan"
}

// StringSSOnlyChan represents a
// send-only
// channel
type StringSSOnlyChan interface {
	ProvideStringS(dat []string) // the send function - aka "MyKind <- some StringS"
}
