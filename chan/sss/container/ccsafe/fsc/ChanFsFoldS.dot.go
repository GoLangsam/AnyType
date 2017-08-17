// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
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

// ChanFsFoldS returns a channel to receive all inputs before close.
func ChanFsFoldS(inp ...fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	go func(out chan<- fs.FsFoldS, inp ...fs.FsFoldS) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanFsFoldSSlice returns a channel to receive all inputs before close.
func ChanFsFoldSSlice(inp ...[]fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	go func(out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinFsFoldS
func JoinFsFoldS(out chan<- fs.FsFoldS, inp ...fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFoldS, inp ...fs.FsFoldS) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsFoldSSlice
func JoinFsFoldSSlice(out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFoldS, inp ...[]fs.FsFoldS) {
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

// JoinFsFoldSChan
func JoinFsFoldSChan(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFsFoldS returns a channel to receive one signal before close after inp has been drained.
func DoneFsFoldS(inp <-chan fs.FsFoldS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan fs.FsFoldS) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFsFoldSSlice returns a channel which will receive a slice
// of all the FsFoldSs received on inp channel before close.
// Unlike DoneFsFoldS, a full slice is sent once, not just an event.
func DoneFsFoldSSlice(inp <-chan fs.FsFoldS) (done <-chan []fs.FsFoldS) {
	cha := make(chan []fs.FsFoldS)
	go func(inp <-chan fs.FsFoldS, done chan<- []fs.FsFoldS) {
		defer close(done)
		FsFoldSS := []fs.FsFoldS{}
		for i := range inp {
			FsFoldSS = append(FsFoldSS, i)
		}
		done <- FsFoldSS
	}(inp, cha)
	return cha
}

// DoneFsFoldSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFoldSFunc(inp <-chan fs.FsFoldS, act func(a fs.FsFoldS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsFoldS) { return }
	}
	go func(done chan<- struct{}, inp <-chan fs.FsFoldS, act func(a fs.FsFoldS)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFsFoldSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFoldSBuffer(inp <-chan fs.FsFoldS, cap int) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS, cap)
	go func(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFsFoldSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFoldSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFoldSFunc(inp <-chan fs.FsFoldS, act func(a fs.FsFoldS) fs.FsFoldS) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	if act == nil {
		act = func(a fs.FsFoldS) fs.FsFoldS { return a }
	}
	go func(out chan<- fs.FsFoldS, inp <-chan fs.FsFoldS, act func(a fs.FsFoldS) fs.FsFoldS) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFsFoldSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFoldSFork(inp <-chan fs.FsFoldS) (out1, out2 <-chan fs.FsFoldS) {
	cha1 := make(chan fs.FsFoldS)
	cha2 := make(chan fs.FsFoldS)
	go func(out1, out2 chan<- fs.FsFoldS, inp <-chan fs.FsFoldS) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// FsFoldSTube is the signature for a pipe function.
type FsFoldSTube func(inp <-chan fs.FsFoldS, out <-chan fs.FsFoldS)

// FsFoldSdaisy returns a channel to receive all inp after having passed thru tube.
func FsFoldSdaisy(inp <-chan fs.FsFoldS, tube FsFoldSTube) (out <-chan fs.FsFoldS) {
	cha := make(chan fs.FsFoldS)
	go tube(inp, cha)
	return cha
}

// FsFoldSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsFoldSDaisyChain(inp <-chan fs.FsFoldS, tubes ...FsFoldSTube) (out <-chan fs.FsFoldS) {
	cha := inp
	for _, tube := range tubes {
		cha = FsFoldSdaisy(cha, tube)
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
