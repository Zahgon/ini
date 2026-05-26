// Copyright 2017 Unknwon
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
	"bytes"
	"io"
	"sync"
)

// File represents a combination of one or more INI files in memory.
type File struct {
	options     LoadOptions
	dataSources []dataSource

	// Should make things safe, but sometimes doesn't matter.
	BlockMode bool
	lock      sync.RWMutex

	// To keep data in order.
	sectionList []string
	// To keep track of the index of a section with same name.
	// This meta list is only used with non-unique section names are allowed.
	sectionIndexes []int

	// Actual data is stored here.
	sections map[string][]*Section

	NameMapper
	ValueMapper
}

// newFile initializes File object with given data sources.
func newFile(dataSources []dataSource, opts LoadOptions) *File {
	_ = "STUB: not implemented"
	return nil
}

// Empty returns an empty file object.
func Empty(opts ...LoadOptions) *File { _ = "STUB: not implemented"; return nil }

// Ignore error here, we are sure our data is good.

// NewSection creates a new section.
func (f *File) NewSection(name string) (*Section, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: Append to indexes must happen before appending to sections,
// otherwise index will have off-by-one problem.

// NewRawSection creates a new section with an unparseable body.
func (f *File) NewRawSection(name, body string) (*Section, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSections creates a list of sections.
func (f *File) NewSections(names ...string) (err error) { _ = "STUB: not implemented"; return nil }

// GetSection returns section by given name.
func (f *File) GetSection(name string) (*Section, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HasSection returns true if the file contains a section with given name.
func (f *File) HasSection(name string) bool { _ = "STUB: not implemented"; return false }

// SectionsByName returns all sections with given name.
func (f *File) SectionsByName(name string) ([]*Section, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Section assumes named section exists and returns a zero-value when not.
func (f *File) Section(name string) *Section { _ = "STUB: not implemented"; return nil }

// SectionWithIndex assumes named section exists and returns a new section when not.
func (f *File) SectionWithIndex(name string, index int) *Section {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: It's OK here because the only possible error is empty section name,
// but if it's empty, this piece of code won't be executed.

// Sections returns a list of Section stored in the current instance.
func (f *File) Sections() []*Section { _ = "STUB: not implemented"; return nil }

// ChildSections returns a list of child sections of given section name.
func (f *File) ChildSections(name string) []*Section { _ = "STUB: not implemented"; return nil }

// SectionStrings returns list of section names.
func (f *File) SectionStrings() []string { _ = "STUB: not implemented"; return nil }

// DeleteSection deletes a section or all sections with given name.
func (f *File) DeleteSection(name string) { _ = "STUB: not implemented"; return }

// For non-unique sections, it is always needed to remove the first one so
// in the next iteration, the subsequent section continue having index 0.
// Ignoring the error as index 0 never returns an error.

// DeleteSectionWithIndex deletes a section with given name and index.
func (f *File) DeleteSectionWithIndex(name string, index int) error {
	_ = "STUB: not implemented"
	return nil
}

// Count occurrences of the sections

// The last one in the map

// Fix section lists

// Fix the indices of all following sections with this name.

func (f *File) reload(s dataSource) error { _ = "STUB: not implemented"; return nil }

// Reload reloads and parses all data sources.
func (f *File) Reload() (err error) { _ = "STUB: not implemented"; return nil }

// In loose mode, we create an empty default section for nonexistent files.

// Append appends one or more data sources and reloads automatically.
func (f *File) Append(source interface{}, others ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) writeToBuffer(indent string) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use buffer to make sure target is safe until finish encoding.

// Support multiline comments

// Write nothing if default section is empty

// Put a line between sections

// Count and generate alignment length and buffer spaces using the
// longest key. Keys may be modified if they contain certain characters so
// we need to take that into account in our calculation.

// First case will surround key by ` and second by """

// Support multiline comments

// Write out alignment spaces before "=" sign

// In case key value contains "\n", "`", "\"", "#" or ";"

// Put a line between sections

// WriteToIndent writes content into io.Writer with given indention.
// If PrettyFormat has been set to be true,
// it will align "=" sign with spaces under each section.
func (f *File) WriteToIndent(w io.Writer, indent string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteTo writes file content into io.Writer.
func (f *File) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// SaveToIndent writes content to file system with given value indention.
func (f *File) SaveToIndent(filename, indent string) error {
	_ = "STUB: not implemented"
	// Note: Because we are truncating with os.Create,
	//
	//	so it's safer to save to a temporary file location and rename after done.
	return nil
}

// SaveTo writes content to file system.
func (f *File) SaveTo(filename string) error { _ = "STUB: not implemented"; return nil }
