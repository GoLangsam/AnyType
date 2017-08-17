// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeSomeTypeSChan returns a new open channel
// (simply a 'chan []SomeType' that is).
//
// Note: No 'SomeTypeS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySomeTypeSPipelineStartsHere := MakeSomeTypeSChan()
//	// ... lot's of code to design and build Your favourite "mySomeTypeSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySomeTypeSPipelineStartsHere <- drop
//	}
//	close(mySomeTypeSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSomeTypeSBuffer) the channel is unbuffered.
//
func MakeSomeTypeSChan() (out chan []SomeType) {
	return make(chan []SomeType)
}

// ChanSomeTypeS returns a channel to receive all inputs before close.
func ChanSomeTypeS(inp ...[]SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	go func(out chan<- []SomeType, inp ...[]SomeType) {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}(cha, inp...)
	return cha
}

// ChanSomeTypeSSlice returns a channel to receive all inputs before close.
func ChanSomeTypeSSlice(inp ...[][]SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	go func(out chan<- []SomeType, inp ...[][]SomeType) {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}(cha, inp...)
	return cha
}

// JoinSomeTypeS
func JoinSomeTypeS(out chan<- []SomeType, inp ...[]SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []SomeType, inp ...[]SomeType) {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSomeTypeSSlice
func JoinSomeTypeSSlice(out chan<- []SomeType, inp ...[][]SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []SomeType, inp ...[][]SomeType) {
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

// JoinSomeTypeSChan
func JoinSomeTypeSChan(out chan<- []SomeType, inp <-chan []SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- []SomeType, inp <-chan []SomeType) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneSomeTypeS returns a channel to receive one signal before close after inp has been drained.
func DoneSomeTypeS(inp <-chan []SomeType) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan []SomeType) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSomeTypeSSlice returns a channel which will receive a slice
// of all the SomeTypeSs received on inp channel before close.
// Unlike DoneSomeTypeS, a full slice is sent once, not just an event.
func DoneSomeTypeSSlice(inp <-chan []SomeType) (done <-chan [][]SomeType) {
	cha := make(chan [][]SomeType)
	go func(inp <-chan []SomeType, done chan<- [][]SomeType) {
		defer close(done)
		SomeTypeSS := [][]SomeType{}
		for i := range inp {
			SomeTypeSS = append(SomeTypeSS, i)
		}
		done <- SomeTypeSS
	}(inp, cha)
	return cha
}

// DoneSomeTypeSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSomeTypeSFunc(inp <-chan []SomeType, act func(a []SomeType)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []SomeType) { return }
	}
	go func(done chan<- struct{}, inp <-chan []SomeType, act func(a []SomeType)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeSomeTypeSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSomeTypeSBuffer(inp <-chan []SomeType, cap int) (out <-chan []SomeType) {
	cha := make(chan []SomeType, cap)
	go func(out chan<- []SomeType, inp <-chan []SomeType) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeSomeTypeSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSomeTypeSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSomeTypeSFunc(inp <-chan []SomeType, act func(a []SomeType) []SomeType) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	if act == nil {
		act = func(a []SomeType) []SomeType { return a }
	}
	go func(out chan<- []SomeType, inp <-chan []SomeType, act func(a []SomeType) []SomeType) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeSomeTypeSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSomeTypeSFork(inp <-chan []SomeType) (out1, out2 <-chan []SomeType) {
	cha1 := make(chan []SomeType)
	cha2 := make(chan []SomeType)
	go func(out1, out2 chan<- []SomeType, inp <-chan []SomeType) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// SomeTypeSTube is the signature for a pipe function.
type SomeTypeSTube func(inp <-chan []SomeType, out <-chan []SomeType)

// SomeTypeSdaisy returns a channel to receive all inp after having passed thru tube.
func SomeTypeSdaisy(inp <-chan []SomeType, tube SomeTypeSTube) (out <-chan []SomeType) {
	cha := make(chan []SomeType)
	go tube(inp, cha)
	return cha
}

// SomeTypeSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SomeTypeSDaisyChain(inp <-chan []SomeType, tubes ...SomeTypeSTube) (out <-chan []SomeType) {
	cha := inp
	for _, tube := range tubes {
		cha = SomeTypeSdaisy(cha, tube)
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
