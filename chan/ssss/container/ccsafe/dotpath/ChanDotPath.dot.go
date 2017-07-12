// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package dotpath

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/dotpath"
)

// MakeDotPathChan returns a new open channel
// (simply a 'chan dotpath.DotPath' that is).
//
// Note: No 'DotPath-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myDotPathPipelineStartsHere := MakeDotPathChan()
//	// ... lot's of code to design and build Your favourite "myDotPathWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myDotPathPipelineStartsHere <- drop
//	}
//	close(myDotPathPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeDotPathBuffer) the channel is unbuffered.
//
func MakeDotPathChan() chan dotpath.DotPath {
	return make(chan dotpath.DotPath)
}

// ChanDotPath returns a channel to receive all inputs before close.
func ChanDotPath(inp ...dotpath.DotPath) chan dotpath.DotPath {
	out := make(chan dotpath.DotPath)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanDotPathSlice returns a channel to receive all inputs before close.
func ChanDotPathSlice(inp ...[]dotpath.DotPath) chan dotpath.DotPath {
	out := make(chan dotpath.DotPath)
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

// JoinDotPath
func JoinDotPath(out chan<- dotpath.DotPath, inp ...dotpath.DotPath) chan struct{} {
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

// JoinDotPathSlice
func JoinDotPathSlice(out chan<- dotpath.DotPath, inp ...[]dotpath.DotPath) chan struct{} {
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

// JoinDotPathChan
func JoinDotPathChan(out chan<- dotpath.DotPath, inp <-chan dotpath.DotPath) chan struct{} {
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

// DoneDotPath returns a channel to receive one signal before close after inp has been drained.
func DoneDotPath(inp <-chan dotpath.DotPath) chan struct{} {
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

// DoneDotPathSlice returns a channel which will receive a slice
// of all the DotPaths received on inp channel before close.
// Unlike DoneDotPath, a full slice is sent once, not just an event.
func DoneDotPathSlice(inp <-chan dotpath.DotPath) chan []dotpath.DotPath {
	done := make(chan []dotpath.DotPath)
	go func() {
		defer close(done)
		DotPathS := []dotpath.DotPath{}
		for i := range inp {
			DotPathS = append(DotPathS, i)
		}
		done <- DotPathS
	}()
	return done
}

// DoneDotPathFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneDotPathFunc(inp <-chan dotpath.DotPath, act func(a dotpath.DotPath)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a dotpath.DotPath) { return }
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

// PipeDotPathBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeDotPathBuffer(inp <-chan dotpath.DotPath, cap int) chan dotpath.DotPath {
	out := make(chan dotpath.DotPath, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeDotPathFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeDotPathMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeDotPathFunc(inp <-chan dotpath.DotPath, act func(a dotpath.DotPath) dotpath.DotPath) chan dotpath.DotPath {
	out := make(chan dotpath.DotPath)
	if act == nil {
		act = func(a dotpath.DotPath) dotpath.DotPath { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeDotPathFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeDotPathFork(inp <-chan dotpath.DotPath) (chan dotpath.DotPath, chan dotpath.DotPath) {
	out1 := make(chan dotpath.DotPath)
	out2 := make(chan dotpath.DotPath)
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
