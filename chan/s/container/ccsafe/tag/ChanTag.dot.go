// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/tag"
)

type TagChan interface { // bidirectional channel
	TagROnlyChan // aka "<-chan" - receive only
	TagSOnlyChan // aka "chan<-" - send only
}

type TagROnlyChan interface { // receive-only channel
	RequestTag() (dat tag.TagAny)        // the receive function - aka "some-new-Tag-var := <-MyKind"
	TryTag() (dat tag.TagAny, open bool) // the multi-valued comma-ok receive function - aka "some-new-Tag-var, ok := <-MyKind"
}

type TagSOnlyChan interface { // send-only channel
	ProvideTag(dat tag.TagAny) // the send function - aka "MyKind <- some Tag"
}

type SChTag struct { // supply channel
	dat chan tag.TagAny
	// req chan struct{}
}

func MakeSupplyTagChan() *SChTag {
	d := new(SChTag)
	d.dat = make(chan tag.TagAny)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyTagBuff(cap int) *SChTag {
	d := new(SChTag)
	d.dat = make(chan tag.TagAny, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideTag is the send function - aka "MyKind <- some Tag"
func (c *SChTag) ProvideTag(dat tag.TagAny) {
	// .req
	c.dat <- dat
}

// RequestTag is the receive function - aka "some Tag <- MyKind"
func (c *SChTag) RequestTag() (dat tag.TagAny) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryTag is the comma-ok multi-valued form of RequestTag and
// reports whether a received value was sent before the Tag channel was closed.
func (c *SChTag) TryTag() (dat tag.TagAny, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
