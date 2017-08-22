// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: SendProxyFsInfoS imports "container/ring" for the expanding buffer.
import (
	"container/ring"
	"github.com/golangsam/container/ccsafe/fs"
)

// FsInfoSCAP is the capacity of the buffered proxy channel
const FsInfoSCAP = 10

// FsInfoSQUE is the allocated size of the circular queue
const FsInfoSQUE = 16

// SendProxyFsInfoS returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//
// Note: the expanding buffer is implemented via "container/ring"
func SendProxyFsInfoS(out chan<- fs.FsInfoS) chan<- fs.FsInfoS {
	proxy := make(chan fs.FsInfoS, FsInfoSCAP)
	go func() {
		n := FsInfoSQUE // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- fs.FsInfoS
		var e fs.FsInfoS
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(fs.FsInfoS)
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

/* usage as found in $GOROOT/test/chan/sieve2.go
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
