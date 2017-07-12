// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteWriterChan returns a new open channel
// (simply a 'chan io.ByteWriter' that is).
//
// Note: No 'ByteWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteWriterPipelineStartsHere := MakeByteWriterChan()
//	// ... lot's of code to design and build Your favourite "myByteWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteWriterPipelineStartsHere <- drop
//	}
//	close(myByteWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteWriterBuffer) the channel is unbuffered.
//
func MakeByteWriterChan() chan io.ByteWriter {
	return make(chan io.ByteWriter)
}

// ChanByteWriter returns a channel to receive all inputs before close.
func ChanByteWriter(inp ...io.ByteWriter) chan io.ByteWriter {
	out := make(chan io.ByteWriter)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanByteWriterSlice returns a channel to receive all inputs before close.
func ChanByteWriterSlice(inp ...[]io.ByteWriter) chan io.ByteWriter {
	out := make(chan io.ByteWriter)
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

// JoinByteWriter
func JoinByteWriter(out chan<- io.ByteWriter, inp ...io.ByteWriter) chan struct{} {
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

// JoinByteWriterSlice
func JoinByteWriterSlice(out chan<- io.ByteWriter, inp ...[]io.ByteWriter) chan struct{} {
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

// JoinByteWriterChan
func JoinByteWriterChan(out chan<- io.ByteWriter, inp <-chan io.ByteWriter) chan struct{} {
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

// DoneByteWriter returns a channel to receive one signal before close after inp has been drained.
func DoneByteWriter(inp <-chan io.ByteWriter) chan struct{} {
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

// DoneByteWriterSlice returns a channel which will receive a slice
// of all the ByteWriters received on inp channel before close.
// Unlike DoneByteWriter, a full slice is sent once, not just an event.
func DoneByteWriterSlice(inp <-chan io.ByteWriter) chan []io.ByteWriter {
	done := make(chan []io.ByteWriter)
	go func() {
		defer close(done)
		ByteWriterS := []io.ByteWriter{}
		for i := range inp {
			ByteWriterS = append(ByteWriterS, i)
		}
		done <- ByteWriterS
	}()
	return done
}

// DoneByteWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ByteWriter) { return }
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

// PipeByteWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteWriterBuffer(inp <-chan io.ByteWriter, cap int) chan io.ByteWriter {
	out := make(chan io.ByteWriter, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeByteWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter) io.ByteWriter) chan io.ByteWriter {
	out := make(chan io.ByteWriter)
	if act == nil {
		act = func(a io.ByteWriter) io.ByteWriter { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeByteWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteWriterFork(inp <-chan io.ByteWriter) (chan io.ByteWriter, chan io.ByteWriter) {
	out1 := make(chan io.ByteWriter)
	out2 := make(chan io.ByteWriter)
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
