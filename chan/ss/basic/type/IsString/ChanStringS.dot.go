// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeStringSChan returns a new open channel
// (simply a 'chan []string' that is).
//
// Note: No 'StringS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myStringSPipelineStartsHere := MakeStringSChan()
//	// ... lot's of code to design and build Your favourite "myStringSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myStringSPipelineStartsHere <- drop
//	}
//	close(myStringSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeStringSBuffer) the channel is unbuffered.
//
func MakeStringSChan() (out chan []string) {
	return make(chan []string)
}

func sendStringS(out chan<- []string, inp ...[]string) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanStringS returns a channel to receive all inputs before close.
func ChanStringS(inp ...[]string) (out <-chan []string) {
	cha := make(chan []string)
	go sendStringS(cha, inp...)
	return cha
}

func sendStringSSlice(out chan<- []string, inp ...[][]string) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanStringSSlice returns a channel to receive all inputs before close.
func ChanStringSSlice(inp ...[][]string) (out <-chan []string) {
	cha := make(chan []string)
	go sendStringSSlice(cha, inp...)
	return cha
}

func joinStringS(done chan<- struct{}, out chan<- []string, inp ...[]string) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinStringS
func JoinStringS(out chan<- []string, inp ...[]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinStringS(cha, out, inp...)
	return cha
}

func joinStringSSlice(done chan<- struct{}, out chan<- []string, inp ...[][]string) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinStringSSlice
func JoinStringSSlice(out chan<- []string, inp ...[][]string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinStringSSlice(cha, out, inp...)
	return cha
}

func joinStringSChan(done chan<- struct{}, out chan<- []string, inp <-chan []string) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinStringSChan
func JoinStringSChan(out chan<- []string, inp <-chan []string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinStringSChan(cha, out, inp)
	return cha
}

func doitStringS(done chan<- struct{}, inp <-chan []string) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneStringS returns a channel to receive one signal before close after inp has been drained.
func DoneStringS(inp <-chan []string) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitStringS(cha, inp)
	return cha
}

func doitStringSSlice(done chan<- ([][]string), inp <-chan []string) {
	defer close(done)
	StringSS := [][]string{}
	for i := range inp {
		StringSS = append(StringSS, i)
	}
	done <- StringSS
}

// DoneStringSSlice returns a channel which will receive a slice
// of all the StringSs received on inp channel before close.
// Unlike DoneStringS, a full slice is sent once, not just an event.
func DoneStringSSlice(inp <-chan []string) (done <-chan ([][]string)) {
	cha := make(chan ([][]string))
	go doitStringSSlice(cha, inp)
	return cha
}

func doitStringSFunc(done chan<- struct{}, inp <-chan []string, act func(a []string)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneStringSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneStringSFunc(inp <-chan []string, act func(a []string)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []string) { return }
	}
	go doitStringSFunc(cha, inp, act)
	return cha
}

func pipeStringSBuffer(out chan<- []string, inp <-chan []string) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeStringSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeStringSBuffer(inp <-chan []string, cap int) (out <-chan []string) {
	cha := make(chan []string, cap)
	go pipeStringSBuffer(cha, inp)
	return cha
}

func pipeStringSFunc(out chan<- []string, inp <-chan []string, act func(a []string) []string) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeStringSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeStringSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeStringSFunc(inp <-chan []string, act func(a []string) []string) (out <-chan []string) {
	cha := make(chan []string)
	if act == nil {
		act = func(a []string) []string { return a }
	}
	go pipeStringSFunc(cha, inp, act)
	return cha
}

func pipeStringSFork(out1, out2 chan<- []string, inp <-chan []string) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeStringSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeStringSFork(inp <-chan []string) (out1, out2 <-chan []string) {
	cha1 := make(chan []string)
	cha2 := make(chan []string)
	go pipeStringSFork(cha1, cha2, inp)
	return cha1, cha2
}
