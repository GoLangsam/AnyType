// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsDataChan returns a new open channel
// (simply a 'chan *fs.FsData' that is).
//
// Note: No 'FsData-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsDataPipelineStartsHere := MakeFsDataChan()
//	// ... lot's of code to design and build Your favourite "myFsDataWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsDataPipelineStartsHere <- drop
//	}
//	close(myFsDataPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsDataBuffer) the channel is unbuffered.
//
func MakeFsDataChan() chan *fs.FsData {
	return make(chan *fs.FsData)
}

// ChanFsData returns a channel to receive all inputs before close.
func ChanFsData(inp ...*fs.FsData) chan *fs.FsData {
	out := make(chan *fs.FsData)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFsDataSlice returns a channel to receive all inputs before close.
func ChanFsDataSlice(inp ...[]*fs.FsData) chan *fs.FsData {
	out := make(chan *fs.FsData)
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

// ChanFsDataFuncNil returns a channel to receive all results of act until nil before close.
func ChanFsDataFuncNil(act func() *fs.FsData) <-chan *fs.FsData {
	out := make(chan *fs.FsData)
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

// ChanFsDataFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsDataFuncNok(act func() (*fs.FsData, bool)) <-chan *fs.FsData {
	out := make(chan *fs.FsData)
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

// ChanFsDataFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsDataFuncErr(act func() (*fs.FsData, error)) <-chan *fs.FsData {
	out := make(chan *fs.FsData)
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

// JoinFsData sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsData(out chan<- *fs.FsData, inp ...*fs.FsData) chan struct{} {
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

// JoinFsDataSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsDataSlice(out chan<- *fs.FsData, inp ...[]*fs.FsData) chan struct{} {
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

// JoinFsDataChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsDataChan(out chan<- *fs.FsData, inp <-chan *fs.FsData) chan struct{} {
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

// DoneFsData returns a channel to receive one signal before close after inp has been drained.
func DoneFsData(inp <-chan *fs.FsData) chan struct{} {
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

// DoneFsDataSlice returns a channel which will receive a slice
// of all the FsDatas received on inp channel before close.
// Unlike DoneFsData, a full slice is sent once, not just an event.
func DoneFsDataSlice(inp <-chan *fs.FsData) chan []*fs.FsData {
	done := make(chan []*fs.FsData)
	go func() {
		defer close(done)
		FsDataS := []*fs.FsData{}
		for i := range inp {
			FsDataS = append(FsDataS, i)
		}
		done <- FsDataS
	}()
	return done
}

// DoneFsDataFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsDataFunc(inp <-chan *fs.FsData, act func(a *fs.FsData)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsData) { return }
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

// PipeFsDataBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsDataBuffer(inp <-chan *fs.FsData, cap int) chan *fs.FsData {
	out := make(chan *fs.FsData, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFsDataFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsDataMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsDataFunc(inp <-chan *fs.FsData, act func(a *fs.FsData) *fs.FsData) chan *fs.FsData {
	out := make(chan *fs.FsData)
	if act == nil {
		act = func(a *fs.FsData) *fs.FsData { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFsDataFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsDataFork(inp <-chan *fs.FsData) (chan *fs.FsData, chan *fs.FsData) {
	out1 := make(chan *fs.FsData)
	out2 := make(chan *fs.FsData)
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

// FsDataTube is the signature for a pipe function.
type FsDataTube func(inp <-chan *fs.FsData, out <-chan *fs.FsData)

// FsDataDaisy returns a channel to receive all inp after having passed thru tube.
func FsDataDaisy(inp <-chan *fs.FsData, tube FsDataTube) (out <-chan *fs.FsData) {
	cha := make(chan *fs.FsData)
	go tube(inp, cha)
	return cha
}

// FsDataDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsDataDaisyChain(inp <-chan *fs.FsData, tubes ...FsDataTube) (out <-chan *fs.FsData) {
	cha := inp
	for i := range tubes {
		cha = FsDataDaisy(cha, tubes[i])
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
