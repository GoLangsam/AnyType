// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteWriterChan returns a new open channel
// (simply a 'chan io.ByteWriter' that is).
//
// Note: No 'ByteWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteWriterPipelineStartsHere := MakeByteWriterChan()
//	// ... lot's of code to design and build Your favourite "myByteWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteWriterPipelineStartsHere <- drop
//	}
//	close(myByteWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteWriterBuffer) the channel is unbuffered.
//
func MakeByteWriterChan() (out chan io.ByteWriter) {
	return make(chan io.ByteWriter)
}

func sendByteWriter(out chan<- io.ByteWriter, inp ...io.ByteWriter) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanByteWriter returns a channel to receive all inputs before close.
func ChanByteWriter(inp ...io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go sendByteWriter(cha, inp...)
	return cha
}

func sendByteWriterSlice(out chan<- io.ByteWriter, inp ...[]io.ByteWriter) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanByteWriterSlice returns a channel to receive all inputs before close.
func ChanByteWriterSlice(inp ...[]io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go sendByteWriterSlice(cha, inp...)
	return cha
}

func joinByteWriter(done chan<- struct{}, out chan<- io.ByteWriter, inp ...io.ByteWriter) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinByteWriter
func JoinByteWriter(out chan<- io.ByteWriter, inp ...io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteWriter(cha, out, inp...)
	return cha
}

func joinByteWriterSlice(done chan<- struct{}, out chan<- io.ByteWriter, inp ...[]io.ByteWriter) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinByteWriterSlice
func JoinByteWriterSlice(out chan<- io.ByteWriter, inp ...[]io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteWriterSlice(cha, out, inp...)
	return cha
}

func joinByteWriterChan(done chan<- struct{}, out chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinByteWriterChan
func JoinByteWriterChan(out chan<- io.ByteWriter, inp <-chan io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteWriterChan(cha, out, inp)
	return cha
}

func doitByteWriter(done chan<- struct{}, inp <-chan io.ByteWriter) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneByteWriter returns a channel to receive one signal before close after inp has been drained.
func DoneByteWriter(inp <-chan io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitByteWriter(cha, inp)
	return cha
}

func doitByteWriterSlice(done chan<- ([]io.ByteWriter), inp <-chan io.ByteWriter) {
	defer close(done)
	ByteWriterS := []io.ByteWriter{}
	for i := range inp {
		ByteWriterS = append(ByteWriterS, i)
	}
	done <- ByteWriterS
}

// DoneByteWriterSlice returns a channel which will receive a slice
// of all the ByteWriters received on inp channel before close.
// Unlike DoneByteWriter, a full slice is sent once, not just an event.
func DoneByteWriterSlice(inp <-chan io.ByteWriter) (done <-chan ([]io.ByteWriter)) {
	cha := make(chan ([]io.ByteWriter))
	go doitByteWriterSlice(cha, inp)
	return cha
}

func doitByteWriterFunc(done chan<- struct{}, inp <-chan io.ByteWriter, act func(a io.ByteWriter)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneByteWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ByteWriter) { return }
	}
	go doitByteWriterFunc(cha, inp, act)
	return cha
}

func pipeByteWriterBuffer(out chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeByteWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteWriterBuffer(inp <-chan io.ByteWriter, cap int) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter, cap)
	go pipeByteWriterBuffer(cha, inp)
	return cha
}

func pipeByteWriterFunc(out chan<- io.ByteWriter, inp <-chan io.ByteWriter, act func(a io.ByteWriter) io.ByteWriter) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeByteWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter) io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	if act == nil {
		act = func(a io.ByteWriter) io.ByteWriter { return a }
	}
	go pipeByteWriterFunc(cha, inp, act)
	return cha
}

func pipeByteWriterFork(out1, out2 chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeByteWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteWriterFork(inp <-chan io.ByteWriter) (out1, out2 <-chan io.ByteWriter) {
	cha1 := make(chan io.ByteWriter)
	cha2 := make(chan io.ByteWriter)
	go pipeByteWriterFork(cha1, cha2, inp)
	return cha1, cha2
}
