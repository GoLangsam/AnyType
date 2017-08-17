// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakePatternSChan returns a new open channel
// (simply a 'chan fs.PatternS' that is).
//
// Note: No 'PatternS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatternSPipelineStartsHere := MakePatternSChan()
//	// ... lot's of code to design and build Your favourite "myPatternSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatternSPipelineStartsHere <- drop
//	}
//	close(myPatternSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatternSBuffer) the channel is unbuffered.
//
func MakePatternSChan() (out chan fs.PatternS) {
	return make(chan fs.PatternS)
}

func sendPatternS(out chan<- fs.PatternS, inp ...fs.PatternS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanPatternS returns a channel to receive all inputs before close.
func ChanPatternS(inp ...fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go sendPatternS(cha, inp...)
	return cha
}

func sendPatternSSlice(out chan<- fs.PatternS, inp ...[]fs.PatternS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanPatternSSlice returns a channel to receive all inputs before close.
func ChanPatternSSlice(inp ...[]fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go sendPatternSSlice(cha, inp...)
	return cha
}

func joinPatternS(done chan<- struct{}, out chan<- fs.PatternS, inp ...fs.PatternS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPatternS
func JoinPatternS(out chan<- fs.PatternS, inp ...fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPatternS(cha, out, inp...)
	return cha
}

func joinPatternSSlice(done chan<- struct{}, out chan<- fs.PatternS, inp ...[]fs.PatternS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinPatternSSlice
func JoinPatternSSlice(out chan<- fs.PatternS, inp ...[]fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPatternSSlice(cha, out, inp...)
	return cha
}

func joinPatternSChan(done chan<- struct{}, out chan<- fs.PatternS, inp <-chan fs.PatternS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPatternSChan
func JoinPatternSChan(out chan<- fs.PatternS, inp <-chan fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPatternSChan(cha, out, inp)
	return cha
}

func doitPatternS(done chan<- struct{}, inp <-chan fs.PatternS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePatternS returns a channel to receive one signal before close after inp has been drained.
func DonePatternS(inp <-chan fs.PatternS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPatternS(cha, inp)
	return cha
}

func doitPatternSSlice(done chan<- ([]fs.PatternS), inp <-chan fs.PatternS) {
	defer close(done)
	PatternSS := []fs.PatternS{}
	for i := range inp {
		PatternSS = append(PatternSS, i)
	}
	done <- PatternSS
}

// DonePatternSSlice returns a channel which will receive a slice
// of all the PatternSs received on inp channel before close.
// Unlike DonePatternS, a full slice is sent once, not just an event.
func DonePatternSSlice(inp <-chan fs.PatternS) (done <-chan ([]fs.PatternS)) {
	cha := make(chan ([]fs.PatternS))
	go doitPatternSSlice(cha, inp)
	return cha
}

func doitPatternSFunc(done chan<- struct{}, inp <-chan fs.PatternS, act func(a fs.PatternS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePatternSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.PatternS) { return }
	}
	go doitPatternSFunc(cha, inp, act)
	return cha
}

func pipePatternSBuffer(out chan<- fs.PatternS, inp <-chan fs.PatternS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePatternSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternSBuffer(inp <-chan fs.PatternS, cap int) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS, cap)
	go pipePatternSBuffer(cha, inp)
	return cha
}

func pipePatternSFunc(out chan<- fs.PatternS, inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePatternSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	if act == nil {
		act = func(a fs.PatternS) fs.PatternS { return a }
	}
	go pipePatternSFunc(cha, inp, act)
	return cha
}

func pipePatternSFork(out1, out2 chan<- fs.PatternS, inp <-chan fs.PatternS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePatternSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternSFork(inp <-chan fs.PatternS) (out1, out2 <-chan fs.PatternS) {
	cha1 := make(chan fs.PatternS)
	cha2 := make(chan fs.PatternS)
	go pipePatternSFork(cha1, cha2, inp)
	return cha1, cha2
}
