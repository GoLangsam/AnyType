// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsUnsigned

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeUInt8Chan returns a new open channel
// (simply a 'chan uint8' that is).
//
// Note: No 'UInt8-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myUInt8PipelineStartsHere := MakeUInt8Chan()
//	// ... lot's of code to design and build Your favourite "myUInt8WorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myUInt8PipelineStartsHere <- drop
//	}
//	close(myUInt8PipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeUInt8Buffer) the channel is unbuffered.
//
func MakeUInt8Chan() chan uint8 {
	return make(chan uint8)
}

// ChanUInt8 returns a channel to receive all inputs before close.
func ChanUInt8(inp ...uint8) chan uint8 {
	out := make(chan uint8)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanUInt8Slice returns a channel to receive all inputs before close.
func ChanUInt8Slice(inp ...[]uint8) chan uint8 {
	out := make(chan uint8)
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

// JoinUInt8
func JoinUInt8(out chan<- uint8, inp ...uint8) chan struct{} {
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

// JoinUInt8Slice
func JoinUInt8Slice(out chan<- uint8, inp ...[]uint8) chan struct{} {
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

// JoinUInt8Chan
func JoinUInt8Chan(out chan<- uint8, inp <-chan uint8) chan struct{} {
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

// DoneUInt8 returns a channel to receive one signal before close after inp has been drained.
func DoneUInt8(inp <-chan uint8) chan struct{} {
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

// DoneUInt8Slice returns a channel which will receive a slice
// of all the UInt8s received on inp channel before close.
// Unlike DoneUInt8, a full slice is sent once, not just an event.
func DoneUInt8Slice(inp <-chan uint8) chan []uint8 {
	done := make(chan []uint8)
	go func() {
		defer close(done)
		UInt8S := []uint8{}
		for i := range inp {
			UInt8S = append(UInt8S, i)
		}
		done <- UInt8S
	}()
	return done
}

// DoneUInt8Func returns a channel to receive one signal before close after act has been applied to all inp.
func DoneUInt8Func(inp <-chan uint8, act func(a uint8)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a uint8) { return }
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

// PipeUInt8Buffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeUInt8Buffer(inp <-chan uint8, cap int) chan uint8 {
	out := make(chan uint8, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeUInt8Func returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeUInt8Map for functional people,
// but 'map' has a very different meaning in go lang.
func PipeUInt8Func(inp <-chan uint8, act func(a uint8) uint8) chan uint8 {
	out := make(chan uint8)
	if act == nil {
		act = func(a uint8) uint8 { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeUInt8Fork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeUInt8Fork(inp <-chan uint8) (chan uint8, chan uint8) {
	out1 := make(chan uint8)
	out2 := make(chan uint8)
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

// MergeUInt82 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeUInt82(i1, i2 <-chan uint8) (out <-chan uint8) {
	cha := make(chan uint8)
	go func(out chan<- uint8, i1, i2 <-chan uint8) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 uint8 // what we've read
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
