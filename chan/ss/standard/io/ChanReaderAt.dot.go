// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReaderAtChan returns a new open channel
// (simply a 'chan io.ReaderAt' that is).
//
// Note: No 'ReaderAt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderAtPipelineStartsHere := MakeReaderAtChan()
//	// ... lot's of code to design and build Your favourite "myReaderAtWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderAtPipelineStartsHere <- drop
//	}
//	close(myReaderAtPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderAtBuffer) the channel is unbuffered.
//
func MakeReaderAtChan() (out chan io.ReaderAt) {
	return make(chan io.ReaderAt)
}

func sendReaderAt(out chan<- io.ReaderAt, inp ...io.ReaderAt) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReaderAt returns a channel to receive all inputs before close.
func ChanReaderAt(inp ...io.ReaderAt) (out <-chan io.ReaderAt) {
	cha := make(chan io.ReaderAt)
	go sendReaderAt(cha, inp...)
	return cha
}

func sendReaderAtSlice(out chan<- io.ReaderAt, inp ...[]io.ReaderAt) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReaderAtSlice returns a channel to receive all inputs before close.
func ChanReaderAtSlice(inp ...[]io.ReaderAt) (out <-chan io.ReaderAt) {
	cha := make(chan io.ReaderAt)
	go sendReaderAtSlice(cha, inp...)
	return cha
}

func joinReaderAt(done chan<- struct{}, out chan<- io.ReaderAt, inp ...io.ReaderAt) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReaderAt
func JoinReaderAt(out chan<- io.ReaderAt, inp ...io.ReaderAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderAt(cha, out, inp...)
	return cha
}

func joinReaderAtSlice(done chan<- struct{}, out chan<- io.ReaderAt, inp ...[]io.ReaderAt) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReaderAtSlice
func JoinReaderAtSlice(out chan<- io.ReaderAt, inp ...[]io.ReaderAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderAtSlice(cha, out, inp...)
	return cha
}

func joinReaderAtChan(done chan<- struct{}, out chan<- io.ReaderAt, inp <-chan io.ReaderAt) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReaderAtChan
func JoinReaderAtChan(out chan<- io.ReaderAt, inp <-chan io.ReaderAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderAtChan(cha, out, inp)
	return cha
}

func doitReaderAt(done chan<- struct{}, inp <-chan io.ReaderAt) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReaderAt returns a channel to receive one signal before close after inp has been drained.
func DoneReaderAt(inp <-chan io.ReaderAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReaderAt(cha, inp)
	return cha
}

func doitReaderAtSlice(done chan<- ([]io.ReaderAt), inp <-chan io.ReaderAt) {
	defer close(done)
	ReaderAtS := []io.ReaderAt{}
	for i := range inp {
		ReaderAtS = append(ReaderAtS, i)
	}
	done <- ReaderAtS
}

// DoneReaderAtSlice returns a channel which will receive a slice
// of all the ReaderAts received on inp channel before close.
// Unlike DoneReaderAt, a full slice is sent once, not just an event.
func DoneReaderAtSlice(inp <-chan io.ReaderAt) (done <-chan ([]io.ReaderAt)) {
	cha := make(chan ([]io.ReaderAt))
	go doitReaderAtSlice(cha, inp)
	return cha
}

func doitReaderAtFunc(done chan<- struct{}, inp <-chan io.ReaderAt, act func(a io.ReaderAt)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReaderAtFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderAtFunc(inp <-chan io.ReaderAt, act func(a io.ReaderAt)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReaderAt) { return }
	}
	go doitReaderAtFunc(cha, inp, act)
	return cha
}

func pipeReaderAtBuffer(out chan<- io.ReaderAt, inp <-chan io.ReaderAt) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReaderAtBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderAtBuffer(inp <-chan io.ReaderAt, cap int) (out <-chan io.ReaderAt) {
	cha := make(chan io.ReaderAt, cap)
	go pipeReaderAtBuffer(cha, inp)
	return cha
}

func pipeReaderAtFunc(out chan<- io.ReaderAt, inp <-chan io.ReaderAt, act func(a io.ReaderAt) io.ReaderAt) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReaderAtFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderAtMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderAtFunc(inp <-chan io.ReaderAt, act func(a io.ReaderAt) io.ReaderAt) (out <-chan io.ReaderAt) {
	cha := make(chan io.ReaderAt)
	if act == nil {
		act = func(a io.ReaderAt) io.ReaderAt { return a }
	}
	go pipeReaderAtFunc(cha, inp, act)
	return cha
}

func pipeReaderAtFork(out1, out2 chan<- io.ReaderAt, inp <-chan io.ReaderAt) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReaderAtFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderAtFork(inp <-chan io.ReaderAt) (out1, out2 <-chan io.ReaderAt) {
	cha1 := make(chan io.ReaderAt)
	cha2 := make(chan io.ReaderAt)
	go pipeReaderAtFork(cha1, cha2, inp)
	return cha1, cha2
}
