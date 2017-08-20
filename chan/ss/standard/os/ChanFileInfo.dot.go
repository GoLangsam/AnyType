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

func sendFileInfo(out chan<- os.FileInfo, inp ...os.FileInfo) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFileInfo returns a channel to receive all inputs before close.
func ChanFileInfo(inp ...os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	go sendFileInfo(cha, inp...)
	return cha
}

func sendFileInfoSlice(out chan<- os.FileInfo, inp ...[]os.FileInfo) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFileInfoSlice returns a channel to receive all inputs before close.
func ChanFileInfoSlice(inp ...[]os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	go sendFileInfoSlice(cha, inp...)
	return cha
}

func joinFileInfo(done chan<- struct{}, out chan<- os.FileInfo, inp ...os.FileInfo) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFileInfo sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfo(out chan<- os.FileInfo, inp ...os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFileInfo(cha, out, inp...)
	return cha
}

func joinFileInfoSlice(done chan<- struct{}, out chan<- os.FileInfo, inp ...[]os.FileInfo) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFileInfoSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfoSlice(out chan<- os.FileInfo, inp ...[]os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFileInfoSlice(cha, out, inp...)
	return cha
}

func joinFileInfoChan(done chan<- struct{}, out chan<- os.FileInfo, inp <-chan os.FileInfo) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFileInfoChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileInfoChan(out chan<- os.FileInfo, inp <-chan os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFileInfoChan(cha, out, inp)
	return cha
}

func doitFileInfo(done chan<- struct{}, inp <-chan os.FileInfo) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFileInfo returns a channel to receive one signal before close after inp has been drained.
func DoneFileInfo(inp <-chan os.FileInfo) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFileInfo(cha, inp)
	return cha
}

func doitFileInfoSlice(done chan<- ([]os.FileInfo), inp <-chan os.FileInfo) {
	defer close(done)
	FileInfoS := []os.FileInfo{}
	for i := range inp {
		FileInfoS = append(FileInfoS, i)
	}
	done <- FileInfoS
}

// DoneFileInfoSlice returns a channel which will receive a slice
// of all the FileInfos received on inp channel before close.
// Unlike DoneFileInfo, a full slice is sent once, not just an event.
func DoneFileInfoSlice(inp <-chan os.FileInfo) (done <-chan ([]os.FileInfo)) {
	cha := make(chan ([]os.FileInfo))
	go doitFileInfoSlice(cha, inp)
	return cha
}

func doitFileInfoFunc(done chan<- struct{}, inp <-chan os.FileInfo, act func(a os.FileInfo)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFileInfoFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a os.FileInfo) { return }
	}
	go doitFileInfoFunc(cha, inp, act)
	return cha
}

func pipeFileInfoBuffer(out chan<- os.FileInfo, inp <-chan os.FileInfo) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFileInfoBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileInfoBuffer(inp <-chan os.FileInfo, cap int) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo, cap)
	go pipeFileInfoBuffer(cha, inp)
	return cha
}

func pipeFileInfoFunc(out chan<- os.FileInfo, inp <-chan os.FileInfo, act func(a os.FileInfo) os.FileInfo) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFileInfoFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileInfoMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileInfoFunc(inp <-chan os.FileInfo, act func(a os.FileInfo) os.FileInfo) (out <-chan os.FileInfo) {
	cha := make(chan os.FileInfo)
	if act == nil {
		act = func(a os.FileInfo) os.FileInfo { return a }
	}
	go pipeFileInfoFunc(cha, inp, act)
	return cha
}

func pipeFileInfoFork(out1, out2 chan<- os.FileInfo, inp <-chan os.FileInfo) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFileInfoFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileInfoFork(inp <-chan os.FileInfo) (out1, out2 <-chan os.FileInfo) {
	cha1 := make(chan os.FileInfo)
	cha2 := make(chan os.FileInfo)
	go pipeFileInfoFork(cha1, cha2, inp)
	return cha1, cha2
}
