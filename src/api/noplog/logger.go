package noplog

import "log"

var NopLog = log.New(NullWriter(1), "", log.LstdFlags)

// NullWriter implements the io.Write interface but doesn't do anything.
type NullWriter int

// Write implements the io.Write interface but is a noop.
func (NullWriter) Write([]byte) (int, error) { return 0, nil }
