// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakeFsFileChan returns a new open channel
// (simply a 'chan *fs.FsFile' that is).
//
// Note: No 'FsFile-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFilePipelineStartsHere := MakeFsFileChan()
//	// ... lot's of code to design and build Your favourite "myFsFileWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFilePipelineStartsHere <- drop
//	}
//	close(myFsFilePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFileBuffer) the channel is unbuffered.
//
func MakeFsFileChan() chan *fs.FsFile {
	return make(chan *fs.FsFile)
}

// ChanFsFile returns a channel to receive all inputs before close.
func ChanFsFile(inp ...*fs.FsFile) chan *fs.FsFile {
	out := make(chan *fs.FsFile)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanFsFileSlice returns a channel to receive all inputs before close.
func ChanFsFileSlice(inp ...[]*fs.FsFile) chan *fs.FsFile {
	out := make(chan *fs.FsFile)
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

// JoinFsFile
func JoinFsFile(out chan<- *fs.FsFile, inp ...*fs.FsFile) chan struct{} {
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

// JoinFsFileSlice
func JoinFsFileSlice(out chan<- *fs.FsFile, inp ...[]*fs.FsFile) chan struct{} {
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

// JoinFsFileChan
func JoinFsFileChan(out chan<- *fs.FsFile, inp <-chan *fs.FsFile) chan struct{} {
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

// DoneFsFile returns a channel to receive one signal before close after inp has been drained.
func DoneFsFile(inp <-chan *fs.FsFile) chan struct{} {
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

// DoneFsFileSlice returns a channel which will receive a slice
// of all the FsFiles received on inp channel before close.
// Unlike DoneFsFile, a full slice is sent once, not just an event.
func DoneFsFileSlice(inp <-chan *fs.FsFile) chan []*fs.FsFile {
	done := make(chan []*fs.FsFile)
	go func() {
		defer close(done)
		FsFileS := []*fs.FsFile{}
		for i := range inp {
			FsFileS = append(FsFileS, i)
		}
		done <- FsFileS
	}()
	return done
}

// DoneFsFileFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFileFunc(inp <-chan *fs.FsFile, act func(a *fs.FsFile)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsFile) { return }
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

// PipeFsFileBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFileBuffer(inp <-chan *fs.FsFile, cap int) chan *fs.FsFile {
	out := make(chan *fs.FsFile, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsFileFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFileMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFileFunc(inp <-chan *fs.FsFile, act func(a *fs.FsFile) *fs.FsFile) chan *fs.FsFile {
	out := make(chan *fs.FsFile)
	if act == nil {
		act = func(a *fs.FsFile) *fs.FsFile { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsFileFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFileFork(inp <-chan *fs.FsFile) (chan *fs.FsFile, chan *fs.FsFile) {
	out1 := make(chan *fs.FsFile)
	out2 := make(chan *fs.FsFile)
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
