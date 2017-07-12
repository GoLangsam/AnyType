// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// MakeListSChan returns a new open channel
// (simply a 'chan []list.List' that is).
//
// Note: No 'ListS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myListSPipelineStartsHere := MakeListSChan()
//	// ... lot's of code to design and build Your favourite "myListSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myListSPipelineStartsHere <- drop
//	}
//	close(myListSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeListSBuffer) the channel is unbuffered.
//
func MakeListSChan() (out chan []list.List) {
	return make(chan []list.List)
}

// ChanListS returns a channel to receive all inputs before close.
func ChanListS(inp ...[]list.List) (out <-chan []list.List) {
	cha := make(chan []list.List)
	go func(out chan<- []list.List, inp ...[]list.List) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanListSSlice returns a channel to receive all inputs before close.
func ChanListSSlice(inp ...[][]list.List) (out <-chan []list.List) {
	cha := make(chan []list.List)
	go func(out chan<- []list.List, inp ...[][]list.List) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinListS
func JoinListS(out chan<- []list.List, inp ...[]list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []list.List, inp ...[]list.List) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinListSSlice
func JoinListSSlice(out chan<- []list.List, inp ...[][]list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []list.List, inp ...[][]list.List) {
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

// JoinListSChan
func JoinListSChan(out chan<- []list.List, inp <-chan []list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []list.List, inp <-chan []list.List) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneListS returns a channel to receive one signal before close after inp has been drained.
func DoneListS(inp <-chan []list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []list.List) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneListSSlice returns a channel which will receive a slice
// of all the ListSs received on inp channel before close.
// Unlike DoneListS, a full slice is sent once, not just an event.
func DoneListSSlice(inp <-chan []list.List) (done <-chan [][]list.List) {
	cha := make(chan [][]list.List)
	go func(inp <-chan []list.List, done chan<- [][]list.List) {
		defer close(done)
		ListSS := [][]list.List{}
		for i := range inp {
			ListSS = append(ListSS, i)
		}
		done <- ListSS
	}(inp, cha)
	return cha
}

// DoneListSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneListSFunc(inp <-chan []list.List, act func(a []list.List)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []list.List) { return }
	}
	go func(done chan<- struct{}, inp <-chan []list.List, act func(a []list.List)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeListSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeListSBuffer(inp <-chan []list.List, cap int) (out <-chan []list.List) {
	cha := make(chan []list.List, cap)
	go func(out chan<- []list.List, inp <-chan []list.List) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeListSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeListSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeListSFunc(inp <-chan []list.List, act func(a []list.List) []list.List) (out <-chan []list.List) {
	cha := make(chan []list.List)
	if act == nil {
		act = func(a []list.List) []list.List { return a }
	}
	go func(out chan<- []list.List, inp <-chan []list.List, act func(a []list.List) []list.List) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeListSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeListSFork(inp <-chan []list.List) (out1, out2 <-chan []list.List) {
	cha1 := make(chan []list.List)
	cha2 := make(chan []list.List)
	go func(out1, out2 chan<- []list.List, inp <-chan []list.List) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
