// Copyright 2015 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package ini

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
)

const minReaderBufferSize = 4096

var pythonMultiline = regexp.MustCompile(`^([\t\f ]+)(.*)`)

type parserOptions struct {
	IgnoreContinuation          bool
	IgnoreInlineComment         bool
	AllowPythonMultilineValues  bool
	SpaceBeforeInlineComment    bool
	UnescapeValueDoubleQuotes   bool
	UnescapeValueCommentSymbols bool
	PreserveSurroundedQuote     bool
	DebugFunc                   DebugFunc
	ReaderBufferSize            int
}

type parser struct {
	buf     *bufio.Reader
	options parserOptions

	isEOF   bool
	count   int
	comment *bytes.Buffer
}

func (p *parser) debug(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func newParser(r io.Reader, opts parserOptions) *parser { _ = "STUB: not implemented"; return nil }

// BOM handles header of UTF-8, UTF-16 LE and UTF-16 BE's BOM format.
// http://en.wikipedia.org/wiki/Byte_order_mark#Representations_of_byte_order_marks_by_encoding
func (p *parser) BOM() error { _ = "STUB: not implemented"; return nil }

func (p *parser) readUntil(delim byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func cleanComment(in []byte) ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func readKeyName(delimiters string, in []byte) (string, int, error) {
	_ = "STUB: not implemented"

	// Check if key name surrounded by quotes.
	return "", 0, nil
}

// Get out key name

// FIXME: fail case -> """"""name"""=value

// Find key-value delimiter

func (p *parser) readMultilines(line, val, valQuote string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if the line ends with backslash continuation after the quote

func (p *parser) readContinuationLines(val string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// hasSurroundedQuote check if and only if the first and last characters
// are quotes \" or \'.
// It returns false if any other parts also contain same kind of quotes.
func hasSurroundedQuote(in string, quote byte) bool { _ = "STUB: not implemented"; return false }

func (p *parser) readValue(in []byte, bufferSize int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check for multi-line value

// Won't be able to reach here if value only contains whitespace

// Check continuation lines when desired

// Check if ignore inline comment

// Trim single and double quotes

func (p *parser) readPythonMultilines(line string, bufferSize int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Return if not a Python multiline value.

// Advance the parser reader (buffer) in-sync with the peek buffer.

// parse parses data through an io.Reader.
func (f *File) parse(reader io.Reader) (err error) { _ = "STUB: not implemented"; return nil }

// Ignore error because default section name is never empty string.

// This "last" is not strictly equivalent to "previous one" if current key is not the first nested key

// NOTE: Iterate and increase `currentPeekSize` until
// the size of the parser buffer is found.
// TODO(unknwon): When Golang 1.10 is the lowest version supported, replace with `parserBufferSize := p.buf.Size()`.

// NOTE: Peek 4kb at a time.

// Comments

// Note: we do not care ending line break,
// it is needed for adding second line,
// so just clean it once at the end when set to value.

// Section

// Read to the next ']' (TODO: support quoted strings)

// Reset auto-counter and comments

// Nested values can't span sections

// Treat as boolean key when desired, and whole line is key name.

// Auto increment.
