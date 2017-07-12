// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadCloserChan returns a new open channel
// (simply a 'chan io.ReadCloser' that is).
//
// Note: No 'ReadCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadCloserPipelineStartsHere := MakeReadCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadCloserPipelineStartsHere <- drop
//	}
//	close(myReadCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadCloserBuffer) the channel is unbuffered.
//
func MakeReadCloserChan() chan io.ReadCloser {
	return make(chan io.ReadCloser)
}

// ChanReadCloser returns a channel to receive all inputs before close.
func ChanReadCloser(inp ...io.ReadCloser) chan io.ReadCloser {
	out := make(chan io.ReadCloser)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReadCloserSlice returns a channel to receive all inputs before close.
func ChanReadCloserSlice(inp ...[]io.ReadCloser) chan io.ReadCloser {
	out := make(chan io.ReadCloser)
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

// JoinReadCloser
func JoinReadCloser(out chan<- io.ReadCloser, inp ...io.ReadCloser) chan struct{} {
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

// JoinReadCloserSlice
func JoinReadCloserSlice(out chan<- io.ReadCloser, inp ...[]io.ReadCloser) chan struct{} {
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

// JoinReadCloserChan
func JoinReadCloserChan(out chan<- io.ReadCloser, inp <-chan io.ReadCloser) chan struct{} {
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

// DoneReadCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadCloser(inp <-chan io.ReadCloser) chan struct{} {
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

// DoneReadCloserSlice returns a channel which will receive a slice
// of all the ReadClosers received on inp channel before close.
// Unlike DoneReadCloser, a full slice is sent once, not just an event.
func DoneReadCloserSlice(inp <-chan io.ReadCloser) chan []io.ReadCloser {
	done := make(chan []io.ReadCloser)
	go func() {
		defer close(done)
		ReadCloserS := []io.ReadCloser{}
		for i := range inp {
			ReadCloserS = append(ReadCloserS, i)
		}
		done <- ReadCloserS
	}()
	return done
}

// DoneReadCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadCloserFunc(inp <-chan io.ReadCloser, act func(a io.ReadCloser)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReadCloser) { return }
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

// PipeReadCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadCloserBuffer(inp <-chan io.ReadCloser, cap int) chan io.ReadCloser {
	out := make(chan io.ReadCloser, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadCloserFunc(inp <-chan io.ReadCloser, act func(a io.ReadCloser) io.ReadCloser) chan io.ReadCloser {
	out := make(chan io.ReadCloser)
	if act == nil {
		act = func(a io.ReadCloser) io.ReadCloser { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadCloserFork(inp <-chan io.ReadCloser) (chan io.ReadCloser, chan io.ReadCloser) {
	out1 := make(chan io.ReadCloser)
	out2 := make(chan io.ReadCloser)
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
