// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsInteger

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUIntChan returns a new open channel
// (simply a 'chan uint' that is).
//
// Note: No 'UInt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUIntPipelineStartsHere := MakeUIntChan()
//	// ... lot's of code to design and build Your favourite "myUIntWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUIntPipelineStartsHere <- drop
//	}
//	close(myUIntPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUIntBuffer) the channel is unbuffered.
//
func MakeUIntChan() (out chan uint) {
	return make(chan uint)
}

func sendUInt(out chan<- uint, inp ...uint) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanUInt returns a channel to receive all inputs before close.
func ChanUInt(inp ...uint) (out <-chan uint) {
	cha := make(chan uint)
	go sendUInt(cha, inp...)
	return cha
}

func sendUIntSlice(out chan<- uint, inp ...[]uint) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanUIntSlice returns a channel to receive all inputs before close.
func ChanUIntSlice(inp ...[]uint) (out <-chan uint) {
	cha := make(chan uint)
	go sendUIntSlice(cha, inp...)
	return cha
}

func joinUInt(done chan<- struct{}, out chan<- uint, inp ...uint) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUInt
func JoinUInt(out chan<- uint, inp ...uint) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUInt(cha, out, inp...)
	return cha
}

func joinUIntSlice(done chan<- struct{}, out chan<- uint, inp ...[]uint) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinUIntSlice
func JoinUIntSlice(out chan<- uint, inp ...[]uint) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUIntSlice(cha, out, inp...)
	return cha
}

func joinUIntChan(done chan<- struct{}, out chan<- uint, inp <-chan uint) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinUIntChan
func JoinUIntChan(out chan<- uint, inp <-chan uint) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinUIntChan(cha, out, inp)
	return cha
}

func doitUInt(done chan<- struct{}, inp <-chan uint) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneUInt returns a channel to receive one signal before close after inp has been drained.
func DoneUInt(inp <-chan uint) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitUInt(cha, inp)
	return cha
}

func doitUIntSlice(done chan<- ([]uint), inp <-chan uint) {
	defer close(done)
	UIntS := []uint{}
	for i := range inp {
		UIntS = append(UIntS, i)
	}
	done <- UIntS
}

// DoneUIntSlice returns a channel which will receive a slice
// of all the UInts received on inp channel before close.
// Unlike DoneUInt, a full slice is sent once, not just an event.
func DoneUIntSlice(inp <-chan uint) (done <-chan ([]uint)) {
	cha := make(chan ([]uint))
	go doitUIntSlice(cha, inp)
	return cha
}

func doitUIntFunc(done chan<- struct{}, inp <-chan uint, act func(a uint)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneUIntFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUIntFunc(inp <-chan uint, act func(a uint)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a uint) { return }
	}
	go doitUIntFunc(cha, inp, act)
	return cha
}

func pipeUIntBuffer(out chan<- uint, inp <-chan uint) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeUIntBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUIntBuffer(inp <-chan uint, cap int) (out <-chan uint) {
	cha := make(chan uint, cap)
	go pipeUIntBuffer(cha, inp)
	return cha
}

func pipeUIntFunc(out chan<- uint, inp <-chan uint, act func(a uint) uint) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeUIntFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUIntMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUIntFunc(inp <-chan uint, act func(a uint) uint) (out <-chan uint) {
	cha := make(chan uint)
	if act == nil {
		act = func(a uint) uint { return a }
	}
	go pipeUIntFunc(cha, inp, act)
	return cha
}

func pipeUIntFork(out1, out2 chan<- uint, inp <-chan uint) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeUIntFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUIntFork(inp <-chan uint) (out1, out2 <-chan uint) {
	cha1 := make(chan uint)
	cha2 := make(chan uint)
	go pipeUIntFork(cha1, cha2, inp)
	return cha1, cha2
}

// MergeUInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeUInt2(i1, i2 <-chan uint) (out <-chan uint) {
	cha := make(chan uint)
	go func(out chan<- uint, i1, i2 <-chan uint) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 uint // what we've read
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
