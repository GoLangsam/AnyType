// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt8Chan returns a new open channel
// (simply a 'chan int8' that is).
//
// Note: No 'Int8-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt8PipelineStartsHere := MakeInt8Chan()
//	// ... lot's of code to design and build Your favourite "myInt8WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt8PipelineStartsHere <- drop
//	}
//	close(myInt8PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt8Buffer) the channel is unbuffered.
//
func MakeInt8Chan() (out chan int8) {
	return make(chan int8)
}

func sendInt8(out chan<- int8, inp ...int8) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanInt8 returns a channel to receive all inputs before close.
func ChanInt8(inp ...int8) (out <-chan int8) {
	cha := make(chan int8)
	go sendInt8(cha, inp...)
	return cha
}

func sendInt8Slice(out chan<- int8, inp ...[]int8) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanInt8Slice returns a channel to receive all inputs before close.
func ChanInt8Slice(inp ...[]int8) (out <-chan int8) {
	cha := make(chan int8)
	go sendInt8Slice(cha, inp...)
	return cha
}

func joinInt8(done chan<- struct{}, out chan<- int8, inp ...int8) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt8
func JoinInt8(out chan<- int8, inp ...int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt8(cha, out, inp...)
	return cha
}

func joinInt8Slice(done chan<- struct{}, out chan<- int8, inp ...[]int8) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinInt8Slice
func JoinInt8Slice(out chan<- int8, inp ...[]int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt8Slice(cha, out, inp...)
	return cha
}

func joinInt8Chan(done chan<- struct{}, out chan<- int8, inp <-chan int8) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt8Chan
func JoinInt8Chan(out chan<- int8, inp <-chan int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt8Chan(cha, out, inp)
	return cha
}

func doitInt8(done chan<- struct{}, inp <-chan int8) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneInt8 returns a channel to receive one signal before close after inp has been drained.
func DoneInt8(inp <-chan int8) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitInt8(cha, inp)
	return cha
}

func doitInt8Slice(done chan<- ([]int8), inp <-chan int8) {
	defer close(done)
	Int8S := []int8{}
	for i := range inp {
		Int8S = append(Int8S, i)
	}
	done <- Int8S
}

// DoneInt8Slice returns a channel which will receive a slice
// of all the Int8s received on inp channel before close.
// Unlike DoneInt8, a full slice is sent once, not just an event.
func DoneInt8Slice(inp <-chan int8) (done <-chan ([]int8)) {
	cha := make(chan ([]int8))
	go doitInt8Slice(cha, inp)
	return cha
}

func doitInt8Func(done chan<- struct{}, inp <-chan int8, act func(a int8)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneInt8Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt8Func(inp <-chan int8, act func(a int8)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int8) { return }
	}
	go doitInt8Func(cha, inp, act)
	return cha
}

func pipeInt8Buffer(out chan<- int8, inp <-chan int8) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeInt8Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt8Buffer(inp <-chan int8, cap int) (out <-chan int8) {
	cha := make(chan int8, cap)
	go pipeInt8Buffer(cha, inp)
	return cha
}

func pipeInt8Func(out chan<- int8, inp <-chan int8, act func(a int8) int8) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeInt8Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt8Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt8Func(inp <-chan int8, act func(a int8) int8) (out <-chan int8) {
	cha := make(chan int8)
	if act == nil {
		act = func(a int8) int8 { return a }
	}
	go pipeInt8Func(cha, inp, act)
	return cha
}

func pipeInt8Fork(out1, out2 chan<- int8, inp <-chan int8) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeInt8Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt8Fork(inp <-chan int8) (out1, out2 <-chan int8) {
	cha1 := make(chan int8)
	cha2 := make(chan int8)
	go pipeInt8Fork(cha1, cha2, inp)
	return cha1, cha2
}