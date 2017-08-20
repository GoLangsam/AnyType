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
func MakeReadWriteSeekerChan() (out chan io.ReadWriteSeeker) {
	return make(chan io.ReadWriteSeeker)
}

func sendReadWriteSeeker(out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReadWriteSeeker returns a channel to receive all inputs before close.
func ChanReadWriteSeeker(inp ...io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go sendReadWriteSeeker(cha, inp...)
	return cha
}

func sendReadWriteSeekerSlice(out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReadWriteSeekerSlice returns a channel to receive all inputs before close.
func ChanReadWriteSeekerSlice(inp ...[]io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go sendReadWriteSeekerSlice(cha, inp...)
	return cha
}

func joinReadWriteSeeker(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriteSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeeker(out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteSeeker(cha, out, inp...)
	return cha
}

func joinReadWriteSeekerSlice(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReadWriteSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeekerSlice(out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteSeekerSlice(cha, out, inp...)
	return cha
}

func joinReadWriteSeekerChan(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriteSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeekerChan(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteSeekerChan(cha, out, inp)
	return cha
}

func doitReadWriteSeeker(done chan<- struct{}, inp <-chan io.ReadWriteSeeker) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadWriteSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteSeeker(inp <-chan io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadWriteSeeker(cha, inp)
	return cha
}

func doitReadWriteSeekerSlice(done chan<- ([]io.ReadWriteSeeker), inp <-chan io.ReadWriteSeeker) {
	defer close(done)
	ReadWriteSeekerS := []io.ReadWriteSeeker{}
	for i := range inp {
		ReadWriteSeekerS = append(ReadWriteSeekerS, i)
	}
	done <- ReadWriteSeekerS
}

// DoneReadWriteSeekerSlice returns a channel which will receive a slice
// of all the ReadWriteSeekers received on inp channel before close.
// Unlike DoneReadWriteSeeker, a full slice is sent once, not just an event.
func DoneReadWriteSeekerSlice(inp <-chan io.ReadWriteSeeker) (done <-chan ([]io.ReadWriteSeeker)) {
	cha := make(chan ([]io.ReadWriteSeeker))
	go doitReadWriteSeekerSlice(cha, inp)
	return cha
}

func doitReadWriteSeekerFunc(done chan<- struct{}, inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadWriteSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteSeeker) { return }
	}
	go doitReadWriteSeekerFunc(cha, inp, act)
	return cha
}

func pipeReadWriteSeekerBuffer(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadWriteSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteSeekerBuffer(inp <-chan io.ReadWriteSeeker, cap int) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker, cap)
	go pipeReadWriteSeekerBuffer(cha, inp)
	return cha
}

func pipeReadWriteSeekerFunc(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker) io.ReadWriteSeeker) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadWriteSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker) io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	if act == nil {
		act = func(a io.ReadWriteSeeker) io.ReadWriteSeeker { return a }
	}
	go pipeReadWriteSeekerFunc(cha, inp, act)
	return cha
}

func pipeReadWriteSeekerFork(out1, out2 chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadWriteSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteSeekerFork(inp <-chan io.ReadWriteSeeker) (out1, out2 <-chan io.ReadWriteSeeker) {
	cha1 := make(chan io.ReadWriteSeeker)
	cha2 := make(chan io.ReadWriteSeeker)
	go pipeReadWriteSeekerFork(cha1, cha2, inp)
	return cha1, cha2
}
