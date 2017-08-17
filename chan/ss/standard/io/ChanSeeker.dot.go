// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeSeekerChan returns a new open channel
// (simply a 'chan io.Seeker' that is).
//
// Note: No 'Seeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySeekerPipelineStartsHere := MakeSeekerChan()
//	// ... lot's of code to design and build Your favourite "mySeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySeekerPipelineStartsHere <- drop
//	}
//	close(mySeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSeekerBuffer) the channel is unbuffered.
//
func MakeSeekerChan() (out chan io.Seeker) {
	return make(chan io.Seeker)
}

func sendSeeker(out chan<- io.Seeker, inp ...io.Seeker) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanSeeker returns a channel to receive all inputs before close.
func ChanSeeker(inp ...io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go sendSeeker(cha, inp...)
	return cha
}

func sendSeekerSlice(out chan<- io.Seeker, inp ...[]io.Seeker) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanSeekerSlice returns a channel to receive all inputs before close.
func ChanSeekerSlice(inp ...[]io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go sendSeekerSlice(cha, inp...)
	return cha
}

func joinSeeker(done chan<- struct{}, out chan<- io.Seeker, inp ...io.Seeker) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSeeker
func JoinSeeker(out chan<- io.Seeker, inp ...io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSeeker(cha, out, inp...)
	return cha
}

func joinSeekerSlice(done chan<- struct{}, out chan<- io.Seeker, inp ...[]io.Seeker) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinSeekerSlice
func JoinSeekerSlice(out chan<- io.Seeker, inp ...[]io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSeekerSlice(cha, out, inp...)
	return cha
}

func joinSeekerChan(done chan<- struct{}, out chan<- io.Seeker, inp <-chan io.Seeker) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSeekerChan
func JoinSeekerChan(out chan<- io.Seeker, inp <-chan io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSeekerChan(cha, out, inp)
	return cha
}

func doitSeeker(done chan<- struct{}, inp <-chan io.Seeker) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneSeeker(inp <-chan io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitSeeker(cha, inp)
	return cha
}

func doitSeekerSlice(done chan<- ([]io.Seeker), inp <-chan io.Seeker) {
	defer close(done)
	SeekerS := []io.Seeker{}
	for i := range inp {
		SeekerS = append(SeekerS, i)
	}
	done <- SeekerS
}

// DoneSeekerSlice returns a channel which will receive a slice
// of all the Seekers received on inp channel before close.
// Unlike DoneSeeker, a full slice is sent once, not just an event.
func DoneSeekerSlice(inp <-chan io.Seeker) (done <-chan ([]io.Seeker)) {
	cha := make(chan ([]io.Seeker))
	go doitSeekerSlice(cha, inp)
	return cha
}

func doitSeekerFunc(done chan<- struct{}, inp <-chan io.Seeker, act func(a io.Seeker)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.Seeker) { return }
	}
	go doitSeekerFunc(cha, inp, act)
	return cha
}

func pipeSeekerBuffer(out chan<- io.Seeker, inp <-chan io.Seeker) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSeekerBuffer(inp <-chan io.Seeker, cap int) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker, cap)
	go pipeSeekerBuffer(cha, inp)
	return cha
}

func pipeSeekerFunc(out chan<- io.Seeker, inp <-chan io.Seeker, act func(a io.Seeker) io.Seeker) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker) io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	if act == nil {
		act = func(a io.Seeker) io.Seeker { return a }
	}
	go pipeSeekerFunc(cha, inp, act)
	return cha
}

func pipeSeekerFork(out1, out2 chan<- io.Seeker, inp <-chan io.Seeker) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSeekerFork(inp <-chan io.Seeker) (out1, out2 <-chan io.Seeker) {
	cha1 := make(chan io.Seeker)
	cha2 := make(chan io.Seeker)
	go pipeSeekerFork(cha1, cha2, inp)
	return cha1, cha2
}
