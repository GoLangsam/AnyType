// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type RuneScannerChan interface { // bidirectional channel
	RuneScannerROnlyChan // aka "<-chan" - receive only
	RuneScannerSOnlyChan // aka "chan<-" - send only
}

type RuneScannerROnlyChan interface { // receive-only channel
	RequestRuneScanner() (dat io.RuneScanner)        // the receive function - aka "some-new-RuneScanner-var := <-MyKind"
	TryRuneScanner() (dat io.RuneScanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-RuneScanner-var, ok := <-MyKind"
}

type RuneScannerSOnlyChan interface { // send-only channel
	ProvideRuneScanner(dat io.RuneScanner) // the send function - aka "MyKind <- some RuneScanner"
}

type SChRuneScanner struct { // supply channel
	dat chan io.RuneScanner
	// req chan struct{}
}

func MakeSupplyRuneScannerChan() *SChRuneScanner {
	d := new(SChRuneScanner)
	d.dat = make(chan io.RuneScanner)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyRuneScannerBuff(cap int) *SChRuneScanner {
	d := new(SChRuneScanner)
	d.dat = make(chan io.RuneScanner, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideRuneScanner is the send function - aka "MyKind <- some RuneScanner"
func (c *SChRuneScanner) ProvideRuneScanner(dat io.RuneScanner) {
	// .req
	c.dat <- dat
}

// RequestRuneScanner is the receive function - aka "some RuneScanner <- MyKind"
func (c *SChRuneScanner) RequestRuneScanner() (dat io.RuneScanner) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryRuneScanner is the comma-ok multi-valued form of RequestRuneScanner and
// reports whether a received value was sent before the RuneScanner channel was closed.
func (c *SChRuneScanner) TryRuneScanner() (dat io.RuneScanner, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
