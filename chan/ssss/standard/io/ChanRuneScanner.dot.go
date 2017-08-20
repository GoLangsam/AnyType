// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeRuneScannerChan returns a new open channel
// (simply a 'chan io.RuneScanner' that is).
//
// Note: No 'RuneScanner-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myRuneScannerPipelineStartsHere := MakeRuneScannerChan()
//	// ... lot's of code to design and build Your favourite "myRuneScannerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myRuneScannerPipelineStartsHere <- drop
//	}
//	close(myRuneScannerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeRuneScannerBuffer) the channel is unbuffered.
//
func MakeRuneScannerChan() chan io.RuneScanner {
	return make(chan io.RuneScanner)
}

// ChanRuneScanner returns a channel to receive all inputs before close.
func ChanRuneScanner(inp ...io.RuneScanner) chan io.RuneScanner {
	out := make(chan io.RuneScanner)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanRuneScannerSlice returns a channel to receive all inputs before close.
func ChanRuneScannerSlice(inp ...[]io.RuneScanner) chan io.RuneScanner {
	out := make(chan io.RuneScanner)
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

// JoinRuneScanner sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScanner(out chan<- io.RuneScanner, inp ...io.RuneScanner) chan struct{} {
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

// JoinRuneScannerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScannerSlice(out chan<- io.RuneScanner, inp ...[]io.RuneScanner) chan struct{} {
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

// JoinRuneScannerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneScannerChan(out chan<- io.RuneScanner, inp <-chan io.RuneScanner) chan struct{} {
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

// DoneRuneScanner returns a channel to receive one signal before close after inp has been drained.
func DoneRuneScanner(inp <-chan io.RuneScanner) chan struct{} {
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

// DoneRuneScannerSlice returns a channel which will receive a slice
// of all the RuneScanners received on inp channel before close.
// Unlike DoneRuneScanner, a full slice is sent once, not just an event.
func DoneRuneScannerSlice(inp <-chan io.RuneScanner) chan []io.RuneScanner {
	done := make(chan []io.RuneScanner)
	go func() {
		defer close(done)
		RuneScannerS := []io.RuneScanner{}
		for i := range inp {
			RuneScannerS = append(RuneScannerS, i)
		}
		done <- RuneScannerS
	}()
	return done
}

// DoneRuneScannerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneRuneScannerFunc(inp <-chan io.RuneScanner, act func(a io.RuneScanner)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.RuneScanner) { return }
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

// PipeRuneScannerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeRuneScannerBuffer(inp <-chan io.RuneScanner, cap int) chan io.RuneScanner {
	out := make(chan io.RuneScanner, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeRuneScannerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeRuneScannerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeRuneScannerFunc(inp <-chan io.RuneScanner, act func(a io.RuneScanner) io.RuneScanner) chan io.RuneScanner {
	out := make(chan io.RuneScanner)
	if act == nil {
		act = func(a io.RuneScanner) io.RuneScanner { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeRuneScannerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeRuneScannerFork(inp <-chan io.RuneScanner) (chan io.RuneScanner, chan io.RuneScanner) {
	out1 := make(chan io.RuneScanner)
	out2 := make(chan io.RuneScanner)
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
