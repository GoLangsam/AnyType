// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsUnsafe

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeChan returns a new open channel
// (simply a 'chan uintptr' that is).
//
// Note: No '-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipelineStartsHere := MakeChan()
//	// ... lot's of code to design and build Your favourite "myWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipelineStartsHere <- drop
//	}
//	close(myPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBuffer) the channel is unbuffered.
//
func MakeChan() (out chan uintptr) {
	return make(chan uintptr)
}

func send(out chan<- uintptr, inp ...uintptr) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...uintptr) (out <-chan uintptr) {
	cha := make(chan uintptr)
	go send(cha, inp...)
	return cha
}

func sendSlice(out chan<- uintptr, inp ...[]uintptr) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]uintptr) (out <-chan uintptr) {
	cha := make(chan uintptr)
	go sendSlice(cha, inp...)
	return cha
}

func join(done chan<- struct{}, out chan<- uintptr, inp ...uintptr) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// Join sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func Join(out chan<- uintptr, inp ...uintptr) (done <-chan struct{}) {
	cha := make(chan struct{})
	go join(cha, out, inp...)
	return cha
}

func joinSlice(done chan<- struct{}, out chan<- uintptr, inp ...[]uintptr) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSlice(out chan<- uintptr, inp ...[]uintptr) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSlice(cha, out, inp...)
	return cha
}

func joinChan(done chan<- struct{}, out chan<- uintptr, inp <-chan uintptr) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinChan(out chan<- uintptr, inp <-chan uintptr) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinChan(cha, out, inp)
	return cha
}

func doit(done chan<- struct{}, inp <-chan uintptr) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan uintptr) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doit(cha, inp)
	return cha
}

func doitSlice(done chan<- ([]uintptr), inp <-chan uintptr) {
	defer close(done)
	S := []uintptr{}
	for i := range inp {
		S = append(S, i)
	}
	done <- S
}

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan uintptr) (done <-chan ([]uintptr)) {
	cha := make(chan ([]uintptr))
	go doitSlice(cha, inp)
	return cha
}

func doitFunc(done chan<- struct{}, inp <-chan uintptr, act func(a uintptr)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan uintptr, act func(a uintptr)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uintptr) { return }
	}
	go doitFunc(cha, inp, act)
	return cha
}

func pipeBuffer(out chan<- uintptr, inp <-chan uintptr) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan uintptr, cap int) (out <-chan uintptr) {
	cha := make(chan uintptr, cap)
	go pipeBuffer(cha, inp)
	return cha
}

func pipeFunc(out chan<- uintptr, inp <-chan uintptr, act func(a uintptr) uintptr) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan uintptr, act func(a uintptr) uintptr) (out <-chan uintptr) {
	cha := make(chan uintptr)
	if act == nil {
		act = func(a uintptr) uintptr { return a }
	}
	go pipeFunc(cha, inp, act)
	return cha
}

func pipeFork(out1, out2 chan<- uintptr, inp <-chan uintptr) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan uintptr) (out1, out2 <-chan uintptr) {
	cha1 := make(chan uintptr)
	cha2 := make(chan uintptr)
	go pipeFork(cha1, cha2, inp)
	return cha1, cha2
}
