// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeSectionReaderChan returns a new open channel
// (simply a 'chan *io.SectionReader' that is).
//
// Note: No 'SectionReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySectionReaderPipelineStartsHere := MakeSectionReaderChan()
//	// ... lot's of code to design and build Your favourite "mySectionReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySectionReaderPipelineStartsHere <- drop
//	}
//	close(mySectionReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSectionReaderBuffer) the channel is unbuffered.
//
func MakeSectionReaderChan() (out chan *io.SectionReader) {
	return make(chan *io.SectionReader)
}

// ChanSectionReader returns a channel to receive all inputs before close.
func ChanSectionReader(inp ...*io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go func(out chan<- *io.SectionReader, inp ...*io.SectionReader) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanSectionReaderSlice returns a channel to receive all inputs before close.
func ChanSectionReaderSlice(inp ...[]*io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go func(out chan<- *io.SectionReader, inp ...[]*io.SectionReader) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinSectionReader
func JoinSectionReader(out chan<- *io.SectionReader, inp ...*io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.SectionReader, inp ...*io.SectionReader) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSectionReaderSlice
func JoinSectionReaderSlice(out chan<- *io.SectionReader, inp ...[]*io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.SectionReader, inp ...[]*io.SectionReader) {
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

// JoinSectionReaderChan
func JoinSectionReaderChan(out chan<- *io.SectionReader, inp <-chan *io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneSectionReader returns a channel to receive one signal before close after inp has been drained.
func DoneSectionReader(inp <-chan *io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *io.SectionReader) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSectionReaderSlice returns a channel which will receive a slice
// of all the SectionReaders received on inp channel before close.
// Unlike DoneSectionReader, a full slice is sent once, not just an event.
func DoneSectionReaderSlice(inp <-chan *io.SectionReader) (done <-chan []*io.SectionReader) {
	cha := make(chan []*io.SectionReader)
	go func(inp <-chan *io.SectionReader, done chan<- []*io.SectionReader) {
		defer close(done)
		SectionReaderS := []*io.SectionReader{}
		for i := range inp {
			SectionReaderS = append(SectionReaderS, i)
		}
		done <- SectionReaderS
	}(inp, cha)
	return cha
}

// DoneSectionReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.SectionReader) { return }
	}
	go func(done chan<- struct{}, inp <-chan *io.SectionReader, act func(a *io.SectionReader)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeSectionReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSectionReaderBuffer(inp <-chan *io.SectionReader, cap int) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader, cap)
	go func(out chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeSectionReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSectionReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader) *io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	if act == nil {
		act = func(a *io.SectionReader) *io.SectionReader { return a }
	}
	go func(out chan<- *io.SectionReader, inp <-chan *io.SectionReader, act func(a *io.SectionReader) *io.SectionReader) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeSectionReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSectionReaderFork(inp <-chan *io.SectionReader) (out1, out2 <-chan *io.SectionReader) {
	cha1 := make(chan *io.SectionReader)
	cha2 := make(chan *io.SectionReader)
	go func(out1, out2 chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
