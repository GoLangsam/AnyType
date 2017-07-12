// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// MakeElementChan returns a new open channel
// (simply a 'chan list.Element' that is).
//
// Note: No 'Element-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myElementPipelineStartsHere := MakeElementChan()
//	// ... lot's of code to design and build Your favourite "myElementWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myElementPipelineStartsHere <- drop
//	}
//	close(myElementPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeElementBuffer) the channel is unbuffered.
//
func MakeElementChan() (out chan list.Element) {
	return make(chan list.Element)
}

// ChanElement returns a channel to receive all inputs before close.
func ChanElement(inp ...list.Element) (out <-chan list.Element) {
	cha := make(chan list.Element)
	go func(out chan<- list.Element, inp ...list.Element) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanElementSlice returns a channel to receive all inputs before close.
func ChanElementSlice(inp ...[]list.Element) (out <-chan list.Element) {
	cha := make(chan list.Element)
	go func(out chan<- list.Element, inp ...[]list.Element) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinElement
func JoinElement(out chan<- list.Element, inp ...list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- list.Element, inp ...list.Element) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinElementSlice
func JoinElementSlice(out chan<- list.Element, inp ...[]list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- list.Element, inp ...[]list.Element) {
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

// JoinElementChan
func JoinElementChan(out chan<- list.Element, inp <-chan list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- list.Element, inp <-chan list.Element) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneElement returns a channel to receive one signal before close after inp has been drained.
func DoneElement(inp <-chan list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan list.Element) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneElementSlice returns a channel which will receive a slice
// of all the Elements received on inp channel before close.
// Unlike DoneElement, a full slice is sent once, not just an event.
func DoneElementSlice(inp <-chan list.Element) (done <-chan []list.Element) {
	cha := make(chan []list.Element)
	go func(inp <-chan list.Element, done chan<- []list.Element) {
		defer close(done)
		ElementS := []list.Element{}
		for i := range inp {
			ElementS = append(ElementS, i)
		}
		done <- ElementS
	}(inp, cha)
	return cha
}

// DoneElementFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneElementFunc(inp <-chan list.Element, act func(a list.Element)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a list.Element) { return }
	}
	go func(done chan<- struct{}, inp <-chan list.Element, act func(a list.Element)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeElementBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeElementBuffer(inp <-chan list.Element, cap int) (out <-chan list.Element) {
	cha := make(chan list.Element, cap)
	go func(out chan<- list.Element, inp <-chan list.Element) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeElementFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeElementMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeElementFunc(inp <-chan list.Element, act func(a list.Element) list.Element) (out <-chan list.Element) {
	cha := make(chan list.Element)
	if act == nil {
		act = func(a list.Element) list.Element { return a }
	}
	go func(out chan<- list.Element, inp <-chan list.Element, act func(a list.Element) list.Element) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeElementFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeElementFork(inp <-chan list.Element) (out1, out2 <-chan list.Element) {
	cha1 := make(chan list.Element)
	cha2 := make(chan list.Element)
	go func(out1, out2 chan<- list.Element, inp <-chan list.Element) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
