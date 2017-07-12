// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// MakeFileInfoChan returns a new open channel
// (simply a 'chan os.FileInfo' that is).
//
// Note: No 'FileInfo-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFileInfoPipelineStartsHere := MakeFileInfoChan()
//	// ... lot's of code to design and build Your favourite "myFileInfoWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFileInfoPipelineStartsHere <- drop
//	}
//	close(myFileInfoPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFileInfoBuffer) the channel is unbuffered.
//
func MakeFileInfoChan() (out chan os.FileInfo) {
	return make(chan os.FileInfo)
}

// ChanFileInfo returns a channel to receive all inputs before close.
func ChanFileInfo(inp ...os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	go func(out chan<- os.FileInfo, inp ...os.FileInfo) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanFileInfoSlice returns a channel to receive all inputs before close.
func ChanFileInfoSlice(inp ...[]os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	go func(out chan<- os.FileInfo, inp ...[]os.FileInfo) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinFileInfo
func JoinFileInfo(out chan<- os.FileInfo, inp ...os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.FileInfo, inp ...os.FileInfo) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFileInfoSlice
func JoinFileInfoSlice(out chan<- os.FileInfo, inp ...[]os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.FileInfo, inp ...[]os.FileInfo) {
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

// JoinFileInfoChan
func JoinFileInfoChan(out chan<- os.FileInfo, inp <-chan os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.FileInfo, inp <-chan os.FileInfo) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFileInfo returns a channel to receive one signal before close after inp has been drained.
func DoneFileInfo(inp <-chan os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan os.FileInfo) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFileInfoSlice returns a channel which will receive a slice
// of all the FileInfos received on inp channel before close.
// Unlike DoneFileInfo, a full slice is sent once, not just an event.
func DoneFileInfoSlice(inp <-chan os.FileInfo) (done <-chan []os.FileInfo) {
	cha := make(chan []os.FileInfo)
	go func(inp <-chan os.FileInfo, done chan<- []os.FileInfo) {
		defer close(done)
		FileInfoS := []os.FileInfo{}
		for i := range inp {
			FileInfoS = append(FileInfoS, i)
		}
		done <- FileInfoS
	}(inp, cha)
	return cha
}

// DoneFileInfoFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a os.FileInfo) { return }
	}
	go func(done chan<- struct{}, inp <-chan os.FileInfo, act func(a os.FileInfo)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFileInfoBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileInfoBuffer(inp <-chan os.FileInfo, cap int) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo, cap)
	go func(out chan<- os.FileInfo, inp <-chan os.FileInfo) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFileInfoFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileInfoMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo) os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	if act == nil {
		act = func(a os.FileInfo) os.FileInfo { return a }
	}
	go func(out chan<- os.FileInfo, inp <-chan os.FileInfo, act func(a os.FileInfo) os.FileInfo) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFileInfoFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileInfoFork(inp <-chan os.FileInfo) (out1, out2 <-chan os.FileInfo) {
	cha1 := make(chan os.FileInfo)
	cha2 := make(chan os.FileInfo)
	go func(out1, out2 chan<- os.FileInfo, inp <-chan os.FileInfo) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
