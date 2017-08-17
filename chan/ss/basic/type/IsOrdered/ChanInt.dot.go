// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeIntChan returns a new open channel
// (simply a 'chan int' that is).
//
// Note: No 'Int-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myIntPipelineStartsHere := MakeIntChan()
//	// ... lot's of code to design and build Your favourite "myIntWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myIntPipelineStartsHere <- drop
//	}
//	close(myIntPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeIntBuffer) the channel is unbuffered.
//
func MakeIntChan() (out chan int) {
	return make(chan int)
}

func sendInt(out chan<- int, inp ...int) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanInt returns a channel to receive all inputs before close.
func ChanInt(inp ...int) (out <-chan int) {
	cha := make(chan int)
	go sendInt(cha, inp...)
	return cha
}

func sendIntSlice(out chan<- int, inp ...[]int) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanIntSlice returns a channel to receive all inputs before close.
func ChanIntSlice(inp ...[]int) (out <-chan int) {
	cha := make(chan int)
	go sendIntSlice(cha, inp...)
	return cha
}

func joinInt(done chan<- struct{}, out chan<- int, inp ...int) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinInt
func JoinInt(out chan<- int, inp ...int) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinInt(cha, out, inp...)
	return cha
}

func joinIntSlice(done chan<- struct{}, out chan<- int, inp ...[]int) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinIntSlice
func JoinIntSlice(out chan<- int, inp ...[]int) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinIntSlice(cha, out, inp...)
	return cha
}

func joinIntChan(done chan<- struct{}, out chan<- int, inp <-chan int) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinIntChan
func JoinIntChan(out chan<- int, inp <-chan int) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinIntChan(cha, out, inp)
	return cha
}

func doitInt(done chan<- struct{}, inp <-chan int) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneInt returns a channel to receive one signal before close after inp has been drained.
func DoneInt(inp <-chan int) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitInt(cha, inp)
	return cha
}

func doitIntSlice(done chan<- ([]int), inp <-chan int) {
	defer close(done)
	IntS := []int{}
	for i := range inp {
		IntS = append(IntS, i)
	}
	done <- IntS
}

// DoneIntSlice returns a channel which will receive a slice
// of all the Ints received on inp channel before close.
// Unlike DoneInt, a full slice is sent once, not just an event.
func DoneIntSlice(inp <-chan int) (done <-chan ([]int)) {
	cha := make(chan ([]int))
	go doitIntSlice(cha, inp)
	return cha
}

func doitIntFunc(done chan<- struct{}, inp <-chan int, act func(a int)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneIntFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneIntFunc(inp <-chan int, act func(a int)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a int) { return }
	}
	go doitIntFunc(cha, inp, act)
	return cha
}

func pipeIntBuffer(out chan<- int, inp <-chan int) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeIntBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeIntBuffer(inp <-chan int, cap int) (out <-chan int) {
	cha := make(chan int, cap)
	go pipeIntBuffer(cha, inp)
	return cha
}

func pipeIntFunc(out chan<- int, inp <-chan int, act func(a int) int) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeIntFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeIntMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeIntFunc(inp <-chan int, act func(a int) int) (out <-chan int) {
	cha := make(chan int)
	if act == nil {
		act = func(a int) int { return a }
	}
	go pipeIntFunc(cha, inp, act)
	return cha
}

func pipeIntFork(out1, out2 chan<- int, inp <-chan int) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeIntFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeIntFork(inp <-chan int) (out1, out2 <-chan int) {
	cha1 := make(chan int)
	cha2 := make(chan int)
	go pipeIntFork(cha1, cha2, inp)
	return cha1, cha2
}

// MergeInt returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Intchannel is returned.
func MergeInt(inps ...<-chan int) (out <-chan int) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt2(inps[0], MergeInt(inps[1:]...))
	}
}

// mergeInt2 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt2(i1, i2 <-chan int) (out <-chan int) {
	cha := make(chan int)
	go func(out chan<- int, i1, i2 <-chan int) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read sucessfully?
			from1, from2 int  // what we've read
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

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
