// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

// MakeFileChan returns a new open channel
// (simply a 'chan zip.File' that is).
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
func MakeFileChan() (out chan zip.File) {
	return make(chan zip.File)
}

func sendFile(out chan<- zip.File, inp ...zip.File) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFile returns a channel to receive all inputs before close.
func ChanFile(inp ...zip.File) (out <-chan zip.File) {
	cha := make(chan zip.File)
	go sendFile(cha, inp...)
	return cha
}

func sendFileSlice(out chan<- zip.File, inp ...[]zip.File) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFileSlice returns a channel to receive all inputs before close.
func ChanFileSlice(inp ...[]zip.File) (out <-chan zip.File) {
	cha := make(chan zip.File)
	go sendFileSlice(cha, inp...)
	return cha
}

func joinFile(done chan<- struct{}, out chan<- zip.File, inp ...zip.File) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFile
func JoinFile(out chan<- zip.File, inp ...zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFile(cha, out, inp...)
	return cha
}

func joinFileSlice(done chan<- struct{}, out chan<- zip.File, inp ...[]zip.File) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFileSlice
func JoinFileSlice(out chan<- zip.File, inp ...[]zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFileSlice(cha, out, inp...)
	return cha
}

func joinFileChan(done chan<- struct{}, out chan<- zip.File, inp <-chan zip.File) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFileChan
func JoinFileChan(out chan<- zip.File, inp <-chan zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFileChan(cha, out, inp)
	return cha
}

func doitFile(done chan<- struct{}, inp <-chan zip.File) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFile returns a channel to receive one signal before close after inp has been drained.
func DoneFile(inp <-chan zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFile(cha, inp)
	return cha
}

func doitFileSlice(done chan<- ([]zip.File), inp <-chan zip.File) {
	defer close(done)
	FileS := []zip.File{}
	for i := range inp {
		FileS = append(FileS, i)
	}
	done <- FileS
}

// DoneFileSlice returns a channel which will receive a slice
// of all the Files received on inp channel before close.
// Unlike DoneFile, a full slice is sent once, not just an event.
func DoneFileSlice(inp <-chan zip.File) (done <-chan ([]zip.File)) {
	cha := make(chan ([]zip.File))
	go doitFileSlice(cha, inp)
	return cha
}

func doitFileFunc(done chan<- struct{}, inp <-chan zip.File, act func(a zip.File)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFileFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileFunc(inp <-chan zip.File, act func(a zip.File)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a zip.File) { return }
	}
	go doitFileFunc(cha, inp, act)
	return cha
}

func pipeFileBuffer(out chan<- zip.File, inp <-chan zip.File) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFileBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileBuffer(inp <-chan zip.File, cap int) (out <-chan zip.File) {
	cha := make(chan zip.File, cap)
	go pipeFileBuffer(cha, inp)
	return cha
}

func pipeFileFunc(out chan<- zip.File, inp <-chan zip.File, act func(a zip.File) zip.File) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFileFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileFunc(inp <-chan zip.File, act func(a zip.File) zip.File) (out <-chan zip.File) {
	cha := make(chan zip.File)
	if act == nil {
		act = func(a zip.File) zip.File { return a }
	}
	go pipeFileFunc(cha, inp, act)
	return cha
}

func pipeFileFork(out1, out2 chan<- zip.File, inp <-chan zip.File) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFileFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileFork(inp <-chan zip.File) (out1, out2 <-chan zip.File) {
	cha1 := make(chan zip.File)
	cha2 := make(chan zip.File)
	go pipeFileFork(cha1, cha2, inp)
	return cha1, cha2
}