// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsInfoSChan returns a new open channel
// (simply a 'chan fs.FsInfoS' that is).
//
// Note: No 'FsInfoS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsInfoSPipelineStartsHere := MakeFsInfoSChan()
//	// ... lot's of code to design and build Your favourite "myFsInfoSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsInfoSPipelineStartsHere <- drop
//	}
//	close(myFsInfoSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsInfoSBuffer) the channel is unbuffered.
//
func MakeFsInfoSChan() (out chan fs.FsInfoS) {
	return make(chan fs.FsInfoS)
}

func sendFsInfoS(out chan<- fs.FsInfoS, inp ...fs.FsInfoS) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanFsInfoS returns a channel to receive all inputs before close.
func ChanFsInfoS(inp ...fs.FsInfoS) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	go sendFsInfoS(cha, inp...)
	return cha
}

func sendFsInfoSSlice(out chan<- fs.FsInfoS, inp ...[]fs.FsInfoS) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanFsInfoSSlice returns a channel to receive all inputs before close.
func ChanFsInfoSSlice(inp ...[]fs.FsInfoS) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	go sendFsInfoSSlice(cha, inp...)
	return cha
}

func chanFsInfoSFuncNok(out chan<- fs.FsInfoS, act func() (fs.FsInfoS, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanFsInfoSFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsInfoSFuncNok(act func() (fs.FsInfoS, bool)) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	go chanFsInfoSFuncNok(cha, act)
	return cha
}

func chanFsInfoSFuncErr(out chan<- fs.FsInfoS, act func() (fs.FsInfoS, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanFsInfoSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsInfoSFuncErr(act func() (fs.FsInfoS, error)) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	go chanFsInfoSFuncErr(cha, act)
	return cha
}

func joinFsInfoS(done chan<- struct{}, out chan<- fs.FsInfoS, inp ...fs.FsInfoS) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinFsInfoS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsInfoS(out chan<- fs.FsInfoS, inp ...fs.FsInfoS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsInfoS(cha, out, inp...)
	return cha
}

func joinFsInfoSSlice(done chan<- struct{}, out chan<- fs.FsInfoS, inp ...[]fs.FsInfoS) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinFsInfoSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsInfoSSlice(out chan<- fs.FsInfoS, inp ...[]fs.FsInfoS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsInfoSSlice(cha, out, inp...)
	return cha
}

func joinFsInfoSChan(done chan<- struct{}, out chan<- fs.FsInfoS, inp <-chan fs.FsInfoS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsInfoSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsInfoSChan(out chan<- fs.FsInfoS, inp <-chan fs.FsInfoS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsInfoSChan(cha, out, inp)
	return cha
}

func doitFsInfoS(done chan<- struct{}, inp <-chan fs.FsInfoS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsInfoS returns a channel to receive one signal before close after inp has been drained.
func DoneFsInfoS(inp <-chan fs.FsInfoS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsInfoS(cha, inp)
	return cha
}

func doitFsInfoSSlice(done chan<- ([]fs.FsInfoS), inp <-chan fs.FsInfoS) {
	defer close(done)
	FsInfoSS := []fs.FsInfoS{}
	for i := range inp {
		FsInfoSS = append(FsInfoSS, i)
	}
	done <- FsInfoSS
}

// DoneFsInfoSSlice returns a channel which will receive a slice
// of all the FsInfoSs received on inp channel before close.
// Unlike DoneFsInfoS, a full slice is sent once, not just an event.
func DoneFsInfoSSlice(inp <-chan fs.FsInfoS) (done <-chan ([]fs.FsInfoS)) {
	cha := make(chan ([]fs.FsInfoS))
	go doitFsInfoSSlice(cha, inp)
	return cha
}

func doitFsInfoSFunc(done chan<- struct{}, inp <-chan fs.FsInfoS, act func(a fs.FsInfoS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsInfoSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsInfoSFunc(inp <-chan fs.FsInfoS, act func(a fs.FsInfoS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsInfoS) { return }
	}
	go doitFsInfoSFunc(cha, inp, act)
	return cha
}

func pipeFsInfoSBuffer(out chan<- fs.FsInfoS, inp <-chan fs.FsInfoS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsInfoSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsInfoSBuffer(inp <-chan fs.FsInfoS, cap int) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS, cap)
	go pipeFsInfoSBuffer(cha, inp)
	return cha
}

func pipeFsInfoSFunc(out chan<- fs.FsInfoS, inp <-chan fs.FsInfoS, act func(a fs.FsInfoS) fs.FsInfoS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsInfoSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsInfoSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsInfoSFunc(inp <-chan fs.FsInfoS, act func(a fs.FsInfoS) fs.FsInfoS) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	if act == nil {
		act = func(a fs.FsInfoS) fs.FsInfoS { return a }
	}
	go pipeFsInfoSFunc(cha, inp, act)
	return cha
}

func pipeFsInfoSFork(out1, out2 chan<- fs.FsInfoS, inp <-chan fs.FsInfoS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsInfoSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsInfoSFork(inp <-chan fs.FsInfoS) (out1, out2 <-chan fs.FsInfoS) {
	cha1 := make(chan fs.FsInfoS)
	cha2 := make(chan fs.FsInfoS)
	go pipeFsInfoSFork(cha1, cha2, inp)
	return cha1, cha2
}

// FsInfoSTube is the signature for a pipe function.
type FsInfoSTube func(inp <-chan fs.FsInfoS, out <-chan fs.FsInfoS)

// FsInfoSDaisy returns a channel to receive all inp after having passed thru tube.
func FsInfoSDaisy(inp <-chan fs.FsInfoS, tube FsInfoSTube) (out <-chan fs.FsInfoS) {
	cha := make(chan fs.FsInfoS)
	go tube(inp, cha)
	return cha
}

// FsInfoSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsInfoSDaisyChain(inp <-chan fs.FsInfoS, tubes ...FsInfoSTube) (out <-chan fs.FsInfoS) {
	cha := inp
	for i := range tubes {
		cha = FsInfoSDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
