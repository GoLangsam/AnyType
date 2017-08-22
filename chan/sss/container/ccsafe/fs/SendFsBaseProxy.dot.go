// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: SendProxyFsBase imports "container/ring" for the expanding buffer.
import (
	"container/ring"
	"github.com/golangsam/container/ccsafe/fs"
)

// FsBaseCAP is the capacity of the buffered proxy channel
const FsBaseCAP = 10

// FsBaseQUE is the allocated size of the circular queue
const FsBaseQUE = 16

// SendProxyFsBase returns a channel to serve as a sending proxy to 'out'.
// Uses a goroutine to receive values from 'out' and store them
// in an expanding buffer, so that sending to 'out' never blocks.
//
// Note: the expanding buffer is implemented via "container/ring"
func SendProxyFsBase(out chan<- *fs.FsBase) chan<- *fs.FsBase {
	proxy := make(chan *fs.FsBase, FsBaseCAP)
	go func() {
		n := FsBaseQUE // the allocated size of the circular queue
		first := ring.New(n)
		last := first
		var c chan<- *fs.FsBase
		var e *fs.FsBase
		for {
			c = out
			if first == last {
				// buffer empty: disable output
				c = nil
			} else {
				e = first.Value.(*fs.FsBase)
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
