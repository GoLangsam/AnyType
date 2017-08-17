// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/lsm"
)

type LSMChan interface { // bidirectional channel
	LSMROnlyChan // aka "<-chan" - receive only
	LSMSOnlyChan // aka "chan<-" - send only
}

type LSMROnlyChan interface { // receive-only channel
	RequestLSM() (dat lsm.LazyStringerMap)        // the receive function - aka "some-new-LSM-var := <-MyKind"
	TryLSM() (dat lsm.LazyStringerMap, open bool) // the multi-valued comma-ok receive function - aka "some-new-LSM-var, ok := <-MyKind"
}

type LSMSOnlyChan interface { // send-only channel
	ProvideLSM(dat lsm.LazyStringerMap) // the send function - aka "MyKind <- some LSM"
}

type DChLSM struct { // demand channel
	dat chan lsm.LazyStringerMap
	req chan struct{}
}

func MakeDemandLSMChan() *DChLSM {
	d := new(DChLSM)
	d.dat = make(chan lsm.LazyStringerMap)
	d.req = make(chan struct{})
	return d
}

func MakeDemandLSMBuff(cap int) *DChLSM {
	d := new(DChLSM)
	d.dat = make(chan lsm.LazyStringerMap, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideLSM is the send function - aka "MyKind <- some LSM"
func (c *DChLSM) ProvideLSM(dat lsm.LazyStringerMap) {
	<-c.req
	c.dat <- dat
}

// RequestLSM is the receive function - aka "some LSM <- MyKind"
func (c *DChLSM) RequestLSM() (dat lsm.LazyStringerMap) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryLSM is the comma-ok multi-valued form of RequestLSM and
// reports whether a received value was sent before the LSM channel was closed.
func (c *DChLSM) TryLSM() (dat lsm.LazyStringerMap, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
