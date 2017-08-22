// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsFoldChan returns a new open channel
// (simply a 'chan *fs.FsFold' that is).
//
// Note: No 'FsFold-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFoldPipelineStartsHere := MakeFsFoldChan()
//	// ... lot's of code to design and build Your favourite "myFsFoldWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFoldPipelineStartsHere <- drop
//	}
//	close(myFsFoldPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFoldBuffer) the channel is unbuffered.
//
func MakeFsFoldChan() (out chan *fs.FsFold) {
	return make(chan *fs.FsFold)
}

// ChanFsFold returns a channel to receive all inputs before close.
func ChanFsFold(inp ...*fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go func(out chan<- *fs.FsFold, inp ...*fs.FsFold) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanFsFoldSlice returns a channel to receive all inputs before close.
func ChanFsFoldSlice(inp ...[]*fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go func(out chan<- *fs.FsFold, inp ...[]*fs.FsFold) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanFsFoldFuncNil returns a channel to receive all results of act until nil before close.
func ChanFsFoldFuncNil(act func() *fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go func(out chan<- *fs.FsFold, act func() *fs.FsFold) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFsFoldFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsFoldFuncNok(act func() (*fs.FsFold, bool)) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go func(out chan<- *fs.FsFold, act func() (*fs.FsFold, bool)) {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFsFoldFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsFoldFuncErr(act func() (*fs.FsFold, error)) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go func(out chan<- *fs.FsFold, act func() (*fs.FsFold, error)) {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// JoinFsFold sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFold(out chan<- *fs.FsFold, inp ...*fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsFold, inp ...*fs.FsFold) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsFoldSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFoldSlice(out chan<- *fs.FsFold, inp ...[]*fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsFold, inp ...[]*fs.FsFold) {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsFoldChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFoldChan(out chan<- *fs.FsFold, inp <-chan *fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFsFold returns a channel to receive one signal before close after inp has been drained.
func DoneFsFold(inp <-chan *fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *fs.FsFold) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFsFoldSlice returns a channel which will receive a slice
// of all the FsFolds received on inp channel before close.
// Unlike DoneFsFold, a full slice is sent once, not just an event.
func DoneFsFoldSlice(inp <-chan *fs.FsFold) (done <-chan []*fs.FsFold) {
	cha := make(chan []*fs.FsFold)
	go func(inp <-chan *fs.FsFold, done chan<- []*fs.FsFold) {
		defer close(done)
		FsFoldS := []*fs.FsFold{}
		for i := range inp {
			FsFoldS = append(FsFoldS, i)
		}
		done <- FsFoldS
	}(inp, cha)
	return cha
}

// DoneFsFoldFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsFold) { return }
	}
	go func(done chan<- struct{}, inp <-chan *fs.FsFold, act func(a *fs.FsFold)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFsFoldBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFoldBuffer(inp <-chan *fs.FsFold, cap int) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold, cap)
	go func(out chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFsFoldFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFoldMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold) *fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	if act == nil {
		act = func(a *fs.FsFold) *fs.FsFold { return a }
	}
	go func(out chan<- *fs.FsFold, inp <-chan *fs.FsFold, act func(a *fs.FsFold) *fs.FsFold) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFsFoldFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFoldFork(inp <-chan *fs.FsFold) (out1, out2 <-chan *fs.FsFold) {
	cha1 := make(chan *fs.FsFold)
	cha2 := make(chan *fs.FsFold)
	go func(out1, out2 chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// FsFoldTube is the signature for a pipe function.
type FsFoldTube func(inp <-chan *fs.FsFold, out <-chan *fs.FsFold)

// FsFoldDaisy returns a channel to receive all inp after having passed thru tube.
func FsFoldDaisy(inp <-chan *fs.FsFold, tube FsFoldTube) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go tube(inp, cha)
	return cha
}

// FsFoldDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsFoldDaisyChain(inp <-chan *fs.FsFold, tubes ...FsFoldTube) (out <-chan *fs.FsFold) {
	cha := inp
	for i := range tubes {
		cha = FsFoldDaisy(cha, tubes[i])
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
