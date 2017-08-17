// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

type ReplacerChan interface { // bidirectional channel
	ReplacerROnlyChan // aka "<-chan" - receive only
	ReplacerSOnlyChan // aka "chan<-" - send only
}

type ReplacerROnlyChan interface { // receive-only channel
	RequestReplacer() (dat *strings.Replacer)        // the receive function - aka "some-new-Replacer-var := <-MyKind"
	TryReplacer() (dat *strings.Replacer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Replacer-var, ok := <-MyKind"
}

type ReplacerSOnlyChan interface { // send-only channel
	ProvideReplacer(dat *strings.Replacer) // the send function - aka "MyKind <- some Replacer"
}

type DChReplacer struct { // demand channel
	dat chan *strings.Replacer
	req chan struct{}
}

func MakeDemandReplacerChan() *DChReplacer {
	d := new(DChReplacer)
	d.dat = make(chan *strings.Replacer)
	d.req = make(chan struct{})
	return d
}

func MakeDemandReplacerBuff(cap int) *DChReplacer {
	d := new(DChReplacer)
	d.dat = make(chan *strings.Replacer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReplacer is the send function - aka "MyKind <- some Replacer"
func (c *DChReplacer) ProvideReplacer(dat *strings.Replacer) {
	<-c.req
	c.dat <- dat
}

// RequestReplacer is the receive function - aka "some Replacer <- MyKind"
func (c *DChReplacer) RequestReplacer() (dat *strings.Replacer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReplacer is the comma-ok multi-valued form of RequestReplacer and
// reports whether a received value was sent before the Replacer channel was closed.
func (c *DChReplacer) TryReplacer() (dat *strings.Replacer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
