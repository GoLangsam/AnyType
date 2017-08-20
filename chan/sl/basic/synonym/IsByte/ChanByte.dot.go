// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// ByteChan represents a
// bidirectional
// channel
type ByteChan interface {
	ByteROnlyChan // aka "<-chan" - receive only
	ByteSOnlyChan // aka "chan<-" - send only
}

// ByteROnlyChan represents a
// receive-only
// channel
type ByteROnlyChan interface {
	RequestByte() (dat byte)        // the receive function - aka "MyByte := <-MyByteROnlyChan"
	TryByte() (dat byte, open bool) // the multi-valued comma-ok receive function - aka "MyByte, ok := <-MyByteROnlyChan"
}

// ByteSOnlyChan represents a
// send-only
// channel
type ByteSOnlyChan interface {
	ProvideByte(dat byte) // the send function - aka "MyKind <- some Byte"
}
