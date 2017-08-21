// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

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
func MakeFsFoldChan() chan *fs.FsFold {
	return make(chan *fs.FsFold)
}

// ChanFsFold returns a channel to receive all inputs before close.
func ChanFsFold(inp ...*fs.FsFold) chan *fs.FsFold {
	out := make(chan *fs.FsFold)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFsFoldSlice returns a channel to receive all inputs before close.
func ChanFsFoldSlice(inp ...[]*fs.FsFold) chan *fs.FsFold {
	out := make(chan *fs.FsFold)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanFsFoldFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsFoldFuncNok(act func() (*fs.FsFold, bool)) <-chan *fs.FsFold {
	out := make(chan *fs.FsFold)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			} else {
				out <- res
			}
		}
	}()
	return out
}

// ChanFsFoldFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsFoldFuncErr(act func() (*fs.FsFold, error)) <-chan *fs.FsFold {
	out := make(chan *fs.FsFold)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			} else {
				out <- res
			}
		}
	}()
	return out
}

// JoinFsFold sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFold(out chan<- *fs.FsFold, inp ...*fs.FsFold) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFsFoldSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFoldSlice(out chan<- *fs.FsFold, inp ...[]*fs.FsFold) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFsFoldChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFoldChan(out chan<- *fs.FsFold, inp <-chan *fs.FsFold) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFsFold returns a channel to receive one signal before close after inp has been drained.
func DoneFsFold(inp <-chan *fs.FsFold) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFsFoldSlice returns a channel which will receive a slice
// of all the FsFolds received on inp channel before close.
// Unlike DoneFsFold, a full slice is sent once, not just an event.
func DoneFsFoldSlice(inp <-chan *fs.FsFold) chan []*fs.FsFold {
	done := make(chan []*fs.FsFold)
	go func() {
		defer close(done)
		FsFoldS := []*fs.FsFold{}
		for i := range inp {
			FsFoldS = append(FsFoldS, i)
		}
		done <- FsFoldS
	}()
	return done
}

// DoneFsFoldFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsFold) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeFsFoldBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFoldBuffer(inp <-chan *fs.FsFold, cap int) chan *fs.FsFold {
	out := make(chan *fs.FsFold, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsFoldFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFoldMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold) *fs.FsFold) chan *fs.FsFold {
	out := make(chan *fs.FsFold)
	if act == nil {
		act = func(a *fs.FsFold) *fs.FsFold { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsFoldFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFoldFork(inp <-chan *fs.FsFold) (chan *fs.FsFold, chan *fs.FsFold) {
	out1 := make(chan *fs.FsFold)
	out2 := make(chan *fs.FsFold)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
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
