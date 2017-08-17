// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriterAtChan returns a new open channel
// (simply a 'chan io.WriterAt' that is).
//
// Note: No 'WriterAt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterAtPipelineStartsHere := MakeWriterAtChan()
//	// ... lot's of code to design and build Your favourite "myWriterAtWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterAtPipelineStartsHere <- drop
//	}
//	close(myWriterAtPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterAtBuffer) the channel is unbuffered.
//
func MakeWriterAtChan() (out chan io.WriterAt) {
	return make(chan io.WriterAt)
}

func sendWriterAt(out chan<- io.WriterAt, inp ...io.WriterAt) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanWriterAt returns a channel to receive all inputs before close.
func ChanWriterAt(inp ...io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	go sendWriterAt(cha, inp...)
	return cha
}

func sendWriterAtSlice(out chan<- io.WriterAt, inp ...[]io.WriterAt) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanWriterAtSlice returns a channel to receive all inputs before close.
func ChanWriterAtSlice(inp ...[]io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	go sendWriterAtSlice(cha, inp...)
	return cha
}

func joinWriterAt(done chan<- struct{}, out chan<- io.WriterAt, inp ...io.WriterAt) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriterAt
func JoinWriterAt(out chan<- io.WriterAt, inp ...io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterAt(cha, out, inp...)
	return cha
}

func joinWriterAtSlice(done chan<- struct{}, out chan<- io.WriterAt, inp ...[]io.WriterAt) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinWriterAtSlice
func JoinWriterAtSlice(out chan<- io.WriterAt, inp ...[]io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterAtSlice(cha, out, inp...)
	return cha
}

func joinWriterAtChan(done chan<- struct{}, out chan<- io.WriterAt, inp <-chan io.WriterAt) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriterAtChan
func JoinWriterAtChan(out chan<- io.WriterAt, inp <-chan io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterAtChan(cha, out, inp)
	return cha
}

func doitWriterAt(done chan<- struct{}, inp <-chan io.WriterAt) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneWriterAt returns a channel to receive one signal before close after inp has been drained.
func DoneWriterAt(inp <-chan io.WriterAt) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitWriterAt(cha, inp)
	return cha
}

func doitWriterAtSlice(done chan<- ([]io.WriterAt), inp <-chan io.WriterAt) {
	defer close(done)
	WriterAtS := []io.WriterAt{}
	for i := range inp {
		WriterAtS = append(WriterAtS, i)
	}
	done <- WriterAtS
}

// DoneWriterAtSlice returns a channel which will receive a slice
// of all the WriterAts received on inp channel before close.
// Unlike DoneWriterAt, a full slice is sent once, not just an event.
func DoneWriterAtSlice(inp <-chan io.WriterAt) (done <-chan ([]io.WriterAt)) {
	cha := make(chan ([]io.WriterAt))
	go doitWriterAtSlice(cha, inp)
	return cha
}

func doitWriterAtFunc(done chan<- struct{}, inp <-chan io.WriterAt, act func(a io.WriterAt)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneWriterAtFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterAtFunc(inp <-chan io.WriterAt, act func(a io.WriterAt)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.WriterAt) { return }
	}
	go doitWriterAtFunc(cha, inp, act)
	return cha
}

func pipeWriterAtBuffer(out chan<- io.WriterAt, inp <-chan io.WriterAt) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeWriterAtBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterAtBuffer(inp <-chan io.WriterAt, cap int) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt, cap)
	go pipeWriterAtBuffer(cha, inp)
	return cha
}

func pipeWriterAtFunc(out chan<- io.WriterAt, inp <-chan io.WriterAt, act func(a io.WriterAt) io.WriterAt) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeWriterAtFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterAtMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterAtFunc(inp <-chan io.WriterAt, act func(a io.WriterAt) io.WriterAt) (out <-chan io.WriterAt) {
	cha := make(chan io.WriterAt)
	if act == nil {
		act = func(a io.WriterAt) io.WriterAt { return a }
	}
	go pipeWriterAtFunc(cha, inp, act)
	return cha
}

func pipeWriterAtFork(out1, out2 chan<- io.WriterAt, inp <-chan io.WriterAt) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeWriterAtFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterAtFork(inp <-chan io.WriterAt) (out1, out2 <-chan io.WriterAt) {
	cha1 := make(chan io.WriterAt)
	cha2 := make(chan io.WriterAt)
	go pipeWriterAtFork(cha1, cha2, inp)
	return cha1, cha2
}
