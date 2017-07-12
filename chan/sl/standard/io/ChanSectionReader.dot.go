// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type SectionReaderChan interface { // bidirectional channel
	SectionReaderROnlyChan // aka "<-chan" - receive only
	SectionReaderSOnlyChan // aka "chan<-" - send only
}

type SectionReaderROnlyChan interface { // receive-only channel
	RequestSectionReader() (dat *io.SectionReader)        // the receive function - aka "some-new-SectionReader-var := <-MyKind"
	TrySectionReader() (dat *io.SectionReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-SectionReader-var, ok := <-MyKind"
}

type SectionReaderSOnlyChan interface { // send-only channel
	ProvideSectionReader(dat *io.SectionReader) // the send function - aka "MyKind <- some SectionReader"
}
