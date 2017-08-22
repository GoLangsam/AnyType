// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsBaseChan returns a new open channel
// (simply a 'chan *fs.FsBase' that is).
//
// Note: No 'FsBase-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsBasePipelineStartsHere := MakeFsBaseChan()
//	// ... lot's of code to design and build Your favourite "myFsBaseWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsBasePipelineStartsHere <- drop
//	}
//	close(myFsBasePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsBaseBuffer) the channel is unbuffered.
//
func MakeFsBaseChan() (out chan *fs.FsBase) {
	return make(chan *fs.FsBase)
}

func sendFsBase(out chan<- *fs.FsBase, inp ...*fs.FsBase) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanFsBase returns a channel to receive all inputs before close.
func ChanFsBase(inp ...*fs.FsBase) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go sendFsBase(cha, inp...)
	return cha
}

func sendFsBaseSlice(out chan<- *fs.FsBase, inp ...[]*fs.FsBase) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanFsBaseSlice returns a channel to receive all inputs before close.
func ChanFsBaseSlice(inp ...[]*fs.FsBase) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go sendFsBaseSlice(cha, inp...)
	return cha
}

func chanFsBaseFuncNil(out chan<- *fs.FsBase, act func() *fs.FsBase) {
	defer close(out)
	for {
		res := act() // Apply action
		if res == nil {
			return
		}
		out <- res
	}
}

// ChanFsBaseFuncNil returns a channel to receive all results of act until nil before close.
func ChanFsBaseFuncNil(act func() *fs.FsBase) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go chanFsBaseFuncNil(cha, act)
	return cha
}

func chanFsBaseFuncNok(out chan<- *fs.FsBase, act func() (*fs.FsBase, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanFsBaseFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsBaseFuncNok(act func() (*fs.FsBase, bool)) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go chanFsBaseFuncNok(cha, act)
	return cha
}

func chanFsBaseFuncErr(out chan<- *fs.FsBase, act func() (*fs.FsBase, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanFsBaseFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsBaseFuncErr(act func() (*fs.FsBase, error)) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go chanFsBaseFuncErr(cha, act)
	return cha
}

func joinFsBase(done chan<- struct{}, out chan<- *fs.FsBase, inp ...*fs.FsBase) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinFsBase sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBase(out chan<- *fs.FsBase, inp ...*fs.FsBase) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBase(cha, out, inp...)
	return cha
}

func joinFsBaseSlice(done chan<- struct{}, out chan<- *fs.FsBase, inp ...[]*fs.FsBase) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinFsBaseSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBaseSlice(out chan<- *fs.FsBase, inp ...[]*fs.FsBase) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBaseSlice(cha, out, inp...)
	return cha
}

func joinFsBaseChan(done chan<- struct{}, out chan<- *fs.FsBase, inp <-chan *fs.FsBase) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsBaseChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBaseChan(out chan<- *fs.FsBase, inp <-chan *fs.FsBase) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBaseChan(cha, out, inp)
	return cha
}

func doitFsBase(done chan<- struct{}, inp <-chan *fs.FsBase) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsBase returns a channel to receive one signal before close after inp has been drained.
func DoneFsBase(inp <-chan *fs.FsBase) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsBase(cha, inp)
	return cha
}

func doitFsBaseSlice(done chan<- ([]*fs.FsBase), inp <-chan *fs.FsBase) {
	defer close(done)
	FsBaseS := []*fs.FsBase{}
	for i := range inp {
		FsBaseS = append(FsBaseS, i)
	}
	done <- FsBaseS
}

// DoneFsBaseSlice returns a channel which will receive a slice
// of all the FsBases received on inp channel before close.
// Unlike DoneFsBase, a full slice is sent once, not just an event.
func DoneFsBaseSlice(inp <-chan *fs.FsBase) (done <-chan ([]*fs.FsBase)) {
	cha := make(chan ([]*fs.FsBase))
	go doitFsBaseSlice(cha, inp)
	return cha
}

func doitFsBaseFunc(done chan<- struct{}, inp <-chan *fs.FsBase, act func(a *fs.FsBase)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsBaseFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsBaseFunc(inp <-chan *fs.FsBase, act func(a *fs.FsBase)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsBase) { return }
	}
	go doitFsBaseFunc(cha, inp, act)
	return cha
}

func pipeFsBaseBuffer(out chan<- *fs.FsBase, inp <-chan *fs.FsBase) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsBaseBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsBaseBuffer(inp <-chan *fs.FsBase, cap int) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase, cap)
	go pipeFsBaseBuffer(cha, inp)
	return cha
}

func pipeFsBaseFunc(out chan<- *fs.FsBase, inp <-chan *fs.FsBase, act func(a *fs.FsBase) *fs.FsBase) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsBaseFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsBaseMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsBaseFunc(inp <-chan *fs.FsBase, act func(a *fs.FsBase) *fs.FsBase) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	if act == nil {
		act = func(a *fs.FsBase) *fs.FsBase { return a }
	}
	go pipeFsBaseFunc(cha, inp, act)
	return cha
}

func pipeFsBaseFork(out1, out2 chan<- *fs.FsBase, inp <-chan *fs.FsBase) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsBaseFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsBaseFork(inp <-chan *fs.FsBase) (out1, out2 <-chan *fs.FsBase) {
	cha1 := make(chan *fs.FsBase)
	cha2 := make(chan *fs.FsBase)
	go pipeFsBaseFork(cha1, cha2, inp)
	return cha1, cha2
}

// FsBaseTube is the signature for a pipe function.
type FsBaseTube func(inp <-chan *fs.FsBase, out <-chan *fs.FsBase)

// FsBaseDaisy returns a channel to receive all inp after having passed thru tube.
func FsBaseDaisy(inp <-chan *fs.FsBase, tube FsBaseTube) (out <-chan *fs.FsBase) {
	cha := make(chan *fs.FsBase)
	go tube(inp, cha)
	return cha
}

// FsBaseDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsBaseDaisyChain(inp <-chan *fs.FsBase, tubes ...FsBaseTube) (out <-chan *fs.FsBase) {
	cha := inp
	for i := range tubes {
		cha = FsBaseDaisy(cha, tubes[i])
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
