// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeErrorSChan returns a new open channel
// (simply a 'chan []error' that is).
//
// Note: No 'ErrorS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myErrorSPipelineStartsHere := MakeErrorSChan()
//	// ... lot's of code to design and build Your favourite "myErrorSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myErrorSPipelineStartsHere <- drop
//	}
//	close(myErrorSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeErrorSBuffer) the channel is unbuffered.
//
func MakeErrorSChan() (out chan []error) {
	return make(chan []error)
}

func sendErrorS(out chan<- []error, inp ...[]error) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanErrorS returns a channel to receive all inputs before close.
func ChanErrorS(inp ...[]error) (out <-chan []error) {
	cha := make(chan []error)
	go sendErrorS(cha, inp...)
	return cha
}

func sendErrorSSlice(out chan<- []error, inp ...[][]error) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanErrorSSlice returns a channel to receive all inputs before close.
func ChanErrorSSlice(inp ...[][]error) (out <-chan []error) {
	cha := make(chan []error)
	go sendErrorSSlice(cha, inp...)
	return cha
}

func joinErrorS(done chan<- struct{}, out chan<- []error, inp ...[]error) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinErrorS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorS(out chan<- []error, inp ...[]error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinErrorS(cha, out, inp...)
	return cha
}

func joinErrorSSlice(done chan<- struct{}, out chan<- []error, inp ...[][]error) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinErrorSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorSSlice(out chan<- []error, inp ...[][]error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinErrorSSlice(cha, out, inp...)
	return cha
}

func joinErrorSChan(done chan<- struct{}, out chan<- []error, inp <-chan []error) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinErrorSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinErrorSChan(out chan<- []error, inp <-chan []error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinErrorSChan(cha, out, inp)
	return cha
}

func doitErrorS(done chan<- struct{}, inp <-chan []error) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneErrorS returns a channel to receive one signal before close after inp has been drained.
func DoneErrorS(inp <-chan []error) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitErrorS(cha, inp)
	return cha
}

func doitErrorSSlice(done chan<- ([][]error), inp <-chan []error) {
	defer close(done)
	ErrorSS := [][]error{}
	for i := range inp {
		ErrorSS = append(ErrorSS, i)
	}
	done <- ErrorSS
}

// DoneErrorSSlice returns a channel which will receive a slice
// of all the ErrorSs received on inp channel before close.
// Unlike DoneErrorS, a full slice is sent once, not just an event.
func DoneErrorSSlice(inp <-chan []error) (done <-chan ([][]error)) {
	cha := make(chan ([][]error))
	go doitErrorSSlice(cha, inp)
	return cha
}

func doitErrorSFunc(done chan<- struct{}, inp <-chan []error, act func(a []error)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneErrorSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneErrorSFunc(inp <-chan []error, act func(a []error)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []error) { return }
	}
	go doitErrorSFunc(cha, inp, act)
	return cha
}

func pipeErrorSBuffer(out chan<- []error, inp <-chan []error) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeErrorSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeErrorSBuffer(inp <-chan []error, cap int) (out <-chan []error) {
	cha := make(chan []error, cap)
	go pipeErrorSBuffer(cha, inp)
	return cha
}

func pipeErrorSFunc(out chan<- []error, inp <-chan []error, act func(a []error) []error) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeErrorSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeErrorSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeErrorSFunc(inp <-chan []error, act func(a []error) []error) (out <-chan []error) {
	cha := make(chan []error)
	if act == nil {
		act = func(a []error) []error { return a }
	}
	go pipeErrorSFunc(cha, inp, act)
	return cha
}

func pipeErrorSFork(out1, out2 chan<- []error, inp <-chan []error) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeErrorSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeErrorSFork(inp <-chan []error) (out1, out2 <-chan []error) {
	cha1 := make(chan []error)
	cha2 := make(chan []error)
	go pipeErrorSFork(cha1, cha2, inp)
	return cha1, cha2
}
