// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeRuneReaderChan returns a new open channel
// (simply a 'chan io.RuneReader' that is).
//
// Note: No 'RuneReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myRuneReaderPipelineStartsHere := MakeRuneReaderChan()
//	// ... lot's of code to design and build Your favourite "myRuneReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myRuneReaderPipelineStartsHere <- drop
//	}
//	close(myRuneReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeRuneReaderBuffer) the channel is unbuffered.
//
func MakeRuneReaderChan() chan io.RuneReader {
	return make(chan io.RuneReader)
}

// ChanRuneReader returns a channel to receive all inputs before close.
func ChanRuneReader(inp ...io.RuneReader) chan io.RuneReader {
	out := make(chan io.RuneReader)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanRuneReaderSlice returns a channel to receive all inputs before close.
func ChanRuneReaderSlice(inp ...[]io.RuneReader) chan io.RuneReader {
	out := make(chan io.RuneReader)
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

// JoinRuneReader
func JoinRuneReader(out chan<- io.RuneReader, inp ...io.RuneReader) chan struct{} {
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

// JoinRuneReaderSlice
func JoinRuneReaderSlice(out chan<- io.RuneReader, inp ...[]io.RuneReader) chan struct{} {
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

// JoinRuneReaderChan
func JoinRuneReaderChan(out chan<- io.RuneReader, inp <-chan io.RuneReader) chan struct{} {
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

// DoneRuneReader returns a channel to receive one signal before close after inp has been drained.
func DoneRuneReader(inp <-chan io.RuneReader) chan struct{} {
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

// DoneRuneReaderSlice returns a channel which will receive a slice
// of all the RuneReaders received on inp channel before close.
// Unlike DoneRuneReader, a full slice is sent once, not just an event.
func DoneRuneReaderSlice(inp <-chan io.RuneReader) chan []io.RuneReader {
	done := make(chan []io.RuneReader)
	go func() {
		defer close(done)
		RuneReaderS := []io.RuneReader{}
		for i := range inp {
			RuneReaderS = append(RuneReaderS, i)
		}
		done <- RuneReaderS
	}()
	return done
}

// DoneRuneReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneRuneReaderFunc(inp <-chan io.RuneReader, act func(a io.RuneReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.RuneReader) { return }
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

// PipeRuneReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeRuneReaderBuffer(inp <-chan io.RuneReader, cap int) chan io.RuneReader {
	out := make(chan io.RuneReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeRuneReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeRuneReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeRuneReaderFunc(inp <-chan io.RuneReader, act func(a io.RuneReader) io.RuneReader) chan io.RuneReader {
	out := make(chan io.RuneReader)
	if act == nil {
		act = func(a io.RuneReader) io.RuneReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeRuneReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeRuneReaderFork(inp <-chan io.RuneReader) (chan io.RuneReader, chan io.RuneReader) {
	out1 := make(chan io.RuneReader)
	out2 := make(chan io.RuneReader)
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
