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

type DChScanner struct { // demand channel
	dat chan *bufio.Scanner
	req chan struct{}
}

func MakeDemandScannerChan() *DChScanner {
	d := new(DChScanner)
	d.dat = make(chan *bufio.Scanner)
	d.req = make(chan struct{})
	return d
}

func MakeDemandScannerBuff(cap int) *DChScanner {
	d := new(DChScanner)
	d.dat = make(chan *bufio.Scanner, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideScanner is the send function - aka "MyKind <- some Scanner"
func (c *DChScanner) ProvideScanner(dat *bufio.Scanner) {
	<-c.req
	c.dat <- dat
}

// RequestScanner is the receive function - aka "some Scanner <- MyKind"
func (c *DChScanner) RequestScanner() (dat *bufio.Scanner) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryScanner is the comma-ok multi-valued form of RequestScanner and
// reports whether a received value was sent before the Scanner channel was closed.
func (c *DChScanner) TryScanner() (dat *bufio.Scanner, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
