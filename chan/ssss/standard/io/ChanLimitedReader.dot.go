// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeLimitedReaderChan returns a new open channel
// (simply a 'chan *io.LimitedReader' that is).
//
// Note: No 'LimitedReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLimitedReaderPipelineStartsHere := MakeLimitedReaderChan()
//	// ... lot's of code to design and build Your favourite "myLimitedReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLimitedReaderPipelineStartsHere <- drop
//	}
//	close(myLimitedReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLimitedReaderBuffer) the channel is unbuffered.
//
func MakeLimitedReaderChan() chan *io.LimitedReader {
	return make(chan *io.LimitedReader)
}

// ChanLimitedReader returns a channel to receive all inputs before close.
func ChanLimitedReader(inp ...*io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanLimitedReaderSlice returns a channel to receive all inputs before close.
func ChanLimitedReaderSlice(inp ...[]*io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
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

// JoinLimitedReader
func JoinLimitedReader(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) chan struct{} {
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

// JoinLimitedReaderSlice
func JoinLimitedReaderSlice(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) chan struct{} {
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

// JoinLimitedReaderChan
func JoinLimitedReaderChan(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) chan struct{} {
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

// DoneLimitedReader returns a channel to receive one signal before close after inp has been drained.
func DoneLimitedReader(inp <-chan *io.LimitedReader) chan struct{} {
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

// DoneLimitedReaderSlice returns a channel which will receive a slice
// of all the LimitedReaders received on inp channel before close.
// Unlike DoneLimitedReader, a full slice is sent once, not just an event.
func DoneLimitedReaderSlice(inp <-chan *io.LimitedReader) chan []*io.LimitedReader {
	done := make(chan []*io.LimitedReader)
	go func() {
		defer close(done)
		LimitedReaderS := []*io.LimitedReader{}
		for i := range inp {
			LimitedReaderS = append(LimitedReaderS, i)
		}
		done <- LimitedReaderS
	}()
	return done
}

// DoneLimitedReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *io.LimitedReader) { return }
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

// PipeLimitedReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLimitedReaderBuffer(inp <-chan *io.LimitedReader, cap int) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeLimitedReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLimitedReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
	if act == nil {
		act = func(a *io.LimitedReader) *io.LimitedReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeLimitedReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLimitedReaderFork(inp <-chan *io.LimitedReader) (chan *io.LimitedReader, chan *io.LimitedReader) {
	out1 := make(chan *io.LimitedReader)
	out2 := make(chan *io.LimitedReader)
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
