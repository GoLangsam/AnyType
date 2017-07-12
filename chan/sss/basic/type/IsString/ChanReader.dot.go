// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// MakeReaderChan returns a new open channel
// (simply a 'chan *strings.Reader' that is).
//
// Note: No 'Reader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderPipelineStartsHere := MakeReaderChan()
//	// ... lot's of code to design and build Your favourite "myReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderPipelineStartsHere <- drop
//	}
//	close(myReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderBuffer) the channel is unbuffered.
//
func MakeReaderChan() (out chan *strings.Reader) {
	return make(chan *strings.Reader)
}

// ChanReader returns a channel to receive all inputs before close.
func ChanReader(inp ...*strings.Reader) (out <-chan *strings.Reader) {
	cha := make(chan *strings.Reader)
	go func(out chan<- *strings.Reader, inp ...*strings.Reader) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanReaderSlice returns a channel to receive all inputs before close.
func ChanReaderSlice(inp ...[]*strings.Reader) (out <-chan *strings.Reader) {
	cha := make(chan *strings.Reader)
	go func(out chan<- *strings.Reader, inp ...[]*strings.Reader) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinReader
func JoinReader(out chan<- *strings.Reader, inp ...*strings.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Reader, inp ...*strings.Reader) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReaderSlice
func JoinReaderSlice(out chan<- *strings.Reader, inp ...[]*strings.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Reader, inp ...[]*strings.Reader) {
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

// JoinReaderChan
func JoinReaderChan(out chan<- *strings.Reader, inp <-chan *strings.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *strings.Reader, inp <-chan *strings.Reader) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReader returns a channel to receive one signal before close after inp has been drained.
func DoneReader(inp <-chan *strings.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *strings.Reader) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReaderSlice returns a channel which will receive a slice
// of all the Readers received on inp channel before close.
// Unlike DoneReader, a full slice is sent once, not just an event.
func DoneReaderSlice(inp <-chan *strings.Reader) (done <-chan []*strings.Reader) {
	cha := make(chan []*strings.Reader)
	go func(inp <-chan *strings.Reader, done chan<- []*strings.Reader) {
		defer close(done)
		ReaderS := []*strings.Reader{}
		for i := range inp {
			ReaderS = append(ReaderS, i)
		}
		done <- ReaderS
	}(inp, cha)
	return cha
}

// DoneReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderFunc(inp <-chan *strings.Reader, act func(a *strings.Reader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *strings.Reader) { return }
	}
	go func(done chan<- struct{}, inp <-chan *strings.Reader, act func(a *strings.Reader)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderBuffer(inp <-chan *strings.Reader, cap int) (out <-chan *strings.Reader) {
	cha := make(chan *strings.Reader, cap)
	go func(out chan<- *strings.Reader, inp <-chan *strings.Reader) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderFunc(inp <-chan *strings.Reader, act func(a *strings.Reader) *strings.Reader) (out <-chan *strings.Reader) {
	cha := make(chan *strings.Reader)
	if act == nil {
		act = func(a *strings.Reader) *strings.Reader { return a }
	}
	go func(out chan<- *strings.Reader, inp <-chan *strings.Reader, act func(a *strings.Reader) *strings.Reader) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderFork(inp <-chan *strings.Reader) (out1, out2 <-chan *strings.Reader) {
	cha1 := make(chan *strings.Reader)
	cha2 := make(chan *strings.Reader)
	go func(out1, out2 chan<- *strings.Reader, inp <-chan *strings.Reader) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ReaderTube is the signature for a pipe function.
type ReaderTube func(inp <-chan *strings.Reader, out <-chan *strings.Reader)

// Readerdaisy returns a channel to receive all inp after having passed thru tube.
func Readerdaisy(inp <-chan *strings.Reader, tube ReaderTube) (out <-chan *strings.Reader) {
	cha := make(chan *strings.Reader)
	go tube(inp, cha)
	return cha
}

// ReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReaderDaisyChain(inp <-chan *strings.Reader, tubes ...ReaderTube) (out <-chan *strings.Reader) {
	cha := inp
	for _, tube := range tubes {
		cha = Readerdaisy(cha, tube)
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
