// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ByteScannerChan interface { // bidirectional channel
	ByteScannerROnlyChan // aka "<-chan" - receive only
	ByteScannerSOnlyChan // aka "chan<-" - send only
}

type ByteScannerROnlyChan interface { // receive-only channel
	RequestByteScanner() (dat io.ByteScanner)        // the receive function - aka "some-new-ByteScanner-var := <-MyKind"
	TryByteScanner() (dat io.ByteScanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-ByteScanner-var, ok := <-MyKind"
}

type ByteScannerSOnlyChan interface { // send-only channel
	ProvideByteScanner(dat io.ByteScanner) // the send function - aka "MyKind <- some ByteScanner"
}

type SChByteScanner struct { // supply channel
	dat chan io.ByteScanner
	// req chan struct{}
}

func MakeSupplyByteScannerChan() *SChByteScanner {
	d := new(SChByteScanner)
	d.dat = make(chan io.ByteScanner)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyByteScannerBuff(cap int) *SChByteScanner {
	d := new(SChByteScanner)
	d.dat = make(chan io.ByteScanner, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideByteScanner is the send function - aka "MyKind <- some ByteScanner"
func (c *SChByteScanner) ProvideByteScanner(dat io.ByteScanner) {
	// .req
	c.dat <- dat
}

// RequestByteScanner is the receive function - aka "some ByteScanner <- MyKind"
func (c *SChByteScanner) RequestByteScanner() (dat io.ByteScanner) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryByteScanner is the comma-ok multi-valued form of RequestByteScanner and
// reports whether a received value was sent before the ByteScanner channel was closed.
func (c *SChByteScanner) TryByteScanner() (dat io.ByteScanner, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len