// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/lsm"
)

// MakeLSMChan returns a new open channel
// (simply a 'chan lsm.LazyStringerMap' that is).
//
// Note: No 'LSM-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLSMPipelineStartsHere := MakeLSMChan()
//	// ... lot's of code to design and build Your favourite "myLSMWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLSMPipelineStartsHere <- drop
//	}
//	close(myLSMPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLSMBuffer) the channel is unbuffered.
//
func MakeLSMChan() (out chan lsm.LazyStringerMap) {
	return make(chan lsm.LazyStringerMap)
}

// ChanLSM returns a channel to receive all inputs before close.
func ChanLSM(inp ...lsm.LazyStringerMap) (out <-chan lsm.LazyStringerMap) {
	cha := make(chan lsm.LazyStringerMap)
	go func(out chan<- lsm.LazyStringerMap, inp ...lsm.LazyStringerMap) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanLSMSlice returns a channel to receive all inputs before close.
func ChanLSMSlice(inp ...[]lsm.LazyStringerMap) (out <-chan lsm.LazyStringerMap) {
	cha := make(chan lsm.LazyStringerMap)
	go func(out chan<- lsm.LazyStringerMap, inp ...[]lsm.LazyStringerMap) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinLSM
func JoinLSM(out chan<- lsm.LazyStringerMap, inp ...lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- lsm.LazyStringerMap, inp ...lsm.LazyStringerMap) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinLSMSlice
func JoinLSMSlice(out chan<- lsm.LazyStringerMap, inp ...[]lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- lsm.LazyStringerMap, inp ...[]lsm.LazyStringerMap) {
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

// JoinLSMChan
func JoinLSMChan(out chan<- lsm.LazyStringerMap, inp <-chan lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- lsm.LazyStringerMap, inp <-chan lsm.LazyStringerMap) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneLSM returns a channel to receive one signal before close after inp has been drained.
func DoneLSM(inp <-chan lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan lsm.LazyStringerMap) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneLSMSlice returns a channel which will receive a slice
// of all the LSMs received on inp channel before close.
// Unlike DoneLSM, a full slice is sent once, not just an event.
func DoneLSMSlice(inp <-chan lsm.LazyStringerMap) (done <-chan []lsm.LazyStringerMap) {
	cha := make(chan []lsm.LazyStringerMap)
	go func(inp <-chan lsm.LazyStringerMap, done chan<- []lsm.LazyStringerMap) {
		defer close(done)
		LSMS := []lsm.LazyStringerMap{}
		for i := range inp {
			LSMS = append(LSMS, i)
		}
		done <- LSMS
	}(inp, cha)
	return cha
}

// DoneLSMFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLSMFunc(inp <-chan lsm.LazyStringerMap, act func(a lsm.LazyStringerMap)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a lsm.LazyStringerMap) { return }
	}
	go func(done chan<- struct{}, inp <-chan lsm.LazyStringerMap, act func(a lsm.LazyStringerMap)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeLSMBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLSMBuffer(inp <-chan lsm.LazyStringerMap, cap int) (out <-chan lsm.LazyStringerMap) {
	cha := make(chan lsm.LazyStringerMap, cap)
	go func(out chan<- lsm.LazyStringerMap, inp <-chan lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeLSMFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLSMMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLSMFunc(inp <-chan lsm.LazyStringerMap, act func(a lsm.LazyStringerMap) lsm.LazyStringerMap) (out <-chan lsm.LazyStringerMap) {
	cha := make(chan lsm.LazyStringerMap)
	if act == nil {
		act = func(a lsm.LazyStringerMap) lsm.LazyStringerMap { return a }
	}
	go func(out chan<- lsm.LazyStringerMap, inp <-chan lsm.LazyStringerMap, act func(a lsm.LazyStringerMap) lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeLSMFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLSMFork(inp <-chan lsm.LazyStringerMap) (out1, out2 <-chan lsm.LazyStringerMap) {
	cha1 := make(chan lsm.LazyStringerMap)
	cha2 := make(chan lsm.LazyStringerMap)
	go func(out1, out2 chan<- lsm.LazyStringerMap, inp <-chan lsm.LazyStringerMap) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
