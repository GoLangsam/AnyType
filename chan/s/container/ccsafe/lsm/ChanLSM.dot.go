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

type SChLSM struct { // supply channel
	dat chan lsm.LazyStringerMap
	// req chan struct{}
}

func MakeSupplyLSMChan() *SChLSM {
	d := new(SChLSM)
	d.dat = make(chan lsm.LazyStringerMap)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyLSMBuff(cap int) *SChLSM {
	d := new(SChLSM)
	d.dat = make(chan lsm.LazyStringerMap, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideLSM is the send function - aka "MyKind <- some LSM"
func (c *SChLSM) ProvideLSM(dat lsm.LazyStringerMap) {
	// .req
	c.dat <- dat
}

// RequestLSM is the receive function - aka "some LSM <- MyKind"
func (c *SChLSM) RequestLSM() (dat lsm.LazyStringerMap) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryLSM is the comma-ok multi-valued form of RequestLSM and
// reports whether a received value was sent before the LSM channel was closed.
func (c *SChLSM) TryLSM() (dat lsm.LazyStringerMap, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
