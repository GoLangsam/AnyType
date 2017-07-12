// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/tag"
)

// MakeTagChan returns a new open channel
// (simply a 'chan tag.TagAny' that is).
//
// Note: No 'Tag-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myTagPipelineStartsHere := MakeTagChan()
//	// ... lot's of code to design and build Your favourite "myTagWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myTagPipelineStartsHere <- drop
//	}
//	close(myTagPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeTagBuffer) the channel is unbuffered.
//
func MakeTagChan() (out chan tag.TagAny) {
	return make(chan tag.TagAny)
}

func sendTag(out chan<- tag.TagAny, inp ...tag.TagAny) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanTag returns a channel to receive all inputs before close.
func ChanTag(inp ...tag.TagAny) (out <-chan tag.TagAny) {
	cha := make(chan tag.TagAny)
	go sendTag(cha, inp...)
	return cha
}

func sendTagSlice(out chan<- tag.TagAny, inp ...[]tag.TagAny) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanTagSlice returns a channel to receive all inputs before close.
func ChanTagSlice(inp ...[]tag.TagAny) (out <-chan tag.TagAny) {
	cha := make(chan tag.TagAny)
	go sendTagSlice(cha, inp...)
	return cha
}

func joinTag(done chan<- struct{}, out chan<- tag.TagAny, inp ...tag.TagAny) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinTag
func JoinTag(out chan<- tag.TagAny, inp ...tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinTag(cha, out, inp...)
	return cha
}

func joinTagSlice(done chan<- struct{}, out chan<- tag.TagAny, inp ...[]tag.TagAny) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinTagSlice
func JoinTagSlice(out chan<- tag.TagAny, inp ...[]tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinTagSlice(cha, out, inp...)
	return cha
}

func joinTagChan(done chan<- struct{}, out chan<- tag.TagAny, inp <-chan tag.TagAny) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinTagChan
func JoinTagChan(out chan<- tag.TagAny, inp <-chan tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinTagChan(cha, out, inp)
	return cha
}

func doitTag(done chan<- struct{}, inp <-chan tag.TagAny) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneTag returns a channel to receive one signal before close after inp has been drained.
func DoneTag(inp <-chan tag.TagAny) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitTag(cha, inp)
	return cha
}

func doitTagSlice(done chan<- ([]tag.TagAny), inp <-chan tag.TagAny) {
	defer close(done)
	TagS := []tag.TagAny{}
	for i := range inp {
		TagS = append(TagS, i)
	}
	done <- TagS
}

// DoneTagSlice returns a channel which will receive a slice
// of all the Tags received on inp channel before close.
// Unlike DoneTag, a full slice is sent once, not just an event.
func DoneTagSlice(inp <-chan tag.TagAny) (done <-chan ([]tag.TagAny)) {
	cha := make(chan ([]tag.TagAny))
	go doitTagSlice(cha, inp)
	return cha
}

func doitTagFunc(done chan<- struct{}, inp <-chan tag.TagAny, act func(a tag.TagAny)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneTagFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneTagFunc(inp <-chan tag.TagAny, act func(a tag.TagAny)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a tag.TagAny) { return }
	}
	go doitTagFunc(cha, inp, act)
	return cha
}

func pipeTagBuffer(out chan<- tag.TagAny, inp <-chan tag.TagAny) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeTagBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeTagBuffer(inp <-chan tag.TagAny, cap int) (out <-chan tag.TagAny) {
	cha := make(chan tag.TagAny, cap)
	go pipeTagBuffer(cha, inp)
	return cha
}

func pipeTagFunc(out chan<- tag.TagAny, inp <-chan tag.TagAny, act func(a tag.TagAny) tag.TagAny) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeTagFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeTagMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeTagFunc(inp <-chan tag.TagAny, act func(a tag.TagAny) tag.TagAny) (out <-chan tag.TagAny) {
	cha := make(chan tag.TagAny)
	if act == nil {
		act = func(a tag.TagAny) tag.TagAny { return a }
	}
	go pipeTagFunc(cha, inp, act)
	return cha
}

func pipeTagFork(out1, out2 chan<- tag.TagAny, inp <-chan tag.TagAny) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeTagFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeTagFork(inp <-chan tag.TagAny) (out1, out2 <-chan tag.TagAny) {
	cha1 := make(chan tag.TagAny)
	cha2 := make(chan tag.TagAny)
	go pipeTagFork(cha1, cha2, inp)
	return cha1, cha2
}
