// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteScannerChan returns a new open channel
// (simply a 'chan io.ByteScanner' that is).
//
// Note: No 'ByteScanner-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteScannerPipelineStartsHere := MakeByteScannerChan()
//	// ... lot's of code to design and build Your favourite "myByteScannerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteScannerPipelineStartsHere <- drop
//	}
//	close(myByteScannerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteScannerBuffer) the channel is unbuffered.
//
func MakeByteScannerChan() chan io.ByteScanner {
	return make(chan io.ByteScanner)
}

// ChanByteScanner returns a channel to receive all inputs before close.
func ChanByteScanner(inp ...io.ByteScanner) chan io.ByteScanner {
	out := make(chan io.ByteScanner)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanByteScannerSlice returns a channel to receive all inputs before close.
func ChanByteScannerSlice(inp ...[]io.ByteScanner) chan io.ByteScanner {
	out := make(chan io.ByteScanner)
	go func() {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}()
	return out
}

// JoinByteScanner
func JoinByteScanner(out chan<- io.ByteScanner, inp ...io.ByteScanner) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// JoinByteScannerSlice
func JoinByteScannerSlice(out chan<- io.ByteScanner, inp ...[]io.ByteScanner) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinByteScannerChan
func JoinByteScannerChan(out chan<- io.ByteScanner, inp <-chan io.ByteScanner) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneByteScanner returns a channel to receive one signal before close after inp has been drained.
func DoneByteScanner(inp <-chan io.ByteScanner) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneByteScannerSlice returns a channel which will receive a slice
// of all the ByteScanners received on inp channel before close.
// Unlike DoneByteScanner, a full slice is sent once, not just an event.
func DoneByteScannerSlice(inp <-chan io.ByteScanner) chan []io.ByteScanner {
	done := make(chan []io.ByteScanner)
	go func() {
		defer close(done)
		ByteScannerS := []io.ByteScanner{}
		for i := range inp {
			ByteScannerS = append(ByteScannerS, i)
		}
		done <- ByteScannerS
	}()
	return done
}

// DoneByteScannerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteScannerFunc(inp <-chan io.ByteScanner, act func(a io.ByteScanner)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ByteScanner) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeByteScannerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteScannerBuffer(inp <-chan io.ByteScanner, cap int) chan io.ByteScanner {
	out := make(chan io.ByteScanner, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeByteScannerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteScannerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteScannerFunc(inp <-chan io.ByteScanner, act func(a io.ByteScanner) io.ByteScanner) chan io.ByteScanner {
	out := make(chan io.ByteScanner)
	if act == nil {
		act = func(a io.ByteScanner) io.ByteScanner { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeByteScannerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteScannerFork(inp <-chan io.ByteScanner) (chan io.ByteScanner, chan io.ByteScanner) {
	out1 := make(chan io.ByteScanner)
	out2 := make(chan io.ByteScanner)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}
