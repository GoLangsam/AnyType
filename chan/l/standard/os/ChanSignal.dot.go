// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

type SignalChan interface { // bidirectional channel
	SignalROnlyChan // aka "<-chan" - receive only
	SignalSOnlyChan // aka "chan<-" - send only
}

type SignalROnlyChan interface { // receive-only channel
	RequestSignal() (dat os.Signal)        // the receive function - aka "some-new-Signal-var := <-MyKind"
	TrySignal() (dat os.Signal, open bool) // the multi-valued comma-ok receive function - aka "some-new-Signal-var, ok := <-MyKind"
}

type SignalSOnlyChan interface { // send-only channel
	ProvideSignal(dat os.Signal) // the send function - aka "MyKind <- some Signal"
}

type DChSignal struct { // demand channel
	dat chan os.Signal
	req chan struct{}
}

func MakeDemandSignalChan() *DChSignal {
	d := new(DChSignal)
	d.dat = make(chan os.Signal)
	d.req = make(chan struct{})
	return d
}

func MakeDemandSignalBuff(cap int) *DChSignal {
	d := new(DChSignal)
	d.dat = make(chan os.Signal, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSignal is the send function - aka "MyKind <- some Signal"
func (c *DChSignal) ProvideSignal(dat os.Signal) {
	<-c.req
	c.dat <- dat
}

// RequestSignal is the receive function - aka "some Signal <- MyKind"
func (c *DChSignal) RequestSignal() (dat os.Signal) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySignal is the comma-ok multi-valued form of RequestSignal and
// reports whether a received value was sent before the Signal channel was closed.
func (c *DChSignal) TrySignal() (dat os.Signal, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
