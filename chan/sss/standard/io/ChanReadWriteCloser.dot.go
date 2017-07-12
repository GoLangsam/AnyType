// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriteCloserChan returns a new open channel
// (simply a 'chan io.ReadWriteCloser' that is).
//
// Note: No 'ReadWriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriteCloserPipelineStartsHere := MakeReadWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriteCloserPipelineStartsHere <- drop
//	}
//	close(myReadWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriteCloserBuffer) the channel is unbuffered.
//
func MakeReadWriteCloserChan() (out chan io.ReadWriteCloser) {
	return make(chan io.ReadWriteCloser)
}

// ChanReadWriteCloser returns a channel to receive all inputs before close.
func ChanReadWriteCloser(inp ...io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go func(out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanReadWriteCloserSlice returns a channel to receive all inputs before close.
func ChanReadWriteCloserSlice(inp ...[]io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go func(out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinReadWriteCloser
func JoinReadWriteCloser(out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReadWriteCloserSlice
func JoinReadWriteCloserSlice(out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReadWriteCloserChan
func JoinReadWriteCloserChan(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReadWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteCloser(inp <-chan io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.ReadWriteCloser) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReadWriteCloserSlice returns a channel which will receive a slice
// of all the ReadWriteClosers received on inp channel before close.
// Unlike DoneReadWriteCloser, a full slice is sent once, not just an event.
func DoneReadWriteCloserSlice(inp <-chan io.ReadWriteCloser) (done <-chan []io.ReadWriteCloser) {
	cha := make(chan []io.ReadWriteCloser)
	go func(inp <-chan io.ReadWriteCloser, done chan<- []io.ReadWriteCloser) {
		defer close(done)
		ReadWriteCloserS := []io.ReadWriteCloser{}
		for i := range inp {
			ReadWriteCloserS = append(ReadWriteCloserS, i)
		}
		done <- ReadWriteCloserS
	}(inp, cha)
	return cha
}

// DoneReadWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteCloser) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReadWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteCloserBuffer(inp <-chan io.ReadWriteCloser, cap int) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser, cap)
	go func(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReadWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser) io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	if act == nil {
		act = func(a io.ReadWriteCloser) io.ReadWriteCloser { return a }
	}
	go func(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser) io.ReadWriteCloser) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReadWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteCloserFork(inp <-chan io.ReadWriteCloser) (out1, out2 <-chan io.ReadWriteCloser) {
	cha1 := make(chan io.ReadWriteCloser)
	cha2 := make(chan io.ReadWriteCloser)
	go func(out1, out2 chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
