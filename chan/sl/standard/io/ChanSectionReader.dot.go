// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// SectionReaderChan represents a
// bidirectional
// channel
type SectionReaderChan interface {
	SectionReaderROnlyChan // aka "<-chan" - receive only
	SectionReaderSOnlyChan // aka "chan<-" - send only
}

// SectionReaderROnlyChan represents a
// receive-only
// channel
type SectionReaderROnlyChan interface {
	RequestSectionReader() (dat *io.SectionReader)        // the receive function - aka "MySectionReader := <-MySectionReaderROnlyChan"
	TrySectionReader() (dat *io.SectionReader, open bool) // the multi-valued comma-ok receive function - aka "MySectionReader, ok := <-MySectionReaderROnlyChan"
}

// SectionReaderSOnlyChan represents a
// send-only
// channel
type SectionReaderSOnlyChan interface {
	ProvideSectionReader(dat *io.SectionReader) // the send function - aka "MyKind <- some SectionReader"
}
