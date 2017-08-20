// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeInt64Chan returns a new open channel
// (simply a 'chan int64' that is).
//
// Note: No 'Int64-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myInt64PipelineStartsHere := MakeInt64Chan()
//	// ... lot's of code to design and build Your favourite "myInt64WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myInt64PipelineStartsHere <- drop
//	}
//	close(myInt64PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeInt64Buffer) the channel is unbuffered.
//
func MakeInt64Chan() chan int64 {
	return make(chan int64)
}

// ChanInt64 returns a channel to receive all inputs before close.
func ChanInt64(inp ...int64) chan int64 {
	out := make(chan int64)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanInt64Slice returns a channel to receive all inputs before close.
func ChanInt64Slice(inp ...[]int64) chan int64 {
	out := make(chan int64)
	go func() {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}()
	return out
}

// JoinInt64 sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64(out chan<- int64, inp ...int64) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// JoinInt64Slice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64Slice(out chan<- int64, inp ...[]int64) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinInt64Chan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinInt64Chan(out chan<- int64, inp <-chan int64) chan struct{} {
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

// DoneInt64 returns a channel to receive one signal before close after inp has been drained.
func DoneInt64(inp <-chan int64) chan struct{} {
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

// DoneInt64Slice returns a channel which will receive a slice
// of all the Int64s received on inp channel before close.
// Unlike DoneInt64, a full slice is sent once, not just an event.
func DoneInt64Slice(inp <-chan int64) chan []int64 {
	done := make(chan []int64)
	go func() {
		defer close(done)
		Int64S := []int64{}
		for i := range inp {
			Int64S = append(Int64S, i)
		}
		done <- Int64S
	}()
	return done
}

// DoneInt64Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneInt64Func(inp <-chan int64, act func(a int64)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a int64) { return }
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

// PipeInt64Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeInt64Buffer(inp <-chan int64, cap int) chan int64 {
	out := make(chan int64, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeInt64Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeInt64Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeInt64Func(inp <-chan int64, act func(a int64) int64) chan int64 {
	out := make(chan int64)
	if act == nil {
		act = func(a int64) int64 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeInt64Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeInt64Fork(inp <-chan int64) (chan int64, chan int64) {
	out1 := make(chan int64)
	out2 := make(chan int64)
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
