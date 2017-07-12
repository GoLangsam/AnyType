// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakePipeReaderChan returns a new open channel
// (simply a 'chan *io.PipeReader' that is).
//
// Note: No 'PipeReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipeReaderPipelineStartsHere := MakePipeReaderChan()
//	// ... lot's of code to design and build Your favourite "myPipeReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipeReaderPipelineStartsHere <- drop
//	}
//	close(myPipeReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePipeReaderBuffer) the channel is unbuffered.
//
func MakePipeReaderChan() (out chan *io.PipeReader) {
	return make(chan *io.PipeReader)
}

func sendPipeReader(out chan<- *io.PipeReader, inp ...*io.PipeReader) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanPipeReader returns a channel to receive all inputs before close.
func ChanPipeReader(inp ...*io.PipeReader) (out <-chan *io.PipeReader) {
	cha := make(chan *io.PipeReader)
	go sendPipeReader(cha, inp...)
	return cha
}

func sendPipeReaderSlice(out chan<- *io.PipeReader, inp ...[]*io.PipeReader) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanPipeReaderSlice returns a channel to receive all inputs before close.
func ChanPipeReaderSlice(inp ...[]*io.PipeReader) (out <-chan *io.PipeReader) {
	cha := make(chan *io.PipeReader)
	go sendPipeReaderSlice(cha, inp...)
	return cha
}

func joinPipeReader(done chan<- struct{}, out chan<- *io.PipeReader, inp ...*io.PipeReader) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPipeReader
func JoinPipeReader(out chan<- *io.PipeReader, inp ...*io.PipeReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeReader(cha, out, inp...)
	return cha
}

func joinPipeReaderSlice(done chan<- struct{}, out chan<- *io.PipeReader, inp ...[]*io.PipeReader) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinPipeReaderSlice
func JoinPipeReaderSlice(out chan<- *io.PipeReader, inp ...[]*io.PipeReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeReaderSlice(cha, out, inp...)
	return cha
}

func joinPipeReaderChan(done chan<- struct{}, out chan<- *io.PipeReader, inp <-chan *io.PipeReader) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPipeReaderChan
func JoinPipeReaderChan(out chan<- *io.PipeReader, inp <-chan *io.PipeReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeReaderChan(cha, out, inp)
	return cha
}

func doitPipeReader(done chan<- struct{}, inp <-chan *io.PipeReader) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePipeReader returns a channel to receive one signal before close after inp has been drained.
func DonePipeReader(inp <-chan *io.PipeReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPipeReader(cha, inp)
	return cha
}

func doitPipeReaderSlice(done chan<- ([]*io.PipeReader), inp <-chan *io.PipeReader) {
	defer close(done)
	PipeReaderS := []*io.PipeReader{}
	for i := range inp {
		PipeReaderS = append(PipeReaderS, i)
	}
	done <- PipeReaderS
}

// DonePipeReaderSlice returns a channel which will receive a slice
// of all the PipeReaders received on inp channel before close.
// Unlike DonePipeReader, a full slice is sent once, not just an event.
func DonePipeReaderSlice(inp <-chan *io.PipeReader) (done <-chan ([]*io.PipeReader)) {
	cha := make(chan ([]*io.PipeReader))
	go doitPipeReaderSlice(cha, inp)
	return cha
}

func doitPipeReaderFunc(done chan<- struct{}, inp <-chan *io.PipeReader, act func(a *io.PipeReader)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePipeReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePipeReaderFunc(inp <-chan *io.PipeReader, act func(a *io.PipeReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.PipeReader) { return }
	}
	go doitPipeReaderFunc(cha, inp, act)
	return cha
}

func pipePipeReaderBuffer(out chan<- *io.PipeReader, inp <-chan *io.PipeReader) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePipeReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePipeReaderBuffer(inp <-chan *io.PipeReader, cap int) (out <-chan *io.PipeReader) {
	cha := make(chan *io.PipeReader, cap)
	go pipePipeReaderBuffer(cha, inp)
	return cha
}

func pipePipeReaderFunc(out chan<- *io.PipeReader, inp <-chan *io.PipeReader, act func(a *io.PipeReader) *io.PipeReader) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePipeReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePipeReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePipeReaderFunc(inp <-chan *io.PipeReader, act func(a *io.PipeReader) *io.PipeReader) (out <-chan *io.PipeReader) {
	cha := make(chan *io.PipeReader)
	if act == nil {
		act = func(a *io.PipeReader) *io.PipeReader { return a }
	}
	go pipePipeReaderFunc(cha, inp, act)
	return cha
}

func pipePipeReaderFork(out1, out2 chan<- *io.PipeReader, inp <-chan *io.PipeReader) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePipeReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePipeReaderFork(inp <-chan *io.PipeReader) (out1, out2 <-chan *io.PipeReader) {
	cha1 := make(chan *io.PipeReader)
	cha2 := make(chan *io.PipeReader)
	go pipePipeReaderFork(cha1, cha2, inp)
	return cha1, cha2
}
