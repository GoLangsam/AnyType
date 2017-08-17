// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsPathSChan returns a new open channel
// (simply a 'chan fs.FsPathS' that is).
//
// Note: No 'FsPathS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsPathSPipelineStartsHere := MakeFsPathSChan()
//	// ... lot's of code to design and build Your favourite "myFsPathSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsPathSPipelineStartsHere <- drop
//	}
//	close(myFsPathSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsPathSBuffer) the channel is unbuffered.
//
func MakeFsPathSChan() (out chan fs.FsPathS) {
	return make(chan fs.FsPathS)
}

func sendFsPathS(out chan<- fs.FsPathS, inp ...fs.FsPathS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsPathS returns a channel to receive all inputs before close.
func ChanFsPathS(inp ...fs.FsPathS) (out <-chan fs.FsPathS) {
	cha := make(chan fs.FsPathS)
	go sendFsPathS(cha, inp...)
	return cha
}

func sendFsPathSSlice(out chan<- fs.FsPathS, inp ...[]fs.FsPathS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsPathSSlice returns a channel to receive all inputs before close.
func ChanFsPathSSlice(inp ...[]fs.FsPathS) (out <-chan fs.FsPathS) {
	cha := make(chan fs.FsPathS)
	go sendFsPathSSlice(cha, inp...)
	return cha
}

func joinFsPathS(done chan<- struct{}, out chan<- fs.FsPathS, inp ...fs.FsPathS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsPathS
func JoinFsPathS(out chan<- fs.FsPathS, inp ...fs.FsPathS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsPathS(cha, out, inp...)
	return cha
}

func joinFsPathSSlice(done chan<- struct{}, out chan<- fs.FsPathS, inp ...[]fs.FsPathS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsPathSSlice
func JoinFsPathSSlice(out chan<- fs.FsPathS, inp ...[]fs.FsPathS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsPathSSlice(cha, out, inp...)
	return cha
}

func joinFsPathSChan(done chan<- struct{}, out chan<- fs.FsPathS, inp <-chan fs.FsPathS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsPathSChan
func JoinFsPathSChan(out chan<- fs.FsPathS, inp <-chan fs.FsPathS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsPathSChan(cha, out, inp)
	return cha
}

func doitFsPathS(done chan<- struct{}, inp <-chan fs.FsPathS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsPathS returns a channel to receive one signal before close after inp has been drained.
func DoneFsPathS(inp <-chan fs.FsPathS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsPathS(cha, inp)
	return cha
}

func doitFsPathSSlice(done chan<- ([]fs.FsPathS), inp <-chan fs.FsPathS) {
	defer close(done)
	FsPathSS := []fs.FsPathS{}
	for i := range inp {
		FsPathSS = append(FsPathSS, i)
	}
	done <- FsPathSS
}

// DoneFsPathSSlice returns a channel which will receive a slice
// of all the FsPathSs received on inp channel before close.
// Unlike DoneFsPathS, a full slice is sent once, not just an event.
func DoneFsPathSSlice(inp <-chan fs.FsPathS) (done <-chan ([]fs.FsPathS)) {
	cha := make(chan ([]fs.FsPathS))
	go doitFsPathSSlice(cha, inp)
	return cha
}

func doitFsPathSFunc(done chan<- struct{}, inp <-chan fs.FsPathS, act func(a fs.FsPathS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsPathSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsPathSFunc(inp <-chan fs.FsPathS, act func(a fs.FsPathS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsPathS) { return }
	}
	go doitFsPathSFunc(cha, inp, act)
	return cha
}

func pipeFsPathSBuffer(out chan<- fs.FsPathS, inp <-chan fs.FsPathS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsPathSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsPathSBuffer(inp <-chan fs.FsPathS, cap int) (out <-chan fs.FsPathS) {
	cha := make(chan fs.FsPathS, cap)
	go pipeFsPathSBuffer(cha, inp)
	return cha
}

func pipeFsPathSFunc(out chan<- fs.FsPathS, inp <-chan fs.FsPathS, act func(a fs.FsPathS) fs.FsPathS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsPathSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsPathSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsPathSFunc(inp <-chan fs.FsPathS, act func(a fs.FsPathS) fs.FsPathS) (out <-chan fs.FsPathS) {
	cha := make(chan fs.FsPathS)
	if act == nil {
		act = func(a fs.FsPathS) fs.FsPathS { return a }
	}
	go pipeFsPathSFunc(cha, inp, act)
	return cha
}

func pipeFsPathSFork(out1, out2 chan<- fs.FsPathS, inp <-chan fs.FsPathS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsPathSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsPathSFork(inp <-chan fs.FsPathS) (out1, out2 <-chan fs.FsPathS) {
	cha1 := make(chan fs.FsPathS)
	cha2 := make(chan fs.FsPathS)
	go pipeFsPathSFork(cha1, cha2, inp)
	return cha1, cha2
}
