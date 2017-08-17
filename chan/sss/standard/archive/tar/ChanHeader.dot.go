// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

// MakeHeaderChan returns a new open channel
// (simply a 'chan *tar.Header' that is).
//
// Note: No 'Header-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myHeaderPipelineStartsHere := MakeHeaderChan()
//	// ... lot's of code to design and build Your favourite "myHeaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myHeaderPipelineStartsHere <- drop
//	}
//	close(myHeaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeHeaderBuffer) the channel is unbuffered.
//
func MakeHeaderChan() (out chan *tar.Header) {
	return make(chan *tar.Header)
}

// ChanHeader returns a channel to receive all inputs before close.
func ChanHeader(inp ...*tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go func(out chan<- *tar.Header, inp ...*tar.Header) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanHeaderSlice returns a channel to receive all inputs before close.
func ChanHeaderSlice(inp ...[]*tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go func(out chan<- *tar.Header, inp ...[]*tar.Header) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinHeader
func JoinHeader(out chan<- *tar.Header, inp ...*tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Header, inp ...*tar.Header) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinHeaderSlice
func JoinHeaderSlice(out chan<- *tar.Header, inp ...[]*tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Header, inp ...[]*tar.Header) {
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

// JoinHeaderChan
func JoinHeaderChan(out chan<- *tar.Header, inp <-chan *tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Header, inp <-chan *tar.Header) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneHeader returns a channel to receive one signal before close after inp has been drained.
func DoneHeader(inp <-chan *tar.Header) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *tar.Header) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneHeaderSlice returns a channel which will receive a slice
// of all the Headers received on inp channel before close.
// Unlike DoneHeader, a full slice is sent once, not just an event.
func DoneHeaderSlice(inp <-chan *tar.Header) (done <-chan []*tar.Header) {
	cha := make(chan []*tar.Header)
	go func(inp <-chan *tar.Header, done chan<- []*tar.Header) {
		defer close(done)
		HeaderS := []*tar.Header{}
		for i := range inp {
			HeaderS = append(HeaderS, i)
		}
		done <- HeaderS
	}(inp, cha)
	return cha
}

// DoneHeaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *tar.Header) { return }
	}
	go func(done chan<- struct{}, inp <-chan *tar.Header, act func(a *tar.Header)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeHeaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeHeaderBuffer(inp <-chan *tar.Header, cap int) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header, cap)
	go func(out chan<- *tar.Header, inp <-chan *tar.Header) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeHeaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeHeaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeHeaderFunc(inp <-chan *tar.Header, act func(a *tar.Header) *tar.Header) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	if act == nil {
		act = func(a *tar.Header) *tar.Header { return a }
	}
	go func(out chan<- *tar.Header, inp <-chan *tar.Header, act func(a *tar.Header) *tar.Header) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeHeaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeHeaderFork(inp <-chan *tar.Header) (out1, out2 <-chan *tar.Header) {
	cha1 := make(chan *tar.Header)
	cha2 := make(chan *tar.Header)
	go func(out1, out2 chan<- *tar.Header, inp <-chan *tar.Header) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// HeaderTube is the signature for a pipe function.
type HeaderTube func(inp <-chan *tar.Header, out <-chan *tar.Header)

// Headerdaisy returns a channel to receive all inp after having passed thru tube.
func Headerdaisy(inp <-chan *tar.Header, tube HeaderTube) (out <-chan *tar.Header) {
	cha := make(chan *tar.Header)
	go tube(inp, cha)
	return cha
}

// HeaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func HeaderDaisyChain(inp <-chan *tar.Header, tubes ...HeaderTube) (out <-chan *tar.Header) {
	cha := inp
	for _, tube := range tubes {
		cha = Headerdaisy(cha, tube)
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
