// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeByteChan returns a new open channel
// (simply a 'chan byte' that is).
//
// Note: No 'Byte-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myBytePipelineStartsHere := MakeByteChan()
//	// ... lot's of code to design and build Your favourite "myByteWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myBytePipelineStartsHere <- drop
//	}
//	close(myBytePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteBuffer) the channel is unbuffered.
//
func MakeByteChan() (out chan byte) {
	return make(chan byte)
}

func sendByte(out chan<- byte, inp ...byte) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanByte returns a channel to receive all inputs before close.
func ChanByte(inp ...byte) (out <-chan byte) {
	cha := make(chan byte)
	go sendByte(cha, inp...)
	return cha
}

func sendByteSlice(out chan<- byte, inp ...[]byte) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanByteSlice returns a channel to receive all inputs before close.
func ChanByteSlice(inp ...[]byte) (out <-chan byte) {
	cha := make(chan byte)
	go sendByteSlice(cha, inp...)
	return cha
}

func joinByte(done chan<- struct{}, out chan<- byte, inp ...byte) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinByte
func JoinByte(out chan<- byte, inp ...byte) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByte(cha, out, inp...)
	return cha
}

func joinByteSlice(done chan<- struct{}, out chan<- byte, inp ...[]byte) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinByteSlice
func JoinByteSlice(out chan<- byte, inp ...[]byte) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteSlice(cha, out, inp...)
	return cha
}

func joinByteChan(done chan<- struct{}, out chan<- byte, inp <-chan byte) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinByteChan
func JoinByteChan(out chan<- byte, inp <-chan byte) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteChan(cha, out, inp)
	return cha
}

func doitByte(done chan<- struct{}, inp <-chan byte) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneByte returns a channel to receive one signal before close after inp has been drained.
func DoneByte(inp <-chan byte) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitByte(cha, inp)
	return cha
}

func doitByteSlice(done chan<- ([]byte), inp <-chan byte) {
	defer close(done)
	ByteS := []byte{}
	for i := range inp {
		ByteS = append(ByteS, i)
	}
	done <- ByteS
}

// DoneByteSlice returns a channel which will receive a slice
// of all the Bytes received on inp channel before close.
// Unlike DoneByte, a full slice is sent once, not just an event.
func DoneByteSlice(inp <-chan byte) (done <-chan ([]byte)) {
	cha := make(chan ([]byte))
	go doitByteSlice(cha, inp)
	return cha
}

func doitByteFunc(done chan<- struct{}, inp <-chan byte, act func(a byte)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneByteFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteFunc(inp <-chan byte, act func(a byte)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a byte) { return }
	}
	go doitByteFunc(cha, inp, act)
	return cha
}

func pipeByteBuffer(out chan<- byte, inp <-chan byte) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeByteBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteBuffer(inp <-chan byte, cap int) (out <-chan byte) {
	cha := make(chan byte, cap)
	go pipeByteBuffer(cha, inp)
	return cha
}

func pipeByteFunc(out chan<- byte, inp <-chan byte, act func(a byte) byte) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeByteFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteFunc(inp <-chan byte, act func(a byte) byte) (out <-chan byte) {
	cha := make(chan byte)
	if act == nil {
		act = func(a byte) byte { return a }
	}
	go pipeByteFunc(cha, inp, act)
	return cha
}

func pipeByteFork(out1, out2 chan<- byte, inp <-chan byte) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeByteFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteFork(inp <-chan byte) (out1, out2 <-chan byte) {
	cha1 := make(chan byte)
	cha2 := make(chan byte)
	go pipeByteFork(cha1, cha2, inp)
	return cha1, cha2
}
