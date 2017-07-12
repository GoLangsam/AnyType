// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReaderFromChan returns a new open channel
// (simply a 'chan io.ReaderFrom' that is).
//
// Note: No 'ReaderFrom-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderFromPipelineStartsHere := MakeReaderFromChan()
//	// ... lot's of code to design and build Your favourite "myReaderFromWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderFromPipelineStartsHere <- drop
//	}
//	close(myReaderFromPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderFromBuffer) the channel is unbuffered.
//
func MakeReaderFromChan() (out chan io.ReaderFrom) {
	return make(chan io.ReaderFrom)
}

func sendReaderFrom(out chan<- io.ReaderFrom, inp ...io.ReaderFrom) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReaderFrom returns a channel to receive all inputs before close.
func ChanReaderFrom(inp ...io.ReaderFrom) (out <-chan io.ReaderFrom) {
	cha := make(chan io.ReaderFrom)
	go sendReaderFrom(cha, inp...)
	return cha
}

func sendReaderFromSlice(out chan<- io.ReaderFrom, inp ...[]io.ReaderFrom) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReaderFromSlice returns a channel to receive all inputs before close.
func ChanReaderFromSlice(inp ...[]io.ReaderFrom) (out <-chan io.ReaderFrom) {
	cha := make(chan io.ReaderFrom)
	go sendReaderFromSlice(cha, inp...)
	return cha
}

func joinReaderFrom(done chan<- struct{}, out chan<- io.ReaderFrom, inp ...io.ReaderFrom) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReaderFrom
func JoinReaderFrom(out chan<- io.ReaderFrom, inp ...io.ReaderFrom) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderFrom(cha, out, inp...)
	return cha
}

func joinReaderFromSlice(done chan<- struct{}, out chan<- io.ReaderFrom, inp ...[]io.ReaderFrom) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReaderFromSlice
func JoinReaderFromSlice(out chan<- io.ReaderFrom, inp ...[]io.ReaderFrom) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderFromSlice(cha, out, inp...)
	return cha
}

func joinReaderFromChan(done chan<- struct{}, out chan<- io.ReaderFrom, inp <-chan io.ReaderFrom) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReaderFromChan
func JoinReaderFromChan(out chan<- io.ReaderFrom, inp <-chan io.ReaderFrom) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderFromChan(cha, out, inp)
	return cha
}

func doitReaderFrom(done chan<- struct{}, inp <-chan io.ReaderFrom) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReaderFrom returns a channel to receive one signal before close after inp has been drained.
func DoneReaderFrom(inp <-chan io.ReaderFrom) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReaderFrom(cha, inp)
	return cha
}

func doitReaderFromSlice(done chan<- ([]io.ReaderFrom), inp <-chan io.ReaderFrom) {
	defer close(done)
	ReaderFromS := []io.ReaderFrom{}
	for i := range inp {
		ReaderFromS = append(ReaderFromS, i)
	}
	done <- ReaderFromS
}

// DoneReaderFromSlice returns a channel which will receive a slice
// of all the ReaderFroms received on inp channel before close.
// Unlike DoneReaderFrom, a full slice is sent once, not just an event.
func DoneReaderFromSlice(inp <-chan io.ReaderFrom) (done <-chan ([]io.ReaderFrom)) {
	cha := make(chan ([]io.ReaderFrom))
	go doitReaderFromSlice(cha, inp)
	return cha
}

func doitReaderFromFunc(done chan<- struct{}, inp <-chan io.ReaderFrom, act func(a io.ReaderFrom)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReaderFromFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderFromFunc(inp <-chan io.ReaderFrom, act func(a io.ReaderFrom)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReaderFrom) { return }
	}
	go doitReaderFromFunc(cha, inp, act)
	return cha
}

func pipeReaderFromBuffer(out chan<- io.ReaderFrom, inp <-chan io.ReaderFrom) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReaderFromBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderFromBuffer(inp <-chan io.ReaderFrom, cap int) (out <-chan io.ReaderFrom) {
	cha := make(chan io.ReaderFrom, cap)
	go pipeReaderFromBuffer(cha, inp)
	return cha
}

func pipeReaderFromFunc(out chan<- io.ReaderFrom, inp <-chan io.ReaderFrom, act func(a io.ReaderFrom) io.ReaderFrom) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReaderFromFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderFromMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderFromFunc(inp <-chan io.ReaderFrom, act func(a io.ReaderFrom) io.ReaderFrom) (out <-chan io.ReaderFrom) {
	cha := make(chan io.ReaderFrom)
	if act == nil {
		act = func(a io.ReaderFrom) io.ReaderFrom { return a }
	}
	go pipeReaderFromFunc(cha, inp, act)
	return cha
}

func pipeReaderFromFork(out1, out2 chan<- io.ReaderFrom, inp <-chan io.ReaderFrom) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReaderFromFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderFromFork(inp <-chan io.ReaderFrom) (out1, out2 <-chan io.ReaderFrom) {
	cha1 := make(chan io.ReaderFrom)
	cha2 := make(chan io.ReaderFrom)
	go pipeReaderFromFork(cha1, cha2, inp)
	return cha1, cha2
}
