// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt32Chan returns a new open channel
// (simply a 'chan uint32' that is).
//
// Note: No 'UInt32-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt32PipelineStartsHere := MakeUInt32Chan()
//	// ... lot's of code to design and build Your favourite "myUInt32WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt32PipelineStartsHere <- drop
//	}
//	close(myUInt32PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt32Buffer) the channel is unbuffered.
//
func MakeUInt32Chan() (out chan uint32) {
	return make(chan uint32)
}

// ChanUInt32 returns a channel to receive all inputs before close.
func ChanUInt32(inp ...uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go func(out chan<- uint32, inp ...uint32) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanUInt32Slice returns a channel to receive all inputs before close.
func ChanUInt32Slice(inp ...[]uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go func(out chan<- uint32, inp ...[]uint32) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinUInt32
func JoinUInt32(out chan<- uint32, inp ...uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint32, inp ...uint32) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinUInt32Slice
func JoinUInt32Slice(out chan<- uint32, inp ...[]uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint32, inp ...[]uint32) {
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

// JoinUInt32Chan
func JoinUInt32Chan(out chan<- uint32, inp <-chan uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- uint32, inp <-chan uint32) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneUInt32 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt32(inp <-chan uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan uint32) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneUInt32Slice returns a channel which will receive a slice
// of all the UInt32s received on inp channel before close.
// Unlike DoneUInt32, a full slice is sent once, not just an event.
func DoneUInt32Slice(inp <-chan uint32) (done <-chan []uint32) {
	cha := make(chan []uint32)
	go func(inp <-chan uint32, done chan<- []uint32) {
		defer close(done)
		UInt32S := []uint32{}
		for i := range inp {
			UInt32S = append(UInt32S, i)
		}
		done <- UInt32S
	}(inp, cha)
	return cha
}

// DoneUInt32Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt32Func(inp <-chan uint32, act func(a uint32)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint32) { return }
	}
	go func(done chan<- struct{}, inp <-chan uint32, act func(a uint32)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeUInt32Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt32Buffer(inp <-chan uint32, cap int) (out <-chan uint32) {
	cha := make(chan uint32, cap)
	go func(out chan<- uint32, inp <-chan uint32) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeUInt32Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt32Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt32Func(inp <-chan uint32, act func(a uint32) uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	if act == nil {
		act = func(a uint32) uint32 { return a }
	}
	go func(out chan<- uint32, inp <-chan uint32, act func(a uint32) uint32) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeUInt32Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt32Fork(inp <-chan uint32) (out1, out2 <-chan uint32) {
	cha1 := make(chan uint32)
	cha2 := make(chan uint32)
	go func(out1, out2 chan<- uint32, inp <-chan uint32) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}
