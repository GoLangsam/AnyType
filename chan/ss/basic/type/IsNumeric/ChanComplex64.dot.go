// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeComplex64Chan returns a new open channel
// (simply a 'chan complex64' that is).
//
// Note: No 'Complex64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myComplex64PipelineStartsHere := MakeComplex64Chan()
//	// ... lot's of code to design and build Your favourite "myComplex64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myComplex64PipelineStartsHere <- drop
//	}
//	close(myComplex64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeComplex64Buffer) the channel is unbuffered.
//
func MakeComplex64Chan() (out chan complex64) {
	return make(chan complex64)
}

func sendComplex64(out chan<- complex64, inp ...complex64) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanComplex64 returns a channel to receive all inputs before close.
func ChanComplex64(inp ...complex64) (out <-chan complex64) {
	cha := make(chan complex64)
	go sendComplex64(cha, inp...)
	return cha
}

func sendComplex64Slice(out chan<- complex64, inp ...[]complex64) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanComplex64Slice returns a channel to receive all inputs before close.
func ChanComplex64Slice(inp ...[]complex64) (out <-chan complex64) {
	cha := make(chan complex64)
	go sendComplex64Slice(cha, inp...)
	return cha
}

func joinComplex64(done chan<- struct{}, out chan<- complex64, inp ...complex64) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinComplex64
func JoinComplex64(out chan<- complex64, inp ...complex64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex64(cha, out, inp...)
	return cha
}

func joinComplex64Slice(done chan<- struct{}, out chan<- complex64, inp ...[]complex64) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinComplex64Slice
func JoinComplex64Slice(out chan<- complex64, inp ...[]complex64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex64Slice(cha, out, inp...)
	return cha
}

func joinComplex64Chan(done chan<- struct{}, out chan<- complex64, inp <-chan complex64) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinComplex64Chan
func JoinComplex64Chan(out chan<- complex64, inp <-chan complex64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinComplex64Chan(cha, out, inp)
	return cha
}

func doitComplex64(done chan<- struct{}, inp <-chan complex64) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneComplex64 returns a channel to receive one signal before close after inp has been drained.
func DoneComplex64(inp <-chan complex64) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitComplex64(cha, inp)
	return cha
}

func doitComplex64Slice(done chan<- ([]complex64), inp <-chan complex64) {
	defer close(done)
	Complex64S := []complex64{}
	for i := range inp {
		Complex64S = append(Complex64S, i)
	}
	done <- Complex64S
}

// DoneComplex64Slice returns a channel which will receive a slice
// of all the Complex64s received on inp channel before close.
// Unlike DoneComplex64, a full slice is sent once, not just an event.
func DoneComplex64Slice(inp <-chan complex64) (done <-chan ([]complex64)) {
	cha := make(chan ([]complex64))
	go doitComplex64Slice(cha, inp)
	return cha
}

func doitComplex64Func(done chan<- struct{}, inp <-chan complex64, act func(a complex64)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneComplex64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneComplex64Func(inp <-chan complex64, act func(a complex64)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a complex64) { return }
	}
	go doitComplex64Func(cha, inp, act)
	return cha
}

func pipeComplex64Buffer(out chan<- complex64, inp <-chan complex64) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeComplex64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeComplex64Buffer(inp <-chan complex64, cap int) (out <-chan complex64) {
	cha := make(chan complex64, cap)
	go pipeComplex64Buffer(cha, inp)
	return cha
}

func pipeComplex64Func(out chan<- complex64, inp <-chan complex64, act func(a complex64) complex64) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeComplex64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeComplex64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeComplex64Func(inp <-chan complex64, act func(a complex64) complex64) (out <-chan complex64) {
	cha := make(chan complex64)
	if act == nil {
		act = func(a complex64) complex64 { return a }
	}
	go pipeComplex64Func(cha, inp, act)
	return cha
}

func pipeComplex64Fork(out1, out2 chan<- complex64, inp <-chan complex64) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeComplex64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeComplex64Fork(inp <-chan complex64) (out1, out2 <-chan complex64) {
	cha1 := make(chan complex64)
	cha2 := make(chan complex64)
	go pipeComplex64Fork(cha1, cha2, inp)
	return cha1, cha2
}
