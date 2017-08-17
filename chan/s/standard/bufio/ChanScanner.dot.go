// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type ScannerChan interface { // bidirectional channel
	ScannerROnlyChan // aka "<-chan" - receive only
	ScannerSOnlyChan // aka "chan<-" - send only
}

type ScannerROnlyChan interface { // receive-only channel
	RequestScanner() (dat *bufio.Scanner)        // the receive function - aka "some-new-Scanner-var := <-MyKind"
	TryScanner() (dat *bufio.Scanner, open bool) // the multi-valued comma-ok receive function - aka "some-new-Scanner-var, ok := <-MyKind"
}

type ScannerSOnlyChan interface { // send-only channel
	ProvideScanner(dat *bufio.Scanner) // the send function - aka "MyKind <- some Scanner"
}

type SChScanner struct { // supply channel
	dat chan *bufio.Scanner
	// req chan struct{}
}

func MakeSupplyScannerChan() *SChScanner {
	d := new(SChScanner)
	d.dat = make(chan *bufio.Scanner)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyScannerBuff(cap int) *SChScanner {
	d := new(SChScanner)
	d.dat = make(chan *bufio.Scanner, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideScanner is the send function - aka "MyKind <- some Scanner"
func (c *SChScanner) ProvideScanner(dat *bufio.Scanner) {
	// .req
	c.dat <- dat
}

// RequestScanner is the receive function - aka "some Scanner <- MyKind"
func (c *SChScanner) RequestScanner() (dat *bufio.Scanner) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryScanner is the comma-ok multi-valued form of RequestScanner and
// reports whether a received value was sent before the Scanner channel was closed.
func (c *SChScanner) TryScanner() (dat *bufio.Scanner, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
