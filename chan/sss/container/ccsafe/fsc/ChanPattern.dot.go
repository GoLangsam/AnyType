// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/fs"
)

// MakePatternChan returns a new open channel
// (simply a 'chan *fs.Pattern' that is).
//
// Note: No 'Pattern-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatternPipelineStartsHere := MakePatternChan()
//	// ... lot's of code to design and build Your favourite "myPatternWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatternPipelineStartsHere <- drop
//	}
//	close(myPatternPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatternBuffer) the channel is unbuffered.
//
func MakePatternChan() (out chan *fs.Pattern) {
	return make(chan *fs.Pattern)
}

// ChanPattern returns a channel to receive all inputs before close.
func ChanPattern(inp ...*fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	go func(out chan<- *fs.Pattern, inp ...*fs.Pattern) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanPatternSlice returns a channel to receive all inputs before close.
func ChanPatternSlice(inp ...[]*fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	go func(out chan<- *fs.Pattern, inp ...[]*fs.Pattern) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinPattern
func JoinPattern(out chan<- *fs.Pattern, inp ...*fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.Pattern, inp ...*fs.Pattern) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinPatternSlice
func JoinPatternSlice(out chan<- *fs.Pattern, inp ...[]*fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.Pattern, inp ...[]*fs.Pattern) {
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

// JoinPatternChan
func JoinPatternChan(out chan<- *fs.Pattern, inp <-chan *fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DonePattern returns a channel to receive one signal before close after inp has been drained.
func DonePattern(inp <-chan *fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *fs.Pattern) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DonePatternSlice returns a channel which will receive a slice
// of all the Patterns received on inp channel before close.
// Unlike DonePattern, a full slice is sent once, not just an event.
func DonePatternSlice(inp <-chan *fs.Pattern) (done <-chan []*fs.Pattern) {
	cha := make(chan []*fs.Pattern)
	go func(inp <-chan *fs.Pattern, done chan<- []*fs.Pattern) {
		defer close(done)
		PatternS := []*fs.Pattern{}
		for i := range inp {
			PatternS = append(PatternS, i)
		}
		done <- PatternS
	}(inp, cha)
	return cha
}

// DonePatternFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternFunc(inp <-chan *fs.Pattern, act func(a *fs.Pattern)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.Pattern) { return }
	}
	go func(done chan<- struct{}, inp <-chan *fs.Pattern, act func(a *fs.Pattern)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipePatternBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternBuffer(inp <-chan *fs.Pattern, cap int) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern, cap)
	go func(out chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipePatternFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternFunc(inp <-chan *fs.Pattern, act func(a *fs.Pattern) *fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	if act == nil {
		act = func(a *fs.Pattern) *fs.Pattern { return a }
	}
	go func(out chan<- *fs.Pattern, inp <-chan *fs.Pattern, act func(a *fs.Pattern) *fs.Pattern) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipePatternFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternFork(inp <-chan *fs.Pattern) (out1, out2 <-chan *fs.Pattern) {
	cha1 := make(chan *fs.Pattern)
	cha2 := make(chan *fs.Pattern)
	go func(out1, out2 chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// PatternTube is the signature for a pipe function.
type PatternTube func(inp <-chan *fs.Pattern, out <-chan *fs.Pattern)

// Patterndaisy returns a channel to receive all inp after having passed thru tube.
func Patterndaisy(inp <-chan *fs.Pattern, tube PatternTube) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	go tube(inp, cha)
	return cha
}

// PatternDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PatternDaisyChain(inp <-chan *fs.Pattern, tubes ...PatternTube) (out <-chan *fs.Pattern) {
	cha := inp
	for _, tube := range tubes {
		cha = Patterndaisy(cha, tube)
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
