// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// MakeReadWriterChan returns a new open channel
// (simply a 'chan *bufio.ReadWriter' that is).
//
// Note: No 'ReadWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriterPipelineStartsHere := MakeReadWriterChan()
//	// ... lot's of code to design and build Your favourite "myReadWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriterPipelineStartsHere <- drop
//	}
//	close(myReadWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriterBuffer) the channel is unbuffered.
//
func MakeReadWriterChan() (out chan *bufio.ReadWriter) {
	return make(chan *bufio.ReadWriter)
}

func sendReadWriter(out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReadWriter returns a channel to receive all inputs before close.
func ChanReadWriter(inp ...*bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go sendReadWriter(cha, inp...)
	return cha
}

func sendReadWriterSlice(out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReadWriterSlice returns a channel to receive all inputs before close.
func ChanReadWriterSlice(inp ...[]*bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go sendReadWriterSlice(cha, inp...)
	return cha
}

func joinReadWriter(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriter
func JoinReadWriter(out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriter(cha, out, inp...)
	return cha
}

func joinReadWriterSlice(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReadWriterSlice
func JoinReadWriterSlice(out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriterSlice(cha, out, inp...)
	return cha
}

func joinReadWriterChan(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriterChan
func JoinReadWriterChan(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriterChan(cha, out, inp)
	return cha
}

func doitReadWriter(done chan<- struct{}, inp <-chan *bufio.ReadWriter) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadWriter returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriter(inp <-chan *bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadWriter(cha, inp)
	return cha
}

func doitReadWriterSlice(done chan<- ([]*bufio.ReadWriter), inp <-chan *bufio.ReadWriter) {
	defer close(done)
	ReadWriterS := []*bufio.ReadWriter{}
	for i := range inp {
		ReadWriterS = append(ReadWriterS, i)
	}
	done <- ReadWriterS
}

// DoneReadWriterSlice returns a channel which will receive a slice
// of all the ReadWriters received on inp channel before close.
// Unlike DoneReadWriter, a full slice is sent once, not just an event.
func DoneReadWriterSlice(inp <-chan *bufio.ReadWriter) (done <-chan ([]*bufio.ReadWriter)) {
	cha := make(chan ([]*bufio.ReadWriter))
	go doitReadWriterSlice(cha, inp)
	return cha
}

func doitReadWriterFunc(done chan<- struct{}, inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriterFunc(inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.ReadWriter) { return }
	}
	go doitReadWriterFunc(cha, inp, act)
	return cha
}

func pipeReadWriterBuffer(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriterBuffer(inp <-chan *bufio.ReadWriter, cap int) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter, cap)
	go pipeReadWriterBuffer(cha, inp)
	return cha
}

func pipeReadWriterFunc(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter) *bufio.ReadWriter) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriterFunc(inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter) *bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	if act == nil {
		act = func(a *bufio.ReadWriter) *bufio.ReadWriter { return a }
	}
	go pipeReadWriterFunc(cha, inp, act)
	return cha
}

func pipeReadWriterFork(out1, out2 chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriterFork(inp <-chan *bufio.ReadWriter) (out1, out2 <-chan *bufio.ReadWriter) {
	cha1 := make(chan *bufio.ReadWriter)
	cha2 := make(chan *bufio.ReadWriter)
	go pipeReadWriterFork(cha1, cha2, inp)
	return cha1, cha2
}
