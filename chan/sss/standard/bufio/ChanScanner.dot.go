// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// MakeScannerChan returns a new open channel
// (simply a 'chan *bufio.Scanner' that is).
//
// Note: No 'Scanner-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myScannerPipelineStartsHere := MakeScannerChan()
//	// ... lot's of code to design and build Your favourite "myScannerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myScannerPipelineStartsHere <- drop
//	}
//	close(myScannerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeScannerBuffer) the channel is unbuffered.
//
func MakeScannerChan() (out chan *bufio.Scanner) {
	return make(chan *bufio.Scanner)
}

// ChanScanner returns a channel to receive all inputs before close.
func ChanScanner(inp ...*bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, inp ...*bufio.Scanner) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanScannerSlice returns a channel to receive all inputs before close.
func ChanScannerSlice(inp ...[]*bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinScanner
func JoinScanner(out chan<- *bufio.Scanner, inp ...*bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp ...*bufio.Scanner) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinScannerSlice
func JoinScannerSlice(out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp ...[]*bufio.Scanner) {
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

// JoinScannerChan
func JoinScannerChan(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneScanner returns a channel to receive one signal before close after inp has been drained.
func DoneScanner(inp <-chan *bufio.Scanner) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *bufio.Scanner) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneScannerSlice returns a channel which will receive a slice
// of all the Scanners received on inp channel before close.
// Unlike DoneScanner, a full slice is sent once, not just an event.
func DoneScannerSlice(inp <-chan *bufio.Scanner) (done <-chan []*bufio.Scanner) {
	cha := make(chan []*bufio.Scanner)
	go func(inp <-chan *bufio.Scanner, done chan<- []*bufio.Scanner) {
		defer close(done)
		ScannerS := []*bufio.Scanner{}
		for i := range inp {
			ScannerS = append(ScannerS, i)
		}
		done <- ScannerS
	}(inp, cha)
	return cha
}

// DoneScannerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneScannerFunc(inp <-chan *bufio.Scanner, act func(a *bufio.Scanner)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.Scanner) { return }
	}
	go func(done chan<- struct{}, inp <-chan *bufio.Scanner, act func(a *bufio.Scanner)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeScannerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeScannerBuffer(inp <-chan *bufio.Scanner, cap int) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner, cap)
	go func(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeScannerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeScannerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeScannerFunc(inp <-chan *bufio.Scanner, act func(a *bufio.Scanner) *bufio.Scanner) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	if act == nil {
		act = func(a *bufio.Scanner) *bufio.Scanner { return a }
	}
	go func(out chan<- *bufio.Scanner, inp <-chan *bufio.Scanner, act func(a *bufio.Scanner) *bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeScannerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeScannerFork(inp <-chan *bufio.Scanner) (out1, out2 <-chan *bufio.Scanner) {
	cha1 := make(chan *bufio.Scanner)
	cha2 := make(chan *bufio.Scanner)
	go func(out1, out2 chan<- *bufio.Scanner, inp <-chan *bufio.Scanner) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
