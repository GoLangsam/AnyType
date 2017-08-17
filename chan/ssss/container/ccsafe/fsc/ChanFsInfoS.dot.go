// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsInfoSChan returns a new open channel
// (simply a 'chan fs.FsInfoS' that is).
//
// Note: No 'FsInfoS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsInfoSPipelineStartsHere := MakeFsInfoSChan()
//	// ... lot's of code to design and build Your favourite "myFsInfoSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsInfoSPipelineStartsHere <- drop
//	}
//	close(myFsInfoSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsInfoSBuffer) the channel is unbuffered.
//
func MakeFsInfoSChan() chan fs.FsInfoS {
	return make(chan fs.FsInfoS)
}

// ChanFsInfoS returns a channel to receive all inputs before close.
func ChanFsInfoS(inp ...fs.FsInfoS) chan fs.FsInfoS {
	out := make(chan fs.FsInfoS)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsInfoSSlice returns a channel to receive all inputs before close.
func ChanFsInfoSSlice(inp ...[]fs.FsInfoS) chan fs.FsInfoS {
	out := make(chan fs.FsInfoS)
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

// JoinFsInfoS
func JoinFsInfoS(out chan<- fs.FsInfoS, inp ...fs.FsInfoS) chan struct{} {
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

// JoinFsInfoSSlice
func JoinFsInfoSSlice(out chan<- fs.FsInfoS, inp ...[]fs.FsInfoS) chan struct{} {
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

// JoinFsInfoSChan
func JoinFsInfoSChan(out chan<- fs.FsInfoS, inp <-chan fs.FsInfoS) chan struct{} {
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

// DoneFsInfoS returns a channel to receive one signal before close after inp has been drained.
func DoneFsInfoS(inp <-chan fs.FsInfoS) chan struct{} {
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

// DoneFsInfoSSlice returns a channel which will receive a slice
// of all the FsInfoSs received on inp channel before close.
// Unlike DoneFsInfoS, a full slice is sent once, not just an event.
func DoneFsInfoSSlice(inp <-chan fs.FsInfoS) chan []fs.FsInfoS {
	done := make(chan []fs.FsInfoS)
	go func() {
		defer close(done)
		FsInfoSS := []fs.FsInfoS{}
		for i := range inp {
			FsInfoSS = append(FsInfoSS, i)
		}
		done <- FsInfoSS
	}()
	return done
}

// DoneFsInfoSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsInfoSFunc(inp <-chan fs.FsInfoS, act func(a fs.FsInfoS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.FsInfoS) { return }
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

// PipeFsInfoSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsInfoSBuffer(inp <-chan fs.FsInfoS, cap int) chan fs.FsInfoS {
	out := make(chan fs.FsInfoS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsInfoSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsInfoSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsInfoSFunc(inp <-chan fs.FsInfoS, act func(a fs.FsInfoS) fs.FsInfoS) chan fs.FsInfoS {
	out := make(chan fs.FsInfoS)
	if act == nil {
		act = func(a fs.FsInfoS) fs.FsInfoS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsInfoSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsInfoSFork(inp <-chan fs.FsInfoS) (chan fs.FsInfoS, chan fs.FsInfoS) {
	out1 := make(chan fs.FsInfoS)
	out2 := make(chan fs.FsInfoS)
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
