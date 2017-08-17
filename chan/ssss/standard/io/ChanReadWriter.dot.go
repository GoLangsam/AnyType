// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriterChan returns a new open channel
// (simply a 'chan io.ReadWriter' that is).
//
// Note: No 'ReadWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriterPipelineStartsHere := MakeReadWriterChan()
//	// ... lot's of code to design and build Your favourite "myReadWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriterPipelineStartsHere <- drop
//	}
//	close(myReadWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriterBuffer) the channel is unbuffered.
//
func MakeReadWriterChan() chan io.ReadWriter {
	return make(chan io.ReadWriter)
}

// ChanReadWriter returns a channel to receive all inputs before close.
func ChanReadWriter(inp ...io.ReadWriter) chan io.ReadWriter {
	out := make(chan io.ReadWriter)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReadWriterSlice returns a channel to receive all inputs before close.
func ChanReadWriterSlice(inp ...[]io.ReadWriter) chan io.ReadWriter {
	out := make(chan io.ReadWriter)
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

// JoinReadWriter
func JoinReadWriter(out chan<- io.ReadWriter, inp ...io.ReadWriter) chan struct{} {
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

// JoinReadWriterSlice
func JoinReadWriterSlice(out chan<- io.ReadWriter, inp ...[]io.ReadWriter) chan struct{} {
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

// JoinReadWriterChan
func JoinReadWriterChan(out chan<- io.ReadWriter, inp <-chan io.ReadWriter) chan struct{} {
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

// DoneReadWriter returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriter(inp <-chan io.ReadWriter) chan struct{} {
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

// DoneReadWriterSlice returns a channel which will receive a slice
// of all the ReadWriters received on inp channel before close.
// Unlike DoneReadWriter, a full slice is sent once, not just an event.
func DoneReadWriterSlice(inp <-chan io.ReadWriter) chan []io.ReadWriter {
	done := make(chan []io.ReadWriter)
	go func() {
		defer close(done)
		ReadWriterS := []io.ReadWriter{}
		for i := range inp {
			ReadWriterS = append(ReadWriterS, i)
		}
		done <- ReadWriterS
	}()
	return done
}

// DoneReadWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriterFunc(inp <-chan io.ReadWriter, act func(a io.ReadWriter)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriter) { return }
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

// PipeReadWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriterBuffer(inp <-chan io.ReadWriter, cap int) chan io.ReadWriter {
	out := make(chan io.ReadWriter, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriterFunc(inp <-chan io.ReadWriter, act func(a io.ReadWriter) io.ReadWriter) chan io.ReadWriter {
	out := make(chan io.ReadWriter)
	if act == nil {
		act = func(a io.ReadWriter) io.ReadWriter { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriterFork(inp <-chan io.ReadWriter) (chan io.ReadWriter, chan io.ReadWriter) {
	out1 := make(chan io.ReadWriter)
	out2 := make(chan io.ReadWriter)
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
