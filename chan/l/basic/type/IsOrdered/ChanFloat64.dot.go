// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Float64Chan represents a
// bidirectional
// channel
type Float64Chan interface {
	Float64ROnlyChan // aka "<-chan" - receive only
	Float64SOnlyChan // aka "chan<-" - send only
}

// Float64ROnlyChan represents a
// receive-only
// channel
type Float64ROnlyChan interface {
	RequestFloat64() (dat float64)        // the receive function - aka "MyFloat64 := <-MyFloat64ROnlyChan"
	TryFloat64() (dat float64, open bool) // the multi-valued comma-ok receive function - aka "MyFloat64, ok := <-MyFloat64ROnlyChan"
}

// Float64SOnlyChan represents a
// send-only
// channel
type Float64SOnlyChan interface {
	ProvideFloat64(dat float64) // the send function - aka "MyKind <- some Float64"
}

// DChFloat64 is a demand channel
type DChFloat64 struct {
	dat chan float64
	req chan struct{}
}

// MakeDemandFloat64Chan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandFloat64Chan() *DChFloat64 {
	d := new(DChFloat64)
	d.dat = make(chan float64)
	d.req = make(chan struct{})
	return d
}

// MakeDemandFloat64Buff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandFloat64Buff(cap int) *DChFloat64 {
	d := new(DChFloat64)
	d.dat = make(chan float64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideFloat64 is the send function - aka "MyKind <- some Float64"
func (c *DChFloat64) ProvideFloat64(dat float64) {
	<-c.req
	c.dat <- dat
}

// RequestFloat64 is the receive function - aka "some Float64 <- MyKind"
func (c *DChFloat64) RequestFloat64() (dat float64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryFloat64 is the comma-ok multi-valued form of RequestFloat64 and
// reports whether a received value was sent before the Float64 channel was closed.
func (c *DChFloat64) TryFloat64() (dat float64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeFloat64 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Float64channel is returned.
func MergeFloat64(inps ...<-chan float64) (out <-chan float64) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan float64)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeFloat642(inps[0], MergeFloat64(inps[1:]...))
	}
}

// mergeFloat642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeFloat642(i1, i2 <-chan float64) (out <-chan float64) {
	cha := make(chan float64)
	go func(out chan<- float64, i1, i2 <-chan float64) {
		defer close(out)
		var (
			clos1, clos2 bool    // we found the chan closed
			buff1, buff2 bool    // we've read 'from', but not sent (yet)
			ok           bool    // did we read sucessfully?
			from1, from2 float64 // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
