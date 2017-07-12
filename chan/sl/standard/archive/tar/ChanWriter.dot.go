// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/tar"
)

type WriterChan interface { // bidirectional channel
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

type WriterROnlyChan interface { // receive-only channel
	RequestWriter() (dat *tar.Writer)        // the receive function - aka "some-new-Writer-var := <-MyKind"
	TryWriter() (dat *tar.Writer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Writer-var, ok := <-MyKind"
}

type WriterSOnlyChan interface { // send-only channel
	ProvideWriter(dat *tar.Writer) // the send function - aka "MyKind <- some Writer"
}
