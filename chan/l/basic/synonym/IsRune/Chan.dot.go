// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Chan interface { // bidirectional channel
	ROnlyChan // aka "<-chan" - receive only
	SOnlyChan // aka "chan<-" - send only
}

type ROnlyChan interface { // receive-only channel
	Request() (dat rune)        // the receive function - aka "some-new--var := <-MyKind"
	Try() (dat rune, open bool) // the multi-valued comma-ok receive function - aka "some-new--var, ok := <-MyKind"
}

type SOnlyChan interface { // send-only channel
	Provide(dat rune) // the send function - aka "MyKind <- some "
}

type DCh struct { // demand channel
	dat chan rune
	req chan struct{}
}

func MakeDemandChan() *DCh {
	d := new(DCh)
	d.dat = make(chan rune)
	d.req = make(chan struct{})
	return d
}

func MakeDemandBuff(cap int) *DCh {
	d := new(DCh)
	d.dat = make(chan rune, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// Provide is the send function - aka "MyKind <- some "
func (c *DCh) Provide(dat rune) {
	<-c.req
	c.dat <- dat
}

// Request is the receive function - aka "some  <- MyKind"
func (c *DCh) Request() (dat rune) {
	c.req <- struct{}{}
	return <-c.dat
}

// Try is the comma-ok multi-valued form of Request and
// reports whether a received value was sent before the  channel was closed.
func (c *DCh) Try() (dat rune, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
