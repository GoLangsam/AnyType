// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt64Chan interface { // bidirectional channel
	UInt64ROnlyChan // aka "<-chan" - receive only
	UInt64SOnlyChan // aka "chan<-" - send only
}

type UInt64ROnlyChan interface { // receive-only channel
	RequestUInt64() (dat uint64)        // the receive function - aka "some-new-UInt64-var := <-MyKind"
	TryUInt64() (dat uint64, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt64-var, ok := <-MyKind"
}

type UInt64SOnlyChan interface { // send-only channel
	ProvideUInt64(dat uint64) // the send function - aka "MyKind <- some UInt64"
}

// MergeUInt642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeUInt642(i1, i2 <-chan uint64) (out <-chan uint64) {
	cha := make(chan uint64)
	go func(out chan<- uint64, i1, i2 <-chan uint64) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint64 // what we've read
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
