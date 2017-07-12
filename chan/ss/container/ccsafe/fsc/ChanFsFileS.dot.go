// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsFileSChan returns a new open channel
// (simply a 'chan fs.FsFileS' that is).
//
// Note: No 'FsFileS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFileSPipelineStartsHere := MakeFsFileSChan()
//	// ... lot's of code to design and build Your favourite "myFsFileSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFileSPipelineStartsHere <- drop
//	}
//	close(myFsFileSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFileSBuffer) the channel is unbuffered.
//
func MakeFsFileSChan() (out chan fs.FsFileS) {
	return make(chan fs.FsFileS)
}

func sendFsFileS(out chan<- fs.FsFileS, inp ...fs.FsFileS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsFileS returns a channel to receive all inputs before close.
func ChanFsFileS(inp ...fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go sendFsFileS(cha, inp...)
	return cha
}

func sendFsFileSSlice(out chan<- fs.FsFileS, inp ...[]fs.FsFileS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsFileSSlice returns a channel to receive all inputs before close.
func ChanFsFileSSlice(inp ...[]fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go sendFsFileSSlice(cha, inp...)
	return cha
}

func joinFsFileS(done chan<- struct{}, out chan<- fs.FsFileS, inp ...fs.FsFileS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFileS
func JoinFsFileS(out chan<- fs.FsFileS, inp ...fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFileS(cha, out, inp...)
	return cha
}

func joinFsFileSSlice(done chan<- struct{}, out chan<- fs.FsFileS, inp ...[]fs.FsFileS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsFileSSlice
func JoinFsFileSSlice(out chan<- fs.FsFileS, inp ...[]fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFileSSlice(cha, out, inp...)
	return cha
}

func joinFsFileSChan(done chan<- struct{}, out chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFileSChan
func JoinFsFileSChan(out chan<- fs.FsFileS, inp <-chan fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFileSChan(cha, out, inp)
	return cha
}

func doitFsFileS(done chan<- struct{}, inp <-chan fs.FsFileS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsFileS returns a channel to receive one signal before close after inp has been drained.
func DoneFsFileS(inp <-chan fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsFileS(cha, inp)
	return cha
}

func doitFsFileSSlice(done chan<- ([]fs.FsFileS), inp <-chan fs.FsFileS) {
	defer close(done)
	FsFileSS := []fs.FsFileS{}
	for i := range inp {
		FsFileSS = append(FsFileSS, i)
	}
	done <- FsFileSS
}

// DoneFsFileSSlice returns a channel which will receive a slice
// of all the FsFileSs received on inp channel before close.
// Unlike DoneFsFileS, a full slice is sent once, not just an event.
func DoneFsFileSSlice(inp <-chan fs.FsFileS) (done <-chan ([]fs.FsFileS)) {
	cha := make(chan ([]fs.FsFileS))
	go doitFsFileSSlice(cha, inp)
	return cha
}

func doitFsFileSFunc(done chan<- struct{}, inp <-chan fs.FsFileS, act func(a fs.FsFileS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsFileSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsFileS) { return }
	}
	go doitFsFileSFunc(cha, inp, act)
	return cha
}

func pipeFsFileSBuffer(out chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsFileSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFileSBuffer(inp <-chan fs.FsFileS, cap int) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS, cap)
	go pipeFsFileSBuffer(cha, inp)
	return cha
}

func pipeFsFileSFunc(out chan<- fs.FsFileS, inp <-chan fs.FsFileS, act func(a fs.FsFileS) fs.FsFileS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsFileSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFileSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS) fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	if act == nil {
		act = func(a fs.FsFileS) fs.FsFileS { return a }
	}
	go pipeFsFileSFunc(cha, inp, act)
	return cha
}

func pipeFsFileSFork(out1, out2 chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsFileSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFileSFork(inp <-chan fs.FsFileS) (out1, out2 <-chan fs.FsFileS) {
	cha1 := make(chan fs.FsFileS)
	cha2 := make(chan fs.FsFileS)
	go pipeFsFileSFork(cha1, cha2, inp)
	return cha1, cha2
}
