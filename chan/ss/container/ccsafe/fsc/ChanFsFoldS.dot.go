// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsFoldSChan returns a new open channel
// (simply a 'chan fs.FsFoldS' that is).
//
// Note: No 'FsFoldS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFoldSPipelineStartsHere := MakeFsFoldSChan()
//	// ... lot's of code to design and build Your favourite "myFsFoldSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFoldSPipelineStartsHere <- drop
//	}
//	close(myFsFoldSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFoldSBuffer) the channel is unbuffered.
//
func MakeFsFoldSChan() (out chan fs.FsFoldS) {
	return make(chan fs.FsFoldS)
}

func sendFsFoldS(out chan<- fs.FsFoldS, inp ...fs.FsFoldS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsFoldS returns a channel to receive all inputs before close.
func ChanFsFoldS(inp ...fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	go sendFsFoldS(cha, inp...)
	return cha
}

func sendFsFoldSSlice(out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsFoldSSlice returns a channel to receive all inputs before close.
func ChanFsFoldSSlice(inp ...[]fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	go sendFsFoldSSlice(cha, inp...)
	return cha
}

func joinFsFoldS(done chan<- struct{}, out chan<- fs.FsFoldS, inp ...fs.FsFoldS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFoldS
func JoinFsFoldS(out chan<- fs.FsFoldS, inp ...fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFoldS(cha, out, inp...)
	return cha
}

func joinFsFoldSSlice(done chan<- struct{}, out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsFoldSSlice
func JoinFsFoldSSlice(out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFoldSSlice(cha, out, inp...)
	return cha
}

func joinFsFoldSChan(done chan<- struct{}, out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFoldSChan
func JoinFsFoldSChan(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFoldSChan(cha, out, inp)
	return cha
}

func doitFsFoldS(done chan<- struct{}, inp <-chan fs.FsFoldS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsFoldS returns a channel to receive one signal before close after inp has been drained.
func DoneFsFoldS(inp <-chan fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsFoldS(cha, inp)
	return cha
}

func doitFsFoldSSlice(done chan<- ([]fs.FsFoldS), inp <-chan fs.FsFoldS) {
	defer close(done)
	FsFoldSS := []fs.FsFoldS{}
	for i := range inp {
		FsFoldSS = append(FsFoldSS, i)
	}
	done <- FsFoldSS
}

// DoneFsFoldSSlice returns a channel which will receive a slice
// of all the FsFoldSs received on inp channel before close.
// Unlike DoneFsFoldS, a full slice is sent once, not just an event.
func DoneFsFoldSSlice(inp <-chan fs.FsFoldS) (done <-chan ([]fs.FsFoldS)) {
	cha := make(chan ([]fs.FsFoldS))
	go doitFsFoldSSlice(cha, inp)
	return cha
}

func doitFsFoldSFunc(done chan<- struct{}, inp <-chan fs.FsFoldS, act func(a fs.FsFoldS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsFoldSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFoldSFunc(inp <-chan fs.FsFoldS, act func(a fs.FsFoldS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsFoldS) { return }
	}
	go doitFsFoldSFunc(cha, inp, act)
	return cha
}

func pipeFsFoldSBuffer(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsFoldSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFoldSBuffer(inp <-chan fs.FsFoldS, cap int) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS, cap)
	go pipeFsFoldSBuffer(cha, inp)
	return cha
}

func pipeFsFoldSFunc(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS, act func(a fs.FsFoldS) fs.FsFoldS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsFoldSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFoldSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFoldSFunc(inp <-chan fs.FsFoldS, act func(a fs.FsFoldS) fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	if act == nil {
		act = func(a fs.FsFoldS) fs.FsFoldS { return a }
	}
	go pipeFsFoldSFunc(cha, inp, act)
	return cha
}

func pipeFsFoldSFork(out1, out2 chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsFoldSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFoldSFork(inp <-chan fs.FsFoldS) (out1, out2 <-chan fs.FsFoldS) {
	cha1 := make(chan fs.FsFoldS)
	cha2 := make(chan fs.FsFoldS)
	go pipeFsFoldSFork(cha1, cha2, inp)
	return cha1, cha2
}
