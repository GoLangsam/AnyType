// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeSomeTypeChan returns a new open channel
// (simply a 'chan SomeType' that is).
//
// Note: No 'SomeType-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySomeTypePipelineStartsHere := MakeSomeTypeChan()
//	// ... lot's of code to design and build Your favourite "mySomeTypeWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySomeTypePipelineStartsHere <- drop
//	}
//	close(mySomeTypePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSomeTypeBuffer) the channel is unbuffered.
//
func MakeSomeTypeChan() chan SomeType {
	return make(chan SomeType)
}

// ChanSomeType returns a channel to receive all inputs before close.
func ChanSomeType(inp ...SomeType) chan SomeType {
	out := make(chan SomeType)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanSomeTypeSlice returns a channel to receive all inputs before close.
func ChanSomeTypeSlice(inp ...[]SomeType) chan SomeType {
	out := make(chan SomeType)
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

// JoinSomeType
func JoinSomeType(out chan<- SomeType, inp ...SomeType) chan struct{} {
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

// JoinSomeTypeSlice
func JoinSomeTypeSlice(out chan<- SomeType, inp ...[]SomeType) chan struct{} {
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

// JoinSomeTypeChan
func JoinSomeTypeChan(out chan<- SomeType, inp <-chan SomeType) chan struct{} {
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

// DoneSomeType returns a channel to receive one signal before close after inp has been drained.
func DoneSomeType(inp <-chan SomeType) chan struct{} {
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

// DoneSomeTypeSlice returns a channel which will receive a slice
// of all the SomeTypes received on inp channel before close.
// Unlike DoneSomeType, a full slice is sent once, not just an event.
func DoneSomeTypeSlice(inp <-chan SomeType) chan []SomeType {
	done := make(chan []SomeType)
	go func() {
		defer close(done)
		SomeTypeS := []SomeType{}
		for i := range inp {
			SomeTypeS = append(SomeTypeS, i)
		}
		done <- SomeTypeS
	}()
	return done
}

// DoneSomeTypeFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSomeTypeFunc(inp <-chan SomeType, act func(a SomeType)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a SomeType) { return }
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

// PipeSomeTypeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSomeTypeBuffer(inp <-chan SomeType, cap int) chan SomeType {
	out := make(chan SomeType, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeSomeTypeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSomeTypeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSomeTypeFunc(inp <-chan SomeType, act func(a SomeType) SomeType) chan SomeType {
	out := make(chan SomeType)
	if act == nil {
		act = func(a SomeType) SomeType { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeSomeTypeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSomeTypeFork(inp <-chan SomeType) (chan SomeType, chan SomeType) {
	out1 := make(chan SomeType)
	out2 := make(chan SomeType)
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
