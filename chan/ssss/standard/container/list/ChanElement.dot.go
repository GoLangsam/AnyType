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
func MakeElementChan() chan list.Element {
	return make(chan list.Element)
}

// ChanElement returns a channel to receive all inputs before close.
func ChanElement(inp ...list.Element) chan list.Element {
	out := make(chan list.Element)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanElementSlice returns a channel to receive all inputs before close.
func ChanElementSlice(inp ...[]list.Element) chan list.Element {
	out := make(chan list.Element)
	go func() {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}()
	return out
}

// JoinElement sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElement(out chan<- list.Element, inp ...list.Element) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// JoinElementSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementSlice(out chan<- list.Element, inp ...[]list.Element) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinElementChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementChan(out chan<- list.Element, inp <-chan list.Element) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneElement returns a channel to receive one signal before close after inp has been drained.
func DoneElement(inp <-chan list.Element) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneElementSlice returns a channel which will receive a slice
// of all the Elements received on inp channel before close.
// Unlike DoneElement, a full slice is sent once, not just an event.
func DoneElementSlice(inp <-chan list.Element) chan []list.Element {
	done := make(chan []list.Element)
	go func() {
		defer close(done)
		ElementS := []list.Element{}
		for i := range inp {
			ElementS = append(ElementS, i)
		}
		done <- ElementS
	}()
	return done
}

// DoneElementFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneElementFunc(inp <-chan list.Element, act func(a list.Element)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a list.Element) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeElementBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeElementBuffer(inp <-chan list.Element, cap int) chan list.Element {
	out := make(chan list.Element, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeElementFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeElementMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeElementFunc(inp <-chan list.Element, act func(a list.Element) list.Element) chan list.Element {
	out := make(chan list.Element)
	if act == nil {
		act = func(a list.Element) list.Element { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeElementFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeElementFork(inp <-chan list.Element) (chan list.Element, chan list.Element) {
	out1 := make(chan list.Element)
	out2 := make(chan list.Element)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}
