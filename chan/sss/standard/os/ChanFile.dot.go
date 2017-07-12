// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// MakeFileChan returns a new open channel
// (simply a 'chan *os.File' that is).
//
// Note: No 'File-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFilePipelineStartsHere := MakeFileChan()
//	// ... lot's of code to design and build Your favourite "myFileWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFilePipelineStartsHere <- drop
//	}
//	close(myFilePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFileBuffer) the channel is unbuffered.
//
func MakeFileChan() (out chan *os.File) {
	return make(chan *os.File)
}

// ChanFile returns a channel to receive all inputs before close.
func ChanFile(inp ...*os.File) (out <-chan *os.File) {
	cha := make(chan *os.File)
	go func(out chan<- *os.File, inp ...*os.File) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanFileSlice returns a channel to receive all inputs before close.
func ChanFileSlice(inp ...[]*os.File) (out <-chan *os.File) {
	cha := make(chan *os.File)
	go func(out chan<- *os.File, inp ...[]*os.File) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinFile
func JoinFile(out chan<- *os.File, inp ...*os.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *os.File, inp ...*os.File) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFileSlice
func JoinFileSlice(out chan<- *os.File, inp ...[]*os.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *os.File, inp ...[]*os.File) {
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

// JoinFileChan
func JoinFileChan(out chan<- *os.File, inp <-chan *os.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *os.File, inp <-chan *os.File) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFile returns a channel to receive one signal before close after inp has been drained.
func DoneFile(inp <-chan *os.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *os.File) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFileSlice returns a channel which will receive a slice
// of all the Files received on inp channel before close.
// Unlike DoneFile, a full slice is sent once, not just an event.
func DoneFileSlice(inp <-chan *os.File) (done <-chan []*os.File) {
	cha := make(chan []*os.File)
	go func(inp <-chan *os.File, done chan<- []*os.File) {
		defer close(done)
		FileS := []*os.File{}
		for i := range inp {
			FileS = append(FileS, i)
		}
		done <- FileS
	}(inp, cha)
	return cha
}

// DoneFileFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileFunc(inp <-chan *os.File, act func(a *os.File)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *os.File) { return }
	}
	go func(done chan<- struct{}, inp <-chan *os.File, act func(a *os.File)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFileBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileBuffer(inp <-chan *os.File, cap int) (out <-chan *os.File) {
	cha := make(chan *os.File, cap)
	go func(out chan<- *os.File, inp <-chan *os.File) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFileFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileFunc(inp <-chan *os.File, act func(a *os.File) *os.File) (out <-chan *os.File) {
	cha := make(chan *os.File)
	if act == nil {
		act = func(a *os.File) *os.File { return a }
	}
	go func(out chan<- *os.File, inp <-chan *os.File, act func(a *os.File) *os.File) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFileFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileFork(inp <-chan *os.File) (out1, out2 <-chan *os.File) {
	cha1 := make(chan *os.File)
	cha2 := make(chan *os.File)
	go func(out1, out2 chan<- *os.File, inp <-chan *os.File) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
