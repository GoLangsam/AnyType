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

type SChRuneS struct { // supply channel
	dat chan []rune
	// req chan struct{}
}

func MakeSupplyRuneSChan() *SChRuneS {
	d := new(SChRuneS)
	d.dat = make(chan []rune)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyRuneSBuff(cap int) *SChRuneS {
	d := new(SChRuneS)
	d.dat = make(chan []rune, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideRuneS is the send function - aka "MyKind <- some RuneS"
func (c *SChRuneS) ProvideRuneS(dat []rune) {
	// .req
	c.dat <- dat
}

// RequestRuneS is the receive function - aka "some RuneS <- MyKind"
func (c *SChRuneS) RequestRuneS() (dat []rune) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryRuneS is the comma-ok multi-valued form of RequestRuneS and
// reports whether a received value was sent before the RuneS channel was closed.
func (c *SChRuneS) TryRuneS() (dat []rune, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
