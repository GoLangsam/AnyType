// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsDataChan returns a new open channel
// (simply a 'chan *fs.FsData' that is).
//
// Note: No 'FsData-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsDataPipelineStartsHere := MakeFsDataChan()
//	// ... lot's of code to design and build Your favourite "myFsDataWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsDataPipelineStartsHere <- drop
//	}
//	close(myFsDataPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsDataBuffer) the channel is unbuffered.
//
func MakeFsDataChan() (out chan *fs.FsData) {
	return make(chan *fs.FsData)
}

// ChanFsData returns a channel to receive all inputs before close.
func ChanFsData(inp ...*fs.FsData) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData)
	go func(out chan<- *fs.FsData, inp ...*fs.FsData) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanFsDataSlice returns a channel to receive all inputs before close.
func ChanFsDataSlice(inp ...[]*fs.FsData) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData)
	go func(out chan<- *fs.FsData, inp ...[]*fs.FsData) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinFsData
func JoinFsData(out chan<- *fs.FsData, inp ...*fs.FsData) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsData, inp ...*fs.FsData) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsDataSlice
func JoinFsDataSlice(out chan<- *fs.FsData, inp ...[]*fs.FsData) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsData, inp ...[]*fs.FsData) {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsDataChan
func JoinFsDataChan(out chan<- *fs.FsData, inp <-chan *fs.FsData) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsData, inp <-chan *fs.FsData) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFsData returns a channel to receive one signal before close after inp has been drained.
func DoneFsData(inp <-chan *fs.FsData) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *fs.FsData) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFsDataSlice returns a channel which will receive a slice
// of all the FsDatas received on inp channel before close.
// Unlike DoneFsData, a full slice is sent once, not just an event.
func DoneFsDataSlice(inp <-chan *fs.FsData) (done <-chan []*fs.FsData) {
	cha := make(chan []*fs.FsData)
	go func(inp <-chan *fs.FsData, done chan<- []*fs.FsData) {
		defer close(done)
		FsDataS := []*fs.FsData{}
		for i := range inp {
			FsDataS = append(FsDataS, i)
		}
		done <- FsDataS
	}(inp, cha)
	return cha
}

// DoneFsDataFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsDataFunc(inp <-chan *fs.FsData, act func(a *fs.FsData)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsData) { return }
	}
	go func(done chan<- struct{}, inp <-chan *fs.FsData, act func(a *fs.FsData)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFsDataBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsDataBuffer(inp <-chan *fs.FsData, cap int) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData, cap)
	go func(out chan<- *fs.FsData, inp <-chan *fs.FsData) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFsDataFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsDataMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsDataFunc(inp <-chan *fs.FsData, act func(a *fs.FsData) *fs.FsData) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData)
	if act == nil {
		act = func(a *fs.FsData) *fs.FsData { return a }
	}
	go func(out chan<- *fs.FsData, inp <-chan *fs.FsData, act func(a *fs.FsData) *fs.FsData) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFsDataFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsDataFork(inp <-chan *fs.FsData) (out1, out2 <-chan *fs.FsData) {
	cha1 := make(chan *fs.FsData)
	cha2 := make(chan *fs.FsData)
	go func(out1, out2 chan<- *fs.FsData, inp <-chan *fs.FsData) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// FsDataTube is the signature for a pipe function.
type FsDataTube func(inp <-chan *fs.FsData, out <-chan *fs.FsData)

// FsDatadaisy returns a channel to receive all inp after having passed thru tube.
func FsDatadaisy(inp <-chan *fs.FsData, tube FsDataTube) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData)
	go tube(inp, cha)
	return cha
}

// FsDataDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsDataDaisyChain(inp <-chan *fs.FsData, tubes ...FsDataTube) (out <-chan *fs.FsData) {
	cha := inp
	for _, tube := range tubes {
		cha = FsDatadaisy(cha, tube)
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
