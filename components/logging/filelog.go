// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package logging

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// This log writer sends output to a file
type FileLogWriter struct {
	rec chan *LogRecord

	// The opened file
	filename string
	curfile  string
	file     *os.File

	// The logging format
	format string

	// Rotate at size
	maxsize int64
	cursize int64

	// File creating time
	opentime time.Time

	// Current suffix number
	curnum int

	rotate bool

	// Max duration of a single log, create a new log when exceeding.
	// DHOURLY: one hour; DDAILY: one day; DPERMANT: permant.
	duration byte
}

// This is the FileLogWriter's output method
func (w *FileLogWriter) LogWrite(rec *LogRecord) {
	w.rec <- rec
}

func (w *FileLogWriter) Close() {
	close(w.rec)
	w.file.Sync()
}

// NewFileLogWriter creates a new LogWriter which writes to the given file and
// has rotation enabled if rotate is true.
//
// If rotate is true, any time a new log file is opened, the old one is renamed
// with a .### extension to preserve it.  The various Set* methods can be used
// to configure log rotation based on lines, size, and daily.
//
// The standard log-line format is:
//   [%D %T] [%L] (%S) %M
func NewFileLogWriter(fname string, format_ string,
	maxsize_ int64, duration_ byte, rot bool) *FileLogWriter {
	w := &FileLogWriter{
		rec:      make(chan *LogRecord, LogBufferLength),
		filename: fname,
		file:     nil,
		format:   format_,
		maxsize:  maxsize_,
		cursize:  0,
		curnum:   0,
		rotate:   rot,
		duration: duration_,
	}

	// open the file for the first time
	if err := w.intRotate(); err != nil {
		fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.filename, err)
		return nil
	}

	go func() {
		defer func() {
			if w.file != nil {
				w.file.Close()
			}
		}()

		for {
			rec := <-w.rec
			now := time.Now()
			if (now.Hour() != w.opentime.Hour() && w.duration == DHOURLY) ||
				(now.Day() != w.opentime.Day() && w.duration == DDAILY) ||
				(w.maxsize > 0 && w.cursize >= w.maxsize) {
				if err := w.intRotate(); err != nil {
					fmt.Fprintf(os.Stderr, "[EMERGE] FileLogWriter(%q): %s\n", w.filename, err)
					return
				}
			}

			// Perform the write
			n, err := fmt.Fprint(w.file, FormatLogRecord(w.format, rec))
			if err != nil {
				fmt.Fprintf(os.Stderr, "[EMERGE] FileLogWriter(%q): %s\n", w.filename, err)
				return
			}

			// Update the counts
			w.cursize += int64(n)
		}
	}()

	return w
}

// If this is called in a threaded context, it MUST be synchronized
func (w *FileLogWriter) intRotate() error {
	now := time.Now()

	// First open
	if w.file == nil {
		var i int
		var tempfile string
		for i = 0; i < 1000; i++ {
			tempfile = getFilename(w.filename, w.duration, now, i)
			_, err := os.Lstat(tempfile)
			if err != nil {
				break
			}
		}

		if i != 0 {
			i--
		}
		w.opentime = now
		w.curnum = i
	} else if w.maxsize > 0 && w.cursize >= w.maxsize {
		if w.curnum == 999 {
			return errors.New("Log files in current hour reach max number 1000")
		}

		w.file.Close()
		w.curnum++
	} else if now.Hour() != w.opentime.Hour() || now.Day() != w.opentime.Day() {
		w.file.Close()
		w.opentime = now
		w.curnum = 0
	} else {
		return errors.New("intRotate unknown status")
	}

	// Open the log file
	w.curfile = getFilename(w.filename, w.duration, now, w.curnum)
	fd, err := os.OpenFile(w.curfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	w.file = fd

	// initialize rotation values
	info, err := fd.Stat()
	if err != nil {
		return err
	}
	w.cursize = info.Size()

	return nil
}

// Set the logging format (chainable).  Must be called before the first log
// message is written.
func (w *FileLogWriter) SetFormat(format string) *FileLogWriter {
	w.format = format
	return w
}

// Set rotate at size (chainable). Must be called before the first log message
// is written.
func (w *FileLogWriter) SetRotateSize(maxsize int64) *FileLogWriter {
	//fmt.Fprintf(os.Stderr, "FileLogWriter.SetRotateSize: %v\n", maxsize)
	w.maxsize = maxsize
	return w
}

// Set duration. Must be called before the first log message is written.
func (w *FileLogWriter) SetDuration(duration byte) *FileLogWriter {
	fmt.Fprintf(os.Stderr, "FileLogWriter.SetDuration: %v\n", duration)
	w.duration = duration
	return w
}

// NewXMLLogWriter is a utility method for creating a FileLogWriter set up to
// output XML record log messages instead of line-based ones.
func NewXMLLogWriter(fname string, maxsize int64, duration byte, rotate bool) *FileLogWriter {
	format := `<record level="%L">
               <timestamp>%D %T</timestamp>
               <source>%S</source>
               <message>%M</message>
	       </record>`
	return NewFileLogWriter(fname, format, maxsize, duration, rotate)
}

func getFilename(filename string, duration byte, now time.Time, suffix int) string {
	if duration == DHOURLY {
		return filename + fmt.Sprintf(".%d%02d%02d%02d",
			now.Year(), now.Month(), now.Day(), now.Hour())
	} else if duration == DDAILY {
		return filename + fmt.Sprintf(".%d%02d%02d",
			now.Year(), now.Month(), now.Day())
	} else if duration == DFILENAME {
		return filename //only filename
	} else {
		return filename + fmt.Sprintf(".%03d", suffix)
	}
}
