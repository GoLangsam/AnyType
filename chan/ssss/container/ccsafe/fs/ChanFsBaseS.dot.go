// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsBaseSChan returns a new open channel
// (simply a 'chan fs.FsBaseS' that is).
//
// Note: No 'FsBaseS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsBaseSPipelineStartsHere := MakeFsBaseSChan()
//	// ... lot's of code to design and build Your favourite "myFsBaseSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsBaseSPipelineStartsHere <- drop
//	}
//	close(myFsBaseSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsBaseSBuffer) the channel is unbuffered.
//
func MakeFsBaseSChan() chan fs.FsBaseS {
	return make(chan fs.FsBaseS)
}

// ChanFsBaseS returns a channel to receive all inputs before close.
func ChanFsBaseS(inp ...fs.FsBaseS) chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFsBaseSSlice returns a channel to receive all inputs before close.
func ChanFsBaseSSlice(inp ...[]fs.FsBaseS) chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
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

// ChanFsBaseSFuncNil returns a channel to receive all results of act until nil before close.
func ChanFsBaseSFuncNil(act func() fs.FsBaseS) <-chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
	go func() {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanFsBaseSFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsBaseSFuncNok(act func() (fs.FsBaseS, bool)) <-chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanFsBaseSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsBaseSFuncErr(act func() (fs.FsBaseS, error)) <-chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinFsBaseS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBaseS(out chan<- fs.FsBaseS, inp ...fs.FsBaseS) chan struct{} {
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

// JoinFsBaseSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBaseSSlice(out chan<- fs.FsBaseS, inp ...[]fs.FsBaseS) chan struct{} {
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

// JoinFsBaseSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsBaseSChan(out chan<- fs.FsBaseS, inp <-chan fs.FsBaseS) chan struct{} {
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

// DoneFsBaseS returns a channel to receive one signal before close after inp has been drained.
func DoneFsBaseS(inp <-chan fs.FsBaseS) chan struct{} {
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

// DoneFsBaseSSlice returns a channel which will receive a slice
// of all the FsBaseSs received on inp channel before close.
// Unlike DoneFsBaseS, a full slice is sent once, not just an event.
func DoneFsBaseSSlice(inp <-chan fs.FsBaseS) chan []fs.FsBaseS {
	done := make(chan []fs.FsBaseS)
	go func() {
		defer close(done)
		FsBaseSS := []fs.FsBaseS{}
		for i := range inp {
			FsBaseSS = append(FsBaseSS, i)
		}
		done <- FsBaseSS
	}()
	return done
}

// DoneFsBaseSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsBaseSFunc(inp <-chan fs.FsBaseS, act func(a fs.FsBaseS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.FsBaseS) { return }
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

// PipeFsBaseSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsBaseSBuffer(inp <-chan fs.FsBaseS, cap int) chan fs.FsBaseS {
	out := make(chan fs.FsBaseS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsBaseSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsBaseSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsBaseSFunc(inp <-chan fs.FsBaseS, act func(a fs.FsBaseS) fs.FsBaseS) chan fs.FsBaseS {
	out := make(chan fs.FsBaseS)
	if act == nil {
		act = func(a fs.FsBaseS) fs.FsBaseS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsBaseSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsBaseSFork(inp <-chan fs.FsBaseS) (chan fs.FsBaseS, chan fs.FsBaseS) {
	out1 := make(chan fs.FsBaseS)
	out2 := make(chan fs.FsBaseS)
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

// FsBaseSTube is the signature for a pipe function.
type FsBaseSTube func(inp <-chan fs.FsBaseS, out <-chan fs.FsBaseS)

// FsBaseSDaisy returns a channel to receive all inp after having passed thru tube.
func FsBaseSDaisy(inp <-chan fs.FsBaseS, tube FsBaseSTube) (out <-chan fs.FsBaseS) {
	cha := make(chan fs.FsBaseS)
	go tube(inp, cha)
	return cha
}

// FsBaseSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsBaseSDaisyChain(inp <-chan fs.FsBaseS, tubes ...FsBaseSTube) (out <-chan fs.FsBaseS) {
	cha := inp
	for i := range tubes {
		cha = FsBaseSDaisy(cha, tubes[i])
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
