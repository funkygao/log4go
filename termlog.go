// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

// This is the standard writer that prints to standard output.
type ConsoleLogWriter struct {
	records chan *LogRecord
	flushed chan struct{}
}

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() *ConsoleLogWriter {
	writer := &ConsoleLogWriter{
		records: make(chan *LogRecord, LogBufferLength),
		flushed: make(chan struct{}),
	}
	go writer.run(stdout)
	return writer
}

func (w *ConsoleLogWriter) run(out io.Writer) {
	var timestr string
	var timestrAt int64

	for rec := range w.records {
		if at := rec.Created.UnixNano() / 1e9; at != timestrAt {
			timestr, timestrAt = rec.Created.Format("01/02/06 15:04:05"), at
		}
		fmt.Fprint(out, "[", timestr, "] [", levelStrings[rec.Level], "] (", rec.Source, ") ", rec.Message, "\n")
	}

	// inflight logs flushed, safe to quit
	close(w.flushed)
}

// This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (w *ConsoleLogWriter) LogWrite(rec *LogRecord) {
	w.records <- rec
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
//
// Caution: call LogWrite after Close will panic: send on closed channel
func (w *ConsoleLogWriter) Close() {
	close(w.records)

	// wait for inflight logs flush
	<-w.flushed
}
