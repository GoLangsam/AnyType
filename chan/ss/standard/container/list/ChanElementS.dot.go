// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

// MakeElementSChan returns a new open channel
// (simply a 'chan []list.Element' that is).
//
// Note: No 'ElementS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myElementSPipelineStartsHere := MakeElementSChan()
//	// ... lot's of code to design and build Your favourite "myElementSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myElementSPipelineStartsHere <- drop
//	}
//	close(myElementSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeElementSBuffer) the channel is unbuffered.
//
func MakeElementSChan() (out chan []list.Element) {
	return make(chan []list.Element)
}

func sendElementS(out chan<- []list.Element, inp ...[]list.Element) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanElementS returns a channel to receive all inputs before close.
func ChanElementS(inp ...[]list.Element) (out <-chan []list.Element) {
	cha := make(chan []list.Element)
	go sendElementS(cha, inp...)
	return cha
}

func sendElementSSlice(out chan<- []list.Element, inp ...[][]list.Element) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanElementSSlice returns a channel to receive all inputs before close.
func ChanElementSSlice(inp ...[][]list.Element) (out <-chan []list.Element) {
	cha := make(chan []list.Element)
	go sendElementSSlice(cha, inp...)
	return cha
}

func joinElementS(done chan<- struct{}, out chan<- []list.Element, inp ...[]list.Element) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinElementS
func JoinElementS(out chan<- []list.Element, inp ...[]list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElementS(cha, out, inp...)
	return cha
}

func joinElementSSlice(done chan<- struct{}, out chan<- []list.Element, inp ...[][]list.Element) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinElementSSlice
func JoinElementSSlice(out chan<- []list.Element, inp ...[][]list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElementSSlice(cha, out, inp...)
	return cha
}

func joinElementSChan(done chan<- struct{}, out chan<- []list.Element, inp <-chan []list.Element) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinElementSChan
func JoinElementSChan(out chan<- []list.Element, inp <-chan []list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElementSChan(cha, out, inp)
	return cha
}

func doitElementS(done chan<- struct{}, inp <-chan []list.Element) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneElementS returns a channel to receive one signal before close after inp has been drained.
func DoneElementS(inp <-chan []list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitElementS(cha, inp)
	return cha
}

func doitElementSSlice(done chan<- ([][]list.Element), inp <-chan []list.Element) {
	defer close(done)
	ElementSS := [][]list.Element{}
	for i := range inp {
		ElementSS = append(ElementSS, i)
	}
	done <- ElementSS
}

// DoneElementSSlice returns a channel which will receive a slice
// of all the ElementSs received on inp channel before close.
// Unlike DoneElementS, a full slice is sent once, not just an event.
func DoneElementSSlice(inp <-chan []list.Element) (done <-chan ([][]list.Element)) {
	cha := make(chan ([][]list.Element))
	go doitElementSSlice(cha, inp)
	return cha
}

func doitElementSFunc(done chan<- struct{}, inp <-chan []list.Element, act func(a []list.Element)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneElementSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneElementSFunc(inp <-chan []list.Element, act func(a []list.Element)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []list.Element) { return }
	}
	go doitElementSFunc(cha, inp, act)
	return cha
}

func pipeElementSBuffer(out chan<- []list.Element, inp <-chan []list.Element) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeElementSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeElementSBuffer(inp <-chan []list.Element, cap int) (out <-chan []list.Element) {
	cha := make(chan []list.Element, cap)
	go pipeElementSBuffer(cha, inp)
	return cha
}

func pipeElementSFunc(out chan<- []list.Element, inp <-chan []list.Element, act func(a []list.Element) []list.Element) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeElementSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeElementSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeElementSFunc(inp <-chan []list.Element, act func(a []list.Element) []list.Element) (out <-chan []list.Element) {
	cha := make(chan []list.Element)
	if act == nil {
		act = func(a []list.Element) []list.Element { return a }
	}
	go pipeElementSFunc(cha, inp, act)
	return cha
}

func pipeElementSFork(out1, out2 chan<- []list.Element, inp <-chan []list.Element) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeElementSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeElementSFork(inp <-chan []list.Element) (out1, out2 <-chan []list.Element) {
	cha1 := make(chan []list.Element)
	cha2 := make(chan []list.Element)
	go pipeElementSFork(cha1, cha2, inp)
	return cha1, cha2
}
