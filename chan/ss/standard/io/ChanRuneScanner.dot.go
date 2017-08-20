// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeRuneScannerChan returns a new open channel
// (simply a 'chan io.RuneScanner' that is).
//
// Note: No 'RuneScanner-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myRuneScannerPipelineStartsHere := MakeRuneScannerChan()
//	// ... lot's of code to design and build Your favourite "myRuneScannerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myRuneScannerPipelineStartsHere <- drop
//	}
//	close(myRuneScannerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeRuneScannerBuffer) the channel is unbuffered.
//
func MakeRuneScannerChan() (out chan io.RuneScanner) {
	return make(chan io.RuneScanner)
}

func sendRuneScanner(out chan<- io.RuneScanner, inp ...io.RuneScanner) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanRuneScanner returns a channel to receive all inputs before close.
func ChanRuneScanner(inp ...io.RuneScanner) (out <-chan io.RuneScanner) {
	cha := make(chan io.RuneScanner)
	go sendRuneScanner(cha, inp...)
	return cha
}

func sendRuneScannerSlice(out chan<- io.RuneScanner, inp ...[]io.RuneScanner) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanRuneScannerSlice returns a channel to receive all inputs before close.
func ChanRuneScannerSlice(inp ...[]io.RuneScanner) (out <-chan io.RuneScanner) {
	cha := make(chan io.RuneScanner)
	go sendRuneScannerSlice(cha, inp...)
	return cha
}

func joinRuneScanner(done chan<- struct{}, out chan<- io.RuneScanner, inp ...io.RuneScanner) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinRuneScanner sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScanner(out chan<- io.RuneScanner, inp ...io.RuneScanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinRuneScanner(cha, out, inp...)
	return cha
}

func joinRuneScannerSlice(done chan<- struct{}, out chan<- io.RuneScanner, inp ...[]io.RuneScanner) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinRuneScannerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScannerSlice(out chan<- io.RuneScanner, inp ...[]io.RuneScanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinRuneScannerSlice(cha, out, inp...)
	return cha
}

func joinRuneScannerChan(done chan<- struct{}, out chan<- io.RuneScanner, inp <-chan io.RuneScanner) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinRuneScannerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScannerChan(out chan<- io.RuneScanner, inp <-chan io.RuneScanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinRuneScannerChan(cha, out, inp)
	return cha
}

func doitRuneScanner(done chan<- struct{}, inp <-chan io.RuneScanner) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneRuneScanner returns a channel to receive one signal before close after inp has been drained.
func DoneRuneScanner(inp <-chan io.RuneScanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitRuneScanner(cha, inp)
	return cha
}

func doitRuneScannerSlice(done chan<- ([]io.RuneScanner), inp <-chan io.RuneScanner) {
	defer close(done)
	RuneScannerS := []io.RuneScanner{}
	for i := range inp {
		RuneScannerS = append(RuneScannerS, i)
	}
	done <- RuneScannerS
}

// DoneRuneScannerSlice returns a channel which will receive a slice
// of all the RuneScanners received on inp channel before close.
// Unlike DoneRuneScanner, a full slice is sent once, not just an event.
func DoneRuneScannerSlice(inp <-chan io.RuneScanner) (done <-chan ([]io.RuneScanner)) {
	cha := make(chan ([]io.RuneScanner))
	go doitRuneScannerSlice(cha, inp)
	return cha
}

func doitRuneScannerFunc(done chan<- struct{}, inp <-chan io.RuneScanner, act func(a io.RuneScanner)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneRuneScannerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneRuneScannerFunc(inp <-chan io.RuneScanner, act func(a io.RuneScanner)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.RuneScanner) { return }
	}
	go doitRuneScannerFunc(cha, inp, act)
	return cha
}

func pipeRuneScannerBuffer(out chan<- io.RuneScanner, inp <-chan io.RuneScanner) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeRuneScannerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeRuneScannerBuffer(inp <-chan io.RuneScanner, cap int) (out <-chan io.RuneScanner) {
	cha := make(chan io.RuneScanner, cap)
	go pipeRuneScannerBuffer(cha, inp)
	return cha
}

func pipeRuneScannerFunc(out chan<- io.RuneScanner, inp <-chan io.RuneScanner, act func(a io.RuneScanner) io.RuneScanner) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeRuneScannerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeRuneScannerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeRuneScannerFunc(inp <-chan io.RuneScanner, act func(a io.RuneScanner) io.RuneScanner) (out <-chan io.RuneScanner) {
	cha := make(chan io.RuneScanner)
	if act == nil {
		act = func(a io.RuneScanner) io.RuneScanner { return a }
	}
	go pipeRuneScannerFunc(cha, inp, act)
	return cha
}

func pipeRuneScannerFork(out1, out2 chan<- io.RuneScanner, inp <-chan io.RuneScanner) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeRuneScannerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeRuneScannerFork(inp <-chan io.RuneScanner) (out1, out2 <-chan io.RuneScanner) {
	cha1 := make(chan io.RuneScanner)
	cha2 := make(chan io.RuneScanner)
	go pipeRuneScannerFork(cha1, cha2, inp)
	return cha1, cha2
}
