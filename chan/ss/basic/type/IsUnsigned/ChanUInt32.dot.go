// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsUnsigned

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

func sendUInt32(out chan<- uint32, inp ...uint32) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanUInt32 returns a channel to receive all inputs before close.
func ChanUInt32(inp ...uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go sendUInt32(cha, inp...)
	return cha
}

func sendUInt32Slice(out chan<- uint32, inp ...[]uint32) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanUInt32Slice returns a channel to receive all inputs before close.
func ChanUInt32Slice(inp ...[]uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go sendUInt32Slice(cha, inp...)
	return cha
}

func joinUInt32(done chan<- struct{}, out chan<- uint32, inp ...uint32) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt32
func JoinUInt32(out chan<- uint32, inp ...uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt32(cha, out, inp...)
	return cha
}

func joinUInt32Slice(done chan<- struct{}, out chan<- uint32, inp ...[]uint32) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinUInt32Slice
func JoinUInt32Slice(out chan<- uint32, inp ...[]uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt32Slice(cha, out, inp...)
	return cha
}

func joinUInt32Chan(done chan<- struct{}, out chan<- uint32, inp <-chan uint32) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt32Chan
func JoinUInt32Chan(out chan<- uint32, inp <-chan uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt32Chan(cha, out, inp)
	return cha
}

func doitUInt32(done chan<- struct{}, inp <-chan uint32) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneUInt32 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt32(inp <-chan uint32) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitUInt32(cha, inp)
	return cha
}

func doitUInt32Slice(done chan<- ([]uint32), inp <-chan uint32) {
	defer close(done)
	UInt32S := []uint32{}
	for i := range inp {
		UInt32S = append(UInt32S, i)
	}
	done <- UInt32S
}

// DoneUInt32Slice returns a channel which will receive a slice
// of all the UInt32s received on inp channel before close.
// Unlike DoneUInt32, a full slice is sent once, not just an event.
func DoneUInt32Slice(inp <-chan uint32) (done <-chan ([]uint32)) {
	cha := make(chan ([]uint32))
	go doitUInt32Slice(cha, inp)
	return cha
}

func doitUInt32Func(done chan<- struct{}, inp <-chan uint32, act func(a uint32)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneUInt32Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt32Func(inp <-chan uint32, act func(a uint32)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint32) { return }
	}
	go doitUInt32Func(cha, inp, act)
	return cha
}

func pipeUInt32Buffer(out chan<- uint32, inp <-chan uint32) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeUInt32Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt32Buffer(inp <-chan uint32, cap int) (out <-chan uint32) {
	cha := make(chan uint32, cap)
	go pipeUInt32Buffer(cha, inp)
	return cha
}

func pipeUInt32Func(out chan<- uint32, inp <-chan uint32, act func(a uint32) uint32) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeUInt32Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt32Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt32Func(inp <-chan uint32, act func(a uint32) uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	if act == nil {
		act = func(a uint32) uint32 { return a }
	}
	go pipeUInt32Func(cha, inp, act)
	return cha
}

func pipeUInt32Fork(out1, out2 chan<- uint32, inp <-chan uint32) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeUInt32Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt32Fork(inp <-chan uint32) (out1, out2 <-chan uint32) {
	cha1 := make(chan uint32)
	cha2 := make(chan uint32)
	go pipeUInt32Fork(cha1, cha2, inp)
	return cha1, cha2
}

// MergeUInt322 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeUInt322(i1, i2 <-chan uint32) (out <-chan uint32) {
	cha := make(chan uint32)
	go func(out chan<- uint32, i1, i2 <-chan uint32) {
		defer close(out)
		var (
			clos1, clos2 bool   // we found the chan closed
			buff1, buff2 bool   // we've read 'from', but not sent (yet)
			ok           bool   // did we read sucessfully?
			from1, from2 uint32 // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}
