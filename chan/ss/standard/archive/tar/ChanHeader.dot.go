// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

// MakeHeaderChan returns a new open channel
// (simply a 'chan *tar.Header' that is).
//
// Note: No 'Header-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myHeaderPipelineStartsHere := MakeHeaderChan()
//	// ... lot's of code to design and build Your favourite "myHeaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myHeaderPipelineStartsHere <- drop
//	}
//	close(myHeaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeHeaderBuffer) the channel is unbuffered.
//
func MakeHeaderChan() (out chan *tar.Header) {
	return make(chan *tar.Header)
}

func sendHeader(out chan<- *tar.Header, inp ...*tar.Header) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanHeader returns a channel to receive all inputs before close.
func ChanHeader(inp ...*tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go sendHeader(cha, inp...)
	return cha
}

func sendHeaderSlice(out chan<- *tar.Header, inp ...[]*tar.Header) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanHeaderSlice returns a channel to receive all inputs before close.
func ChanHeaderSlice(inp ...[]*tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go sendHeaderSlice(cha, inp...)
	return cha
}

func joinHeader(done chan<- struct{}, out chan<- *tar.Header, inp ...*tar.Header) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinHeader
func JoinHeader(out chan<- *tar.Header, inp ...*tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinHeader(cha, out, inp...)
	return cha
}

func joinHeaderSlice(done chan<- struct{}, out chan<- *tar.Header, inp ...[]*tar.Header) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinHeaderSlice
func JoinHeaderSlice(out chan<- *tar.Header, inp ...[]*tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinHeaderSlice(cha, out, inp...)
	return cha
}

func joinHeaderChan(done chan<- struct{}, out chan<- *tar.Header, inp <-chan *tar.Header) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinHeaderChan
func JoinHeaderChan(out chan<- *tar.Header, inp <-chan *tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinHeaderChan(cha, out, inp)
	return cha
}

func doitHeader(done chan<- struct{}, inp <-chan *tar.Header) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneHeader returns a channel to receive one signal before close after inp has been drained.
func DoneHeader(inp <-chan *tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitHeader(cha, inp)
	return cha
}

func doitHeaderSlice(done chan<- ([]*tar.Header), inp <-chan *tar.Header) {
	defer close(done)
	HeaderS := []*tar.Header{}
	for i := range inp {
		HeaderS = append(HeaderS, i)
	}
	done <- HeaderS
}

// DoneHeaderSlice returns a channel which will receive a slice
// of all the Headers received on inp channel before close.
// Unlike DoneHeader, a full slice is sent once, not just an event.
func DoneHeaderSlice(inp <-chan *tar.Header) (done <-chan ([]*tar.Header)) {
	cha := make(chan ([]*tar.Header))
	go doitHeaderSlice(cha, inp)
	return cha
}

func doitHeaderFunc(done chan<- struct{}, inp <-chan *tar.Header, act func(a *tar.Header)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneHeaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *tar.Header) { return }
	}
	go doitHeaderFunc(cha, inp, act)
	return cha
}

func pipeHeaderBuffer(out chan<- *tar.Header, inp <-chan *tar.Header) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeHeaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeHeaderBuffer(inp <-chan *tar.Header, cap int) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header, cap)
	go pipeHeaderBuffer(cha, inp)
	return cha
}

func pipeHeaderFunc(out chan<- *tar.Header, inp <-chan *tar.Header, act func(a *tar.Header) *tar.Header) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeHeaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeHeaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header) *tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	if act == nil {
		act = func(a *tar.Header) *tar.Header { return a }
	}
	go pipeHeaderFunc(cha, inp, act)
	return cha
}

func pipeHeaderFork(out1, out2 chan<- *tar.Header, inp <-chan *tar.Header) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeHeaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeHeaderFork(inp <-chan *tar.Header) (out1, out2 <-chan *tar.Header) {
	cha1 := make(chan *tar.Header)
	cha2 := make(chan *tar.Header)
	go pipeHeaderFork(cha1, cha2, inp)
	return cha1, cha2
}
