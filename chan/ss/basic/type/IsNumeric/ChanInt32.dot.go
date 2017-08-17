// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt32Chan returns a new open channel
// (simply a 'chan int32' that is).
//
// Note: No 'Int32-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt32PipelineStartsHere := MakeInt32Chan()
//	// ... lot's of code to design and build Your favourite "myInt32WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt32PipelineStartsHere <- drop
//	}
//	close(myInt32PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt32Buffer) the channel is unbuffered.
//
func MakeInt32Chan() (out chan int32) {
	return make(chan int32)
}

func sendInt32(out chan<- int32, inp ...int32) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanInt32 returns a channel to receive all inputs before close.
func ChanInt32(inp ...int32) (out <-chan int32) {
	cha := make(chan int32)
	go sendInt32(cha, inp...)
	return cha
}

func sendInt32Slice(out chan<- int32, inp ...[]int32) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanInt32Slice returns a channel to receive all inputs before close.
func ChanInt32Slice(inp ...[]int32) (out <-chan int32) {
	cha := make(chan int32)
	go sendInt32Slice(cha, inp...)
	return cha
}

func joinInt32(done chan<- struct{}, out chan<- int32, inp ...int32) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt32
func JoinInt32(out chan<- int32, inp ...int32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt32(cha, out, inp...)
	return cha
}

func joinInt32Slice(done chan<- struct{}, out chan<- int32, inp ...[]int32) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinInt32Slice
func JoinInt32Slice(out chan<- int32, inp ...[]int32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt32Slice(cha, out, inp...)
	return cha
}

func joinInt32Chan(done chan<- struct{}, out chan<- int32, inp <-chan int32) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt32Chan
func JoinInt32Chan(out chan<- int32, inp <-chan int32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt32Chan(cha, out, inp)
	return cha
}

func doitInt32(done chan<- struct{}, inp <-chan int32) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneInt32 returns a channel to receive one signal before close after inp has been drained.
func DoneInt32(inp <-chan int32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitInt32(cha, inp)
	return cha
}

func doitInt32Slice(done chan<- ([]int32), inp <-chan int32) {
	defer close(done)
	Int32S := []int32{}
	for i := range inp {
		Int32S = append(Int32S, i)
	}
	done <- Int32S
}

// DoneInt32Slice returns a channel which will receive a slice
// of all the Int32s received on inp channel before close.
// Unlike DoneInt32, a full slice is sent once, not just an event.
func DoneInt32Slice(inp <-chan int32) (done <-chan ([]int32)) {
	cha := make(chan ([]int32))
	go doitInt32Slice(cha, inp)
	return cha
}

func doitInt32Func(done chan<- struct{}, inp <-chan int32, act func(a int32)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneInt32Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt32Func(inp <-chan int32, act func(a int32)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int32) { return }
	}
	go doitInt32Func(cha, inp, act)
	return cha
}

func pipeInt32Buffer(out chan<- int32, inp <-chan int32) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeInt32Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt32Buffer(inp <-chan int32, cap int) (out <-chan int32) {
	cha := make(chan int32, cap)
	go pipeInt32Buffer(cha, inp)
	return cha
}

func pipeInt32Func(out chan<- int32, inp <-chan int32, act func(a int32) int32) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeInt32Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt32Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt32Func(inp <-chan int32, act func(a int32) int32) (out <-chan int32) {
	cha := make(chan int32)
	if act == nil {
		act = func(a int32) int32 { return a }
	}
	go pipeInt32Func(cha, inp, act)
	return cha
}

func pipeInt32Fork(out1, out2 chan<- int32, inp <-chan int32) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeInt32Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt32Fork(inp <-chan int32) (out1, out2 <-chan int32) {
	cha1 := make(chan int32)
	cha2 := make(chan int32)
	go pipeInt32Fork(cha1, cha2, inp)
	return cha1, cha2
}
