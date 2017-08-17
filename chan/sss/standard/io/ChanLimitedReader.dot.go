// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeLimitedReaderChan returns a new open channel
// (simply a 'chan *io.LimitedReader' that is).
//
// Note: No 'LimitedReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLimitedReaderPipelineStartsHere := MakeLimitedReaderChan()
//	// ... lot's of code to design and build Your favourite "myLimitedReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLimitedReaderPipelineStartsHere <- drop
//	}
//	close(myLimitedReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLimitedReaderBuffer) the channel is unbuffered.
//
func MakeLimitedReaderChan() (out chan *io.LimitedReader) {
	return make(chan *io.LimitedReader)
}

// ChanLimitedReader returns a channel to receive all inputs before close.
func ChanLimitedReader(inp ...*io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go func(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanLimitedReaderSlice returns a channel to receive all inputs before close.
func ChanLimitedReaderSlice(inp ...[]*io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go func(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinLimitedReader
func JoinLimitedReader(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.LimitedReader, inp ...*io.LimitedReader) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinLimitedReaderSlice
func JoinLimitedReaderSlice(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) {
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

// JoinLimitedReaderChan
func JoinLimitedReaderChan(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneLimitedReader returns a channel to receive one signal before close after inp has been drained.
func DoneLimitedReader(inp <-chan *io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *io.LimitedReader) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneLimitedReaderSlice returns a channel which will receive a slice
// of all the LimitedReaders received on inp channel before close.
// Unlike DoneLimitedReader, a full slice is sent once, not just an event.
func DoneLimitedReaderSlice(inp <-chan *io.LimitedReader) (done <-chan []*io.LimitedReader) {
	cha := make(chan []*io.LimitedReader)
	go func(inp <-chan *io.LimitedReader, done chan<- []*io.LimitedReader) {
		defer close(done)
		LimitedReaderS := []*io.LimitedReader{}
		for i := range inp {
			LimitedReaderS = append(LimitedReaderS, i)
		}
		done <- LimitedReaderS
	}(inp, cha)
	return cha
}

// DoneLimitedReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.LimitedReader) { return }
	}
	go func(done chan<- struct{}, inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeLimitedReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLimitedReaderBuffer(inp <-chan *io.LimitedReader, cap int) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader, cap)
	go func(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeLimitedReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLimitedReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	if act == nil {
		act = func(a *io.LimitedReader) *io.LimitedReader { return a }
	}
	go func(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeLimitedReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLimitedReaderFork(inp <-chan *io.LimitedReader) (out1, out2 <-chan *io.LimitedReader) {
	cha1 := make(chan *io.LimitedReader)
	cha2 := make(chan *io.LimitedReader)
	go func(out1, out2 chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// LimitedReaderTube is the signature for a pipe function.
type LimitedReaderTube func(inp <-chan *io.LimitedReader, out <-chan *io.LimitedReader)

// LimitedReaderdaisy returns a channel to receive all inp after having passed thru tube.
func LimitedReaderdaisy(inp <-chan *io.LimitedReader, tube LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go tube(inp, cha)
	return cha
}

// LimitedReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func LimitedReaderDaisyChain(inp <-chan *io.LimitedReader, tubes ...LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := inp
	for _, tube := range tubes {
		cha = LimitedReaderdaisy(cha, tube)
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
