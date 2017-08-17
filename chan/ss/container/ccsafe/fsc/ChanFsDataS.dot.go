// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsDataSChan returns a new open channel
// (simply a 'chan fs.FsDataS' that is).
//
// Note: No 'FsDataS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsDataSPipelineStartsHere := MakeFsDataSChan()
//	// ... lot's of code to design and build Your favourite "myFsDataSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsDataSPipelineStartsHere <- drop
//	}
//	close(myFsDataSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsDataSBuffer) the channel is unbuffered.
//
func MakeFsDataSChan() (out chan fs.FsDataS) {
	return make(chan fs.FsDataS)
}

func sendFsDataS(out chan<- fs.FsDataS, inp ...fs.FsDataS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsDataS returns a channel to receive all inputs before close.
func ChanFsDataS(inp ...fs.FsDataS) (out <-chan fs.FsDataS) {
	cha := make(chan fs.FsDataS)
	go sendFsDataS(cha, inp...)
	return cha
}

func sendFsDataSSlice(out chan<- fs.FsDataS, inp ...[]fs.FsDataS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsDataSSlice returns a channel to receive all inputs before close.
func ChanFsDataSSlice(inp ...[]fs.FsDataS) (out <-chan fs.FsDataS) {
	cha := make(chan fs.FsDataS)
	go sendFsDataSSlice(cha, inp...)
	return cha
}

func joinFsDataS(done chan<- struct{}, out chan<- fs.FsDataS, inp ...fs.FsDataS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsDataS
func JoinFsDataS(out chan<- fs.FsDataS, inp ...fs.FsDataS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsDataS(cha, out, inp...)
	return cha
}

func joinFsDataSSlice(done chan<- struct{}, out chan<- fs.FsDataS, inp ...[]fs.FsDataS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsDataSSlice
func JoinFsDataSSlice(out chan<- fs.FsDataS, inp ...[]fs.FsDataS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsDataSSlice(cha, out, inp...)
	return cha
}

func joinFsDataSChan(done chan<- struct{}, out chan<- fs.FsDataS, inp <-chan fs.FsDataS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsDataSChan
func JoinFsDataSChan(out chan<- fs.FsDataS, inp <-chan fs.FsDataS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsDataSChan(cha, out, inp)
	return cha
}

func doitFsDataS(done chan<- struct{}, inp <-chan fs.FsDataS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsDataS returns a channel to receive one signal before close after inp has been drained.
func DoneFsDataS(inp <-chan fs.FsDataS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsDataS(cha, inp)
	return cha
}

func doitFsDataSSlice(done chan<- ([]fs.FsDataS), inp <-chan fs.FsDataS) {
	defer close(done)
	FsDataSS := []fs.FsDataS{}
	for i := range inp {
		FsDataSS = append(FsDataSS, i)
	}
	done <- FsDataSS
}

// DoneFsDataSSlice returns a channel which will receive a slice
// of all the FsDataSs received on inp channel before close.
// Unlike DoneFsDataS, a full slice is sent once, not just an event.
func DoneFsDataSSlice(inp <-chan fs.FsDataS) (done <-chan ([]fs.FsDataS)) {
	cha := make(chan ([]fs.FsDataS))
	go doitFsDataSSlice(cha, inp)
	return cha
}

func doitFsDataSFunc(done chan<- struct{}, inp <-chan fs.FsDataS, act func(a fs.FsDataS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsDataSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsDataSFunc(inp <-chan fs.FsDataS, act func(a fs.FsDataS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsDataS) { return }
	}
	go doitFsDataSFunc(cha, inp, act)
	return cha
}

func pipeFsDataSBuffer(out chan<- fs.FsDataS, inp <-chan fs.FsDataS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsDataSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsDataSBuffer(inp <-chan fs.FsDataS, cap int) (out <-chan fs.FsDataS) {
	cha := make(chan fs.FsDataS, cap)
	go pipeFsDataSBuffer(cha, inp)
	return cha
}

func pipeFsDataSFunc(out chan<- fs.FsDataS, inp <-chan fs.FsDataS, act func(a fs.FsDataS) fs.FsDataS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsDataSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsDataSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsDataSFunc(inp <-chan fs.FsDataS, act func(a fs.FsDataS) fs.FsDataS) (out <-chan fs.FsDataS) {
	cha := make(chan fs.FsDataS)
	if act == nil {
		act = func(a fs.FsDataS) fs.FsDataS { return a }
	}
	go pipeFsDataSFunc(cha, inp, act)
	return cha
}

func pipeFsDataSFork(out1, out2 chan<- fs.FsDataS, inp <-chan fs.FsDataS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsDataSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsDataSFork(inp <-chan fs.FsDataS) (out1, out2 <-chan fs.FsDataS) {
	cha1 := make(chan fs.FsDataS)
	cha2 := make(chan fs.FsDataS)
	go pipeFsDataSFork(cha1, cha2, inp)
	return cha1, cha2
}
