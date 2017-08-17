// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: SendProxyFsFold uses "container/ring"

import (
	"container/ring"
	"github.com/golangsam/container/ccsafe/fs"
)

/* usage as found in go/test/chan/sieve2.go
func Sieve() {
	// ...
	primes := make(chan int, 10)
	primes <- 3
	// ...
	go func() {
		// In order to generate the nth prime we only need multiples of primes ≤ sqrt(nth prime).
		// Thus, the merging goroutine will receive from 'primes' much slower than this goroutine will send to it,
		// making the buffer accumulate and block this goroutine from sending, causing a deadlock.
		// The solution is to use a proxy goroutine to do automatic buffering.
		primes := sendproxy(primes)
		// ...

	}()
}
*/

// FsFoldCAP is the capacity of the buffered proxy channel
const FsFoldCAP = 10

// FsFoldQUE is the allocated size of the circular queue
const FsFoldQUE = 16

// SendProxyFsFold returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//
// Note: the expanding buffer is implemented via "container/ring"
func SendProxyFsFold(out chan<- *fs.FsFold) chan<- *fs.FsFold {
	proxy := make(chan *fs.FsFold, FsFoldCAP)
	go func() {
		n := FsFoldQUE // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- *fs.FsFold
		var e *fs.FsFold
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(*fs.FsFold)
			}
			select {
			case e = <-proxy:
				last.Value = e
				if last.Next() == first {
					// buffer full: expand it
					last.Link(ring.New(n))
					n *= 2
				}
				last = last.Next()
			case c <- e:
				first = first.Next()
			}
		}
	}()
	return proxy
}
