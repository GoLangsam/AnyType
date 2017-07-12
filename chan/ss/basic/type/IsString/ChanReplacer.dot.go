// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// MakeReplacerChan returns a new open channel
// (simply a 'chan *strings.Replacer' that is).
//
// Note: No 'Replacer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReplacerPipelineStartsHere := MakeReplacerChan()
//	// ... lot's of code to design and build Your favourite "myReplacerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReplacerPipelineStartsHere <- drop
//	}
//	close(myReplacerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReplacerBuffer) the channel is unbuffered.
//
func MakeReplacerChan() (out chan *strings.Replacer) {
	return make(chan *strings.Replacer)
}

func sendReplacer(out chan<- *strings.Replacer, inp ...*strings.Replacer) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReplacer returns a channel to receive all inputs before close.
func ChanReplacer(inp ...*strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go sendReplacer(cha, inp...)
	return cha
}

func sendReplacerSlice(out chan<- *strings.Replacer, inp ...[]*strings.Replacer) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReplacerSlice returns a channel to receive all inputs before close.
func ChanReplacerSlice(inp ...[]*strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	go sendReplacerSlice(cha, inp...)
	return cha
}

func joinReplacer(done chan<- struct{}, out chan<- *strings.Replacer, inp ...*strings.Replacer) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReplacer
func JoinReplacer(out chan<- *strings.Replacer, inp ...*strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReplacer(cha, out, inp...)
	return cha
}

func joinReplacerSlice(done chan<- struct{}, out chan<- *strings.Replacer, inp ...[]*strings.Replacer) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReplacerSlice
func JoinReplacerSlice(out chan<- *strings.Replacer, inp ...[]*strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReplacerSlice(cha, out, inp...)
	return cha
}

func joinReplacerChan(done chan<- struct{}, out chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReplacerChan
func JoinReplacerChan(out chan<- *strings.Replacer, inp <-chan *strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReplacerChan(cha, out, inp)
	return cha
}

func doitReplacer(done chan<- struct{}, inp <-chan *strings.Replacer) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReplacer returns a channel to receive one signal before close after inp has been drained.
func DoneReplacer(inp <-chan *strings.Replacer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReplacer(cha, inp)
	return cha
}

func doitReplacerSlice(done chan<- ([]*strings.Replacer), inp <-chan *strings.Replacer) {
	defer close(done)
	ReplacerS := []*strings.Replacer{}
	for i := range inp {
		ReplacerS = append(ReplacerS, i)
	}
	done <- ReplacerS
}

// DoneReplacerSlice returns a channel which will receive a slice
// of all the Replacers received on inp channel before close.
// Unlike DoneReplacer, a full slice is sent once, not just an event.
func DoneReplacerSlice(inp <-chan *strings.Replacer) (done <-chan ([]*strings.Replacer)) {
	cha := make(chan ([]*strings.Replacer))
	go doitReplacerSlice(cha, inp)
	return cha
}

func doitReplacerFunc(done chan<- struct{}, inp <-chan *strings.Replacer, act func(a *strings.Replacer)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReplacerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *strings.Replacer) { return }
	}
	go doitReplacerFunc(cha, inp, act)
	return cha
}

func pipeReplacerBuffer(out chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReplacerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReplacerBuffer(inp <-chan *strings.Replacer, cap int) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer, cap)
	go pipeReplacerBuffer(cha, inp)
	return cha
}

func pipeReplacerFunc(out chan<- *strings.Replacer, inp <-chan *strings.Replacer, act func(a *strings.Replacer) *strings.Replacer) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReplacerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReplacerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer) *strings.Replacer) (out <-chan *strings.Replacer) {
	cha := make(chan *strings.Replacer)
	if act == nil {
		act = func(a *strings.Replacer) *strings.Replacer { return a }
	}
	go pipeReplacerFunc(cha, inp, act)
	return cha
}

func pipeReplacerFork(out1, out2 chan<- *strings.Replacer, inp <-chan *strings.Replacer) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReplacerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReplacerFork(inp <-chan *strings.Replacer) (out1, out2 <-chan *strings.Replacer) {
	cha1 := make(chan *strings.Replacer)
	cha2 := make(chan *strings.Replacer)
	go pipeReplacerFork(cha1, cha2, inp)
	return cha1, cha2
}
