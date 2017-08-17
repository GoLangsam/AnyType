// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsRune

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeRuneSChan returns a new open channel
// (simply a 'chan []rune' that is).
//
// Note: No 'RuneS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myRuneSPipelineStartsHere := MakeRuneSChan()
//	// ... lot's of code to design and build Your favourite "myRuneSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myRuneSPipelineStartsHere <- drop
//	}
//	close(myRuneSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeRuneSBuffer) the channel is unbuffered.
//
func MakeRuneSChan() chan []rune {
	return make(chan []rune)
}

// ChanRuneS returns a channel to receive all inputs before close.
func ChanRuneS(inp ...[]rune) chan []rune {
	out := make(chan []rune)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanRuneSSlice returns a channel to receive all inputs before close.
func ChanRuneSSlice(inp ...[][]rune) chan []rune {
	out := make(chan []rune)
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

// JoinRuneS
func JoinRuneS(out chan<- []rune, inp ...[]rune) chan struct{} {
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

// JoinRuneSSlice
func JoinRuneSSlice(out chan<- []rune, inp ...[][]rune) chan struct{} {
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

// JoinRuneSChan
func JoinRuneSChan(out chan<- []rune, inp <-chan []rune) chan struct{} {
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

// DoneRuneS returns a channel to receive one signal before close after inp has been drained.
func DoneRuneS(inp <-chan []rune) chan struct{} {
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

// DoneRuneSSlice returns a channel which will receive a slice
// of all the RuneSs received on inp channel before close.
// Unlike DoneRuneS, a full slice is sent once, not just an event.
func DoneRuneSSlice(inp <-chan []rune) chan [][]rune {
	done := make(chan [][]rune)
	go func() {
		defer close(done)
		RuneSS := [][]rune{}
		for i := range inp {
			RuneSS = append(RuneSS, i)
		}
		done <- RuneSS
	}()
	return done
}

// DoneRuneSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneRuneSFunc(inp <-chan []rune, act func(a []rune)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a []rune) { return }
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

// PipeRuneSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeRuneSBuffer(inp <-chan []rune, cap int) chan []rune {
	out := make(chan []rune, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeRuneSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeRuneSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeRuneSFunc(inp <-chan []rune, act func(a []rune) []rune) chan []rune {
	out := make(chan []rune)
	if act == nil {
		act = func(a []rune) []rune { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeRuneSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeRuneSFork(inp <-chan []rune) (chan []rune, chan []rune) {
	out1 := make(chan []rune)
	out2 := make(chan []rune)
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
