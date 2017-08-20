// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsDataSChan returns a new open channel
// (simply a 'chan fs.FsDataS' that is).
//
// Note: No 'FsDataS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsDataSPipelineStartsHere := MakeFsDataSChan()
//	// ... lot's of code to design and build Your favourite "myFsDataSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsDataSPipelineStartsHere <- drop
//	}
//	close(myFsDataSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsDataSBuffer) the channel is unbuffered.
//
func MakeFsDataSChan() chan fs.FsDataS {
	return make(chan fs.FsDataS)
}

// ChanFsDataS returns a channel to receive all inputs before close.
func ChanFsDataS(inp ...fs.FsDataS) chan fs.FsDataS {
	out := make(chan fs.FsDataS)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsDataSSlice returns a channel to receive all inputs before close.
func ChanFsDataSSlice(inp ...[]fs.FsDataS) chan fs.FsDataS {
	out := make(chan fs.FsDataS)
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

// JoinFsDataS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsDataS(out chan<- fs.FsDataS, inp ...fs.FsDataS) chan struct{} {
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

// JoinFsDataSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsDataSSlice(out chan<- fs.FsDataS, inp ...[]fs.FsDataS) chan struct{} {
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

// JoinFsDataSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsDataSChan(out chan<- fs.FsDataS, inp <-chan fs.FsDataS) chan struct{} {
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

// DoneFsDataS returns a channel to receive one signal before close after inp has been drained.
func DoneFsDataS(inp <-chan fs.FsDataS) chan struct{} {
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

// DoneFsDataSSlice returns a channel which will receive a slice
// of all the FsDataSs received on inp channel before close.
// Unlike DoneFsDataS, a full slice is sent once, not just an event.
func DoneFsDataSSlice(inp <-chan fs.FsDataS) chan []fs.FsDataS {
	done := make(chan []fs.FsDataS)
	go func() {
		defer close(done)
		FsDataSS := []fs.FsDataS{}
		for i := range inp {
			FsDataSS = append(FsDataSS, i)
		}
		done <- FsDataSS
	}()
	return done
}

// DoneFsDataSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsDataSFunc(inp <-chan fs.FsDataS, act func(a fs.FsDataS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.FsDataS) { return }
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

// PipeFsDataSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsDataSBuffer(inp <-chan fs.FsDataS, cap int) chan fs.FsDataS {
	out := make(chan fs.FsDataS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsDataSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsDataSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsDataSFunc(inp <-chan fs.FsDataS, act func(a fs.FsDataS) fs.FsDataS) chan fs.FsDataS {
	out := make(chan fs.FsDataS)
	if act == nil {
		act = func(a fs.FsDataS) fs.FsDataS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsDataSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsDataSFork(inp <-chan fs.FsDataS) (chan fs.FsDataS, chan fs.FsDataS) {
	out1 := make(chan fs.FsDataS)
	out2 := make(chan fs.FsDataS)
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
