// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsBaseChan returns a new open channel
// (simply a 'chan *fs.FsBase' that is).
//
// Note: No 'FsBase-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsBasePipelineStartsHere := MakeFsBaseChan()
//	// ... lot's of code to design and build Your favourite "myFsBaseWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsBasePipelineStartsHere <- drop
//	}
//	close(myFsBasePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsBaseBuffer) the channel is unbuffered.
//
func MakeFsBaseChan() chan *fs.FsBase {
	return make(chan *fs.FsBase)
}

// ChanFsBase returns a channel to receive all inputs before close.
func ChanFsBase(inp ...*fs.FsBase) chan *fs.FsBase {
	out := make(chan *fs.FsBase)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsBaseSlice returns a channel to receive all inputs before close.
func ChanFsBaseSlice(inp ...[]*fs.FsBase) chan *fs.FsBase {
	out := make(chan *fs.FsBase)
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

// JoinFsBase
func JoinFsBase(out chan<- *fs.FsBase, inp ...*fs.FsBase) chan struct{} {
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

// JoinFsBaseSlice
func JoinFsBaseSlice(out chan<- *fs.FsBase, inp ...[]*fs.FsBase) chan struct{} {
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

// JoinFsBaseChan
func JoinFsBaseChan(out chan<- *fs.FsBase, inp <-chan *fs.FsBase) chan struct{} {
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

// DoneFsBase returns a channel to receive one signal before close after inp has been drained.
func DoneFsBase(inp <-chan *fs.FsBase) chan struct{} {
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

// DoneFsBaseSlice returns a channel which will receive a slice
// of all the FsBases received on inp channel before close.
// Unlike DoneFsBase, a full slice is sent once, not just an event.
func DoneFsBaseSlice(inp <-chan *fs.FsBase) chan []*fs.FsBase {
	done := make(chan []*fs.FsBase)
	go func() {
		defer close(done)
		FsBaseS := []*fs.FsBase{}
		for i := range inp {
			FsBaseS = append(FsBaseS, i)
		}
		done <- FsBaseS
	}()
	return done
}

// DoneFsBaseFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsBaseFunc(inp <-chan *fs.FsBase, act func(a *fs.FsBase)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsBase) { return }
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

// PipeFsBaseBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsBaseBuffer(inp <-chan *fs.FsBase, cap int) chan *fs.FsBase {
	out := make(chan *fs.FsBase, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsBaseFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsBaseMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsBaseFunc(inp <-chan *fs.FsBase, act func(a *fs.FsBase) *fs.FsBase) chan *fs.FsBase {
	out := make(chan *fs.FsBase)
	if act == nil {
		act = func(a *fs.FsBase) *fs.FsBase { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsBaseFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsBaseFork(inp <-chan *fs.FsBase) (chan *fs.FsBase, chan *fs.FsBase) {
	out1 := make(chan *fs.FsBase)
	out2 := make(chan *fs.FsBase)
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
