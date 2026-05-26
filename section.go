// Copyright 2014 Unknwon
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

// Section represents a config section.
type Section struct {
	f        *File
	Comment  string
	name     string
	keys     map[string]*Key
	keyList  []string
	keysHash map[string]string

	isRawSection bool
	rawBody      string
}

func newSection(f *File, name string) *Section { _ = "STUB: not implemented"; return nil }

// Name returns name of Section.
func (s *Section) Name() string {
	_ = "STUB: not implemented"

	// Body returns rawBody of Section if the section was marked as unparseable.
	// It still follows the other rules of the INI format surrounding leading/trailing whitespace.
	return ""
}

func (s *Section) Body() string { _ = "STUB: not implemented"; return "" }

// SetBody updates body content only if section is raw.
func (s *Section) SetBody(body string) { _ = "STUB: not implemented"; return }

// NewKey creates a new key to given section.
func (s *Section) NewKey(name, val string) (*Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBooleanKey creates a new boolean type key to given section.
func (s *Section) NewBooleanKey(name string) (*Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetKey returns key in section by given name.
func (s *Section) GetKey(name string) (*Key, error) { _ = "STUB: not implemented"; return nil, nil }

// Check if it is a child-section.

// HasKey returns true if section contains a key with given name.
func (s *Section) HasKey(name string) bool { _ = "STUB: not implemented"; return false }

// Deprecated: Use "HasKey" instead.
func (s *Section) Haskey(name string) bool { _ = "STUB: not implemented"; return false }

// HasValue returns true if section contains given raw value.
func (s *Section) HasValue(value string) bool { _ = "STUB: not implemented"; return false }

// Key assumes named Key exists in section and returns a zero-value when not.
func (s *Section) Key(name string) *Key { _ = "STUB: not implemented"; return nil }

// It's OK here because the only possible error is empty key name,
// but if it's empty, this piece of code won't be executed.

// Keys returns list of keys of section.
func (s *Section) Keys() []*Key { _ = "STUB: not implemented"; return nil }

// ParentKeys returns list of keys of parent section.
func (s *Section) ParentKeys() []*Key { _ = "STUB: not implemented"; return nil }

// KeyStrings returns list of key names of section.
func (s *Section) KeyStrings() []string { _ = "STUB: not implemented"; return nil }

// KeysHash returns keys hash consisting of names and values.
func (s *Section) KeysHash() map[string]string { _ = "STUB: not implemented"; return nil }

// DeleteKey deletes a key from section.
func (s *Section) DeleteKey(name string) { _ = "STUB: not implemented"; return }

// ChildSections returns a list of child sections of current section.
// For example, "[parent.child1]" and "[parent.child12]" are child sections
// of section "[parent]".
func (s *Section) ChildSections() []*Section { _ = "STUB: not implemented"; return nil }
