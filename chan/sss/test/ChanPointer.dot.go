// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakePointerChan returns a new open channel
// (simply a 'chan *SomeType' that is).
//
// Note: No 'Pointer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPointerPipelineStartsHere := MakePointerChan()
//	// ... lot's of code to design and build Your favourite "myPointerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPointerPipelineStartsHere <- drop
//	}
//	close(myPointerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePointerBuffer) the channel is unbuffered.
//
func MakePointerChan() (out chan *SomeType) {
	return make(chan *SomeType)
}

// ChanPointer returns a channel to receive all inputs before close.
func ChanPointer(inp ...*SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go func(out chan<- *SomeType, inp ...*SomeType) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanPointerSlice returns a channel to receive all inputs before close.
func ChanPointerSlice(inp ...[]*SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	go func(out chan<- *SomeType, inp ...[]*SomeType) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinPointer
func JoinPointer(out chan<- *SomeType, inp ...*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *SomeType, inp ...*SomeType) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinPointerSlice
func JoinPointerSlice(out chan<- *SomeType, inp ...[]*SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *SomeType, inp ...[]*SomeType) {
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

// JoinPointerChan
func JoinPointerChan(out chan<- *SomeType, inp <-chan *SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *SomeType, inp <-chan *SomeType) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DonePointer returns a channel to receive one signal before close after inp has been drained.
func DonePointer(inp <-chan *SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *SomeType) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DonePointerSlice returns a channel which will receive a slice
// of all the Pointers received on inp channel before close.
// Unlike DonePointer, a full slice is sent once, not just an event.
func DonePointerSlice(inp <-chan *SomeType) (done <-chan []*SomeType) {
	cha := make(chan []*SomeType)
	go func(inp <-chan *SomeType, done chan<- []*SomeType) {
		defer close(done)
		PointerS := []*SomeType{}
		for i := range inp {
			PointerS = append(PointerS, i)
		}
		done <- PointerS
	}(inp, cha)
	return cha
}

// DonePointerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePointerFunc(inp <-chan *SomeType, act func(a *SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *SomeType) { return }
	}
	go func(done chan<- struct{}, inp <-chan *SomeType, act func(a *SomeType)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipePointerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePointerBuffer(inp <-chan *SomeType, cap int) (out <-chan *SomeType) {
	cha := make(chan *SomeType, cap)
	go func(out chan<- *SomeType, inp <-chan *SomeType) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipePointerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePointerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePointerFunc(inp <-chan *SomeType, act func(a *SomeType) *SomeType) (out <-chan *SomeType) {
	cha := make(chan *SomeType)
	if act == nil {
		act = func(a *SomeType) *SomeType { return a }
	}
	go func(out chan<- *SomeType, inp <-chan *SomeType, act func(a *SomeType) *SomeType) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipePointerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePointerFork(inp <-chan *SomeType) (out1, out2 <-chan *SomeType) {
	cha1 := make(chan *SomeType)
	cha2 := make(chan *SomeType)
	go func(out1, out2 chan<- *SomeType, inp <-chan *SomeType) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
