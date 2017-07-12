// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriteSeekerChan returns a new open channel
// (simply a 'chan io.ReadWriteSeeker' that is).
//
// Note: No 'ReadWriteSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriteSeekerPipelineStartsHere := MakeReadWriteSeekerChan()
//	// ... lot's of code to design and build Your favourite "myReadWriteSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriteSeekerPipelineStartsHere <- drop
//	}
//	close(myReadWriteSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriteSeekerBuffer) the channel is unbuffered.
//
func MakeReadWriteSeekerChan() chan io.ReadWriteSeeker {
	return make(chan io.ReadWriteSeeker)
}

// ChanReadWriteSeeker returns a channel to receive all inputs before close.
func ChanReadWriteSeeker(inp ...io.ReadWriteSeeker) chan io.ReadWriteSeeker {
	out := make(chan io.ReadWriteSeeker)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReadWriteSeekerSlice returns a channel to receive all inputs before close.
func ChanReadWriteSeekerSlice(inp ...[]io.ReadWriteSeeker) chan io.ReadWriteSeeker {
	out := make(chan io.ReadWriteSeeker)
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

// JoinReadWriteSeeker
func JoinReadWriteSeeker(out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) chan struct{} {
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

// JoinReadWriteSeekerSlice
func JoinReadWriteSeekerSlice(out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) chan struct{} {
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

// JoinReadWriteSeekerChan
func JoinReadWriteSeekerChan(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) chan struct{} {
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

// DoneReadWriteSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteSeeker(inp <-chan io.ReadWriteSeeker) chan struct{} {
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

// DoneReadWriteSeekerSlice returns a channel which will receive a slice
// of all the ReadWriteSeekers received on inp channel before close.
// Unlike DoneReadWriteSeeker, a full slice is sent once, not just an event.
func DoneReadWriteSeekerSlice(inp <-chan io.ReadWriteSeeker) chan []io.ReadWriteSeeker {
	done := make(chan []io.ReadWriteSeeker)
	go func() {
		defer close(done)
		ReadWriteSeekerS := []io.ReadWriteSeeker{}
		for i := range inp {
			ReadWriteSeekerS = append(ReadWriteSeekerS, i)
		}
		done <- ReadWriteSeekerS
	}()
	return done
}

// DoneReadWriteSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteSeeker) { return }
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

// PipeReadWriteSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteSeekerBuffer(inp <-chan io.ReadWriteSeeker, cap int) chan io.ReadWriteSeeker {
	out := make(chan io.ReadWriteSeeker, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadWriteSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker) io.ReadWriteSeeker) chan io.ReadWriteSeeker {
	out := make(chan io.ReadWriteSeeker)
	if act == nil {
		act = func(a io.ReadWriteSeeker) io.ReadWriteSeeker { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadWriteSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteSeekerFork(inp <-chan io.ReadWriteSeeker) (chan io.ReadWriteSeeker, chan io.ReadWriteSeeker) {
	out1 := make(chan io.ReadWriteSeeker)
	out2 := make(chan io.ReadWriteSeeker)
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
