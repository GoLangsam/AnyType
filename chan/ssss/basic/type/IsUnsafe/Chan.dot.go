// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsUnsafe

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeChan returns a new open channel
// (simply a 'chan uintptr' that is).
//
// Note: No '-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipelineStartsHere := MakeChan()
//	// ... lot's of code to design and build Your favourite "myWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipelineStartsHere <- drop
//	}
//	close(myPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBuffer) the channel is unbuffered.
//
func MakeChan() chan uintptr {
	return make(chan uintptr)
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...uintptr) chan uintptr {
	out := make(chan uintptr)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]uintptr) chan uintptr {
	out := make(chan uintptr)
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

// Join
func Join(out chan<- uintptr, inp ...uintptr) chan struct{} {
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

// JoinSlice
func JoinSlice(out chan<- uintptr, inp ...[]uintptr) chan struct{} {
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

// JoinChan
func JoinChan(out chan<- uintptr, inp <-chan uintptr) chan struct{} {
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

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan uintptr) chan struct{} {
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

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan uintptr) chan []uintptr {
	done := make(chan []uintptr)
	go func() {
		defer close(done)
		S := []uintptr{}
		for i := range inp {
			S = append(S, i)
		}
		done <- S
	}()
	return done
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan uintptr, act func(a uintptr)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a uintptr) { return }
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

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan uintptr, cap int) chan uintptr {
	out := make(chan uintptr, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan uintptr, act func(a uintptr) uintptr) chan uintptr {
	out := make(chan uintptr)
	if act == nil {
		act = func(a uintptr) uintptr { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan uintptr) (chan uintptr, chan uintptr) {
	out1 := make(chan uintptr)
	out2 := make(chan uintptr)
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
