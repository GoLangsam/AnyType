// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int8Chan represents a
// bidirectional
// channel
type Int8Chan interface {
	Int8ROnlyChan // aka "<-chan" - receive only
	Int8SOnlyChan // aka "chan<-" - send only
}

// Int8ROnlyChan represents a
// receive-only
// channel
type Int8ROnlyChan interface {
	RequestInt8() (dat int8)        // the receive function - aka "MyInt8 := <-MyInt8ROnlyChan"
	TryInt8() (dat int8, open bool) // the multi-valued comma-ok receive function - aka "MyInt8, ok := <-MyInt8ROnlyChan"
}

// Int8SOnlyChan represents a
// send-only
// channel
type Int8SOnlyChan interface {
	ProvideInt8(dat int8) // the send function - aka "MyKind <- some Int8"
}

// MergeInt8 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int8channel is returned.
func MergeInt8(inps ...<-chan int8) (out <-chan int8) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int8)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt82(inps[0], MergeInt8(inps[1:]...))
	}
}

// mergeInt82 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt82(i1, i2 <-chan int8) (out <-chan int8) {
	cha := make(chan int8)
	go func(out chan<- int8, i1, i2 <-chan int8) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 int8 // what we've read
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
