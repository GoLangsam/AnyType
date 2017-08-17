// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type SomeTypeChan interface { // bidirectional channel
	SomeTypeROnlyChan // aka "<-chan" - receive only
	SomeTypeSOnlyChan // aka "chan<-" - send only
}

type SomeTypeROnlyChan interface { // receive-only channel
	RequestSomeType() (dat SomeType)        // the receive function - aka "some-new-SomeType-var := <-MyKind"
	TrySomeType() (dat SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-SomeType-var, ok := <-MyKind"
}

type SomeTypeSOnlyChan interface { // send-only channel
	ProvideSomeType(dat SomeType) // the send function - aka "MyKind <- some SomeType"
}

type SChSomeType struct { // supply channel
	dat chan SomeType
	// req chan struct{}
}

func MakeSupplySomeTypeChan() *SChSomeType {
	d := new(SChSomeType)
	d.dat = make(chan SomeType)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplySomeTypeBuff(cap int) *SChSomeType {
	d := new(SChSomeType)
	d.dat = make(chan SomeType, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSomeType is the send function - aka "MyKind <- some SomeType"
func (c *SChSomeType) ProvideSomeType(dat SomeType) {
	// .req
	c.dat <- dat
}

// RequestSomeType is the receive function - aka "some SomeType <- MyKind"
func (c *SChSomeType) RequestSomeType() (dat SomeType) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySomeType is the comma-ok multi-valued form of RequestSomeType and
// reports whether a received value was sent before the SomeType channel was closed.
func (c *SChSomeType) TrySomeType() (dat SomeType, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
