// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeSomeTypeSChan returns a new open channel
// (simply a 'chan []SomeType' that is).
//
// Note: No 'SomeTypeS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySomeTypeSPipelineStartsHere := MakeSomeTypeSChan()
//	// ... lot's of code to design and build Your favourite "mySomeTypeSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySomeTypeSPipelineStartsHere <- drop
//	}
//	close(mySomeTypeSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSomeTypeSBuffer) the channel is unbuffered.
//
func MakeSomeTypeSChan() (out chan []SomeType) {
	return make(chan []SomeType)
}

func sendSomeTypeS(out chan<- []SomeType, inp ...[]SomeType) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanSomeTypeS returns a channel to receive all inputs before close.
func ChanSomeTypeS(inp ...[]SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	go sendSomeTypeS(cha, inp...)
	return cha
}

func sendSomeTypeSSlice(out chan<- []SomeType, inp ...[][]SomeType) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanSomeTypeSSlice returns a channel to receive all inputs before close.
func ChanSomeTypeSSlice(inp ...[][]SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	go sendSomeTypeSSlice(cha, inp...)
	return cha
}

func joinSomeTypeS(done chan<- struct{}, out chan<- []SomeType, inp ...[]SomeType) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSomeTypeS
func JoinSomeTypeS(out chan<- []SomeType, inp ...[]SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeTypeS(cha, out, inp...)
	return cha
}

func joinSomeTypeSSlice(done chan<- struct{}, out chan<- []SomeType, inp ...[][]SomeType) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinSomeTypeSSlice
func JoinSomeTypeSSlice(out chan<- []SomeType, inp ...[][]SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeTypeSSlice(cha, out, inp...)
	return cha
}

func joinSomeTypeSChan(done chan<- struct{}, out chan<- []SomeType, inp <-chan []SomeType) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSomeTypeSChan
func JoinSomeTypeSChan(out chan<- []SomeType, inp <-chan []SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSomeTypeSChan(cha, out, inp)
	return cha
}

func doitSomeTypeS(done chan<- struct{}, inp <-chan []SomeType) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneSomeTypeS returns a channel to receive one signal before close after inp has been drained.
func DoneSomeTypeS(inp <-chan []SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitSomeTypeS(cha, inp)
	return cha
}

func doitSomeTypeSSlice(done chan<- ([][]SomeType), inp <-chan []SomeType) {
	defer close(done)
	SomeTypeSS := [][]SomeType{}
	for i := range inp {
		SomeTypeSS = append(SomeTypeSS, i)
	}
	done <- SomeTypeSS
}

// DoneSomeTypeSSlice returns a channel which will receive a slice
// of all the SomeTypeSs received on inp channel before close.
// Unlike DoneSomeTypeS, a full slice is sent once, not just an event.
func DoneSomeTypeSSlice(inp <-chan []SomeType) (done <-chan ([][]SomeType)) {
	cha := make(chan ([][]SomeType))
	go doitSomeTypeSSlice(cha, inp)
	return cha
}

func doitSomeTypeSFunc(done chan<- struct{}, inp <-chan []SomeType, act func(a []SomeType)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneSomeTypeSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSomeTypeSFunc(inp <-chan []SomeType, act func(a []SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []SomeType) { return }
	}
	go doitSomeTypeSFunc(cha, inp, act)
	return cha
}

func pipeSomeTypeSBuffer(out chan<- []SomeType, inp <-chan []SomeType) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeSomeTypeSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSomeTypeSBuffer(inp <-chan []SomeType, cap int) (out <-chan []SomeType) {
	cha := make(chan []SomeType, cap)
	go pipeSomeTypeSBuffer(cha, inp)
	return cha
}

func pipeSomeTypeSFunc(out chan<- []SomeType, inp <-chan []SomeType, act func(a []SomeType) []SomeType) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeSomeTypeSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSomeTypeSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSomeTypeSFunc(inp <-chan []SomeType, act func(a []SomeType) []SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	if act == nil {
		act = func(a []SomeType) []SomeType { return a }
	}
	go pipeSomeTypeSFunc(cha, inp, act)
	return cha
}

func pipeSomeTypeSFork(out1, out2 chan<- []SomeType, inp <-chan []SomeType) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeSomeTypeSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSomeTypeSFork(inp <-chan []SomeType) (out1, out2 <-chan []SomeType) {
	cha1 := make(chan []SomeType)
	cha2 := make(chan []SomeType)
	go pipeSomeTypeSFork(cha1, cha2, inp)
	return cha1, cha2
}