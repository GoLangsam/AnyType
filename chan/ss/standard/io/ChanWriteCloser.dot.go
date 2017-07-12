// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriteCloserChan returns a new open channel
// (simply a 'chan io.WriteCloser' that is).
//
// Note: No 'WriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriteCloserPipelineStartsHere := MakeWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriteCloserPipelineStartsHere <- drop
//	}
//	close(myWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriteCloserBuffer) the channel is unbuffered.
//
func MakeWriteCloserChan() (out chan io.WriteCloser) {
	return make(chan io.WriteCloser)
}

func sendWriteCloser(out chan<- io.WriteCloser, inp ...io.WriteCloser) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanWriteCloser returns a channel to receive all inputs before close.
func ChanWriteCloser(inp ...io.WriteCloser) (out <-chan io.WriteCloser) {
	cha := make(chan io.WriteCloser)
	go sendWriteCloser(cha, inp...)
	return cha
}

func sendWriteCloserSlice(out chan<- io.WriteCloser, inp ...[]io.WriteCloser) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanWriteCloserSlice returns a channel to receive all inputs before close.
func ChanWriteCloserSlice(inp ...[]io.WriteCloser) (out <-chan io.WriteCloser) {
	cha := make(chan io.WriteCloser)
	go sendWriteCloserSlice(cha, inp...)
	return cha
}

func joinWriteCloser(done chan<- struct{}, out chan<- io.WriteCloser, inp ...io.WriteCloser) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriteCloser
func JoinWriteCloser(out chan<- io.WriteCloser, inp ...io.WriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriteCloser(cha, out, inp...)
	return cha
}

func joinWriteCloserSlice(done chan<- struct{}, out chan<- io.WriteCloser, inp ...[]io.WriteCloser) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinWriteCloserSlice
func JoinWriteCloserSlice(out chan<- io.WriteCloser, inp ...[]io.WriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriteCloserSlice(cha, out, inp...)
	return cha
}

func joinWriteCloserChan(done chan<- struct{}, out chan<- io.WriteCloser, inp <-chan io.WriteCloser) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriteCloserChan
func JoinWriteCloserChan(out chan<- io.WriteCloser, inp <-chan io.WriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriteCloserChan(cha, out, inp)
	return cha
}

func doitWriteCloser(done chan<- struct{}, inp <-chan io.WriteCloser) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneWriteCloser(inp <-chan io.WriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitWriteCloser(cha, inp)
	return cha
}

func doitWriteCloserSlice(done chan<- ([]io.WriteCloser), inp <-chan io.WriteCloser) {
	defer close(done)
	WriteCloserS := []io.WriteCloser{}
	for i := range inp {
		WriteCloserS = append(WriteCloserS, i)
	}
	done <- WriteCloserS
}

// DoneWriteCloserSlice returns a channel which will receive a slice
// of all the WriteClosers received on inp channel before close.
// Unlike DoneWriteCloser, a full slice is sent once, not just an event.
func DoneWriteCloserSlice(inp <-chan io.WriteCloser) (done <-chan ([]io.WriteCloser)) {
	cha := make(chan ([]io.WriteCloser))
	go doitWriteCloserSlice(cha, inp)
	return cha
}

func doitWriteCloserFunc(done chan<- struct{}, inp <-chan io.WriteCloser, act func(a io.WriteCloser)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.WriteCloser) { return }
	}
	go doitWriteCloserFunc(cha, inp, act)
	return cha
}

func pipeWriteCloserBuffer(out chan<- io.WriteCloser, inp <-chan io.WriteCloser) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriteCloserBuffer(inp <-chan io.WriteCloser, cap int) (out <-chan io.WriteCloser) {
	cha := make(chan io.WriteCloser, cap)
	go pipeWriteCloserBuffer(cha, inp)
	return cha
}

func pipeWriteCloserFunc(out chan<- io.WriteCloser, inp <-chan io.WriteCloser, act func(a io.WriteCloser) io.WriteCloser) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser) io.WriteCloser) (out <-chan io.WriteCloser) {
	cha := make(chan io.WriteCloser)
	if act == nil {
		act = func(a io.WriteCloser) io.WriteCloser { return a }
	}
	go pipeWriteCloserFunc(cha, inp, act)
	return cha
}

func pipeWriteCloserFork(out1, out2 chan<- io.WriteCloser, inp <-chan io.WriteCloser) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriteCloserFork(inp <-chan io.WriteCloser) (out1, out2 <-chan io.WriteCloser) {
	cha1 := make(chan io.WriteCloser)
	cha2 := make(chan io.WriteCloser)
	go pipeWriteCloserFork(cha1, cha2, inp)
	return cha1, cha2
}
