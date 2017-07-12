// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriterToChan returns a new open channel
// (simply a 'chan io.WriterTo' that is).
//
// Note: No 'WriterTo-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterToPipelineStartsHere := MakeWriterToChan()
//	// ... lot's of code to design and build Your favourite "myWriterToWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterToPipelineStartsHere <- drop
//	}
//	close(myWriterToPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterToBuffer) the channel is unbuffered.
//
func MakeWriterToChan() (out chan io.WriterTo) {
	return make(chan io.WriterTo)
}

func sendWriterTo(out chan<- io.WriterTo, inp ...io.WriterTo) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanWriterTo returns a channel to receive all inputs before close.
func ChanWriterTo(inp ...io.WriterTo) (out <-chan io.WriterTo) {
	cha := make(chan io.WriterTo)
	go sendWriterTo(cha, inp...)
	return cha
}

func sendWriterToSlice(out chan<- io.WriterTo, inp ...[]io.WriterTo) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanWriterToSlice returns a channel to receive all inputs before close.
func ChanWriterToSlice(inp ...[]io.WriterTo) (out <-chan io.WriterTo) {
	cha := make(chan io.WriterTo)
	go sendWriterToSlice(cha, inp...)
	return cha
}

func joinWriterTo(done chan<- struct{}, out chan<- io.WriterTo, inp ...io.WriterTo) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriterTo
func JoinWriterTo(out chan<- io.WriterTo, inp ...io.WriterTo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterTo(cha, out, inp...)
	return cha
}

func joinWriterToSlice(done chan<- struct{}, out chan<- io.WriterTo, inp ...[]io.WriterTo) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinWriterToSlice
func JoinWriterToSlice(out chan<- io.WriterTo, inp ...[]io.WriterTo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterToSlice(cha, out, inp...)
	return cha
}

func joinWriterToChan(done chan<- struct{}, out chan<- io.WriterTo, inp <-chan io.WriterTo) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriterToChan
func JoinWriterToChan(out chan<- io.WriterTo, inp <-chan io.WriterTo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterToChan(cha, out, inp)
	return cha
}

func doitWriterTo(done chan<- struct{}, inp <-chan io.WriterTo) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneWriterTo returns a channel to receive one signal before close after inp has been drained.
func DoneWriterTo(inp <-chan io.WriterTo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitWriterTo(cha, inp)
	return cha
}

func doitWriterToSlice(done chan<- ([]io.WriterTo), inp <-chan io.WriterTo) {
	defer close(done)
	WriterToS := []io.WriterTo{}
	for i := range inp {
		WriterToS = append(WriterToS, i)
	}
	done <- WriterToS
}

// DoneWriterToSlice returns a channel which will receive a slice
// of all the WriterTos received on inp channel before close.
// Unlike DoneWriterTo, a full slice is sent once, not just an event.
func DoneWriterToSlice(inp <-chan io.WriterTo) (done <-chan ([]io.WriterTo)) {
	cha := make(chan ([]io.WriterTo))
	go doitWriterToSlice(cha, inp)
	return cha
}

func doitWriterToFunc(done chan<- struct{}, inp <-chan io.WriterTo, act func(a io.WriterTo)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneWriterToFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterToFunc(inp <-chan io.WriterTo, act func(a io.WriterTo)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.WriterTo) { return }
	}
	go doitWriterToFunc(cha, inp, act)
	return cha
}

func pipeWriterToBuffer(out chan<- io.WriterTo, inp <-chan io.WriterTo) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeWriterToBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterToBuffer(inp <-chan io.WriterTo, cap int) (out <-chan io.WriterTo) {
	cha := make(chan io.WriterTo, cap)
	go pipeWriterToBuffer(cha, inp)
	return cha
}

func pipeWriterToFunc(out chan<- io.WriterTo, inp <-chan io.WriterTo, act func(a io.WriterTo) io.WriterTo) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeWriterToFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterToMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterToFunc(inp <-chan io.WriterTo, act func(a io.WriterTo) io.WriterTo) (out <-chan io.WriterTo) {
	cha := make(chan io.WriterTo)
	if act == nil {
		act = func(a io.WriterTo) io.WriterTo { return a }
	}
	go pipeWriterToFunc(cha, inp, act)
	return cha
}

func pipeWriterToFork(out1, out2 chan<- io.WriterTo, inp <-chan io.WriterTo) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeWriterToFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterToFork(inp <-chan io.WriterTo) (out1, out2 <-chan io.WriterTo) {
	cha1 := make(chan io.WriterTo)
	cha2 := make(chan io.WriterTo)
	go pipeWriterToFork(cha1, cha2, inp)
	return cha1, cha2
}