// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePointerSChan returns a new open channel
// (simply a 'chan []*SomeType' that is).
//
// Note: No 'PointerS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPointerSPipelineStartsHere := MakePointerSChan()
//	// ... lot's of code to design and build Your favourite "myPointerSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPointerSPipelineStartsHere <- drop
//	}
//	close(myPointerSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePointerSBuffer) the channel is unbuffered.
//
func MakePointerSChan() (out chan []*SomeType) {
	return make(chan []*SomeType)
}

func sendPointerS(out chan<- []*SomeType, inp ...[]*SomeType) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanPointerS returns a channel to receive all inputs before close.
func ChanPointerS(inp ...[]*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go sendPointerS(cha, inp...)
	return cha
}

func sendPointerSSlice(out chan<- []*SomeType, inp ...[][]*SomeType) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanPointerSSlice returns a channel to receive all inputs before close.
func ChanPointerSSlice(inp ...[][]*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go sendPointerSSlice(cha, inp...)
	return cha
}

func joinPointerS(done chan<- struct{}, out chan<- []*SomeType, inp ...[]*SomeType) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPointerS
func JoinPointerS(out chan<- []*SomeType, inp ...[]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointerS(cha, out, inp...)
	return cha
}

func joinPointerSSlice(done chan<- struct{}, out chan<- []*SomeType, inp ...[][]*SomeType) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinPointerSSlice
func JoinPointerSSlice(out chan<- []*SomeType, inp ...[][]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointerSSlice(cha, out, inp...)
	return cha
}

func joinPointerSChan(done chan<- struct{}, out chan<- []*SomeType, inp <-chan []*SomeType) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPointerSChan
func JoinPointerSChan(out chan<- []*SomeType, inp <-chan []*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPointerSChan(cha, out, inp)
	return cha
}

func doitPointerS(done chan<- struct{}, inp <-chan []*SomeType) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePointerS returns a channel to receive one signal before close after inp has been drained.
func DonePointerS(inp <-chan []*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPointerS(cha, inp)
	return cha
}

func doitPointerSSlice(done chan<- ([][]*SomeType), inp <-chan []*SomeType) {
	defer close(done)
	PointerSS := [][]*SomeType{}
	for i := range inp {
		PointerSS = append(PointerSS, i)
	}
	done <- PointerSS
}

// DonePointerSSlice returns a channel which will receive a slice
// of all the PointerSs received on inp channel before close.
// Unlike DonePointerS, a full slice is sent once, not just an event.
func DonePointerSSlice(inp <-chan []*SomeType) (done <-chan ([][]*SomeType)) {
	cha := make(chan ([][]*SomeType))
	go doitPointerSSlice(cha, inp)
	return cha
}

func doitPointerSFunc(done chan<- struct{}, inp <-chan []*SomeType, act func(a []*SomeType)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePointerSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePointerSFunc(inp <-chan []*SomeType, act func(a []*SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []*SomeType) { return }
	}
	go doitPointerSFunc(cha, inp, act)
	return cha
}

func pipePointerSBuffer(out chan<- []*SomeType, inp <-chan []*SomeType) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePointerSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePointerSBuffer(inp <-chan []*SomeType, cap int) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType, cap)
	go pipePointerSBuffer(cha, inp)
	return cha
}

func pipePointerSFunc(out chan<- []*SomeType, inp <-chan []*SomeType, act func(a []*SomeType) []*SomeType) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePointerSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePointerSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePointerSFunc(inp <-chan []*SomeType, act func(a []*SomeType) []*SomeType) (out <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	if act == nil {
		act = func(a []*SomeType) []*SomeType { return a }
	}
	go pipePointerSFunc(cha, inp, act)
	return cha
}

func pipePointerSFork(out1, out2 chan<- []*SomeType, inp <-chan []*SomeType) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePointerSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePointerSFork(inp <-chan []*SomeType) (out1, out2 <-chan []*SomeType) {
	cha1 := make(chan []*SomeType)
	cha2 := make(chan []*SomeType)
	go pipePointerSFork(cha1, cha2, inp)
	return cha1, cha2
}
