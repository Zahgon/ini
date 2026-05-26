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

import (
	"time"
)

// Key represents a key under a section.
type Key struct {
	s               *Section
	Comment         string
	name            string
	value           string
	isAutoIncrement bool
	isBooleanType   bool

	isShadow bool
	shadows  []*Key

	nestedValues []string
}

// newKey simply return a key object with given values.
func newKey(s *Section, name, val string) *Key { _ = "STUB: not implemented"; return nil }

func (k *Key) addShadow(val string) error { _ = "STUB: not implemented"; return nil }

// Deduplicate shadows based on their values.

// AddShadow adds a new shadow key to itself.
func (k *Key) AddShadow(val string) error { _ = "STUB: not implemented"; return nil }

func (k *Key) addNestedValue(val string) error { _ = "STUB: not implemented"; return nil }

// AddNestedValue adds a nested value to the key.
func (k *Key) AddNestedValue(val string) error { _ = "STUB: not implemented"; return nil }

// ValueMapper represents a mapping function for values, e.g. os.ExpandEnv
type ValueMapper func(string) string

// Name returns name of key.
func (k *Key) Name() string {
	_ = "STUB: not implemented"

	// Value returns raw value of key for performance purpose.
	return ""
}

func (k *Key) Value() string {
	_ = "STUB: not implemented"

	// ValueWithShadows returns raw values of key and its shadows if any. Shadow
	// keys with empty values are ignored from the returned list.
	return ""
}

func (k *Key) ValueWithShadows() []string { _ = "STUB: not implemented"; return nil }

// NestedValues returns nested values stored in the key.
// It is possible returned value is nil if no nested values stored in the key.
func (k *Key) NestedValues() []string { _ = "STUB: not implemented"; return nil }

// transformValue takes a raw value and transforms to its final string.
func (k *Key) transformValue(val string) string { _ = "STUB: not implemented"; return "" }

// Fail-fast if no indicate char found for recursive value

// Take off leading '%(' and trailing ')s'.

// Search in the same section.
// If not found or found the key itself, then search again in default section.

// Stop when no results found in the default section,
// and returns the value as-is.

// Substitute by new value and take off leading '%(' and trailing ')s'.

// String returns string representation of value.
func (k *Key) String() string { _ = "STUB: not implemented"; return "" }

// Validate accepts a validate function which can
// return modifed result as key value.
func (k *Key) Validate(fn func(string) string) string { _ = "STUB: not implemented"; return "" }

// parseBool returns the boolean value represented by the string.
//
// It accepts 1, t, T, TRUE, true, True, YES, yes, Yes, y, ON, on, On,
// 0, f, F, FALSE, false, False, NO, no, No, n, OFF, off, Off.
// Any other value returns an error.
func parseBool(str string) (value bool, err error) { _ = "STUB: not implemented"; return false, nil }

// Bool returns bool type value.
func (k *Key) Bool() (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Float64 returns float64 type value.
		nil
}

func (k *Key) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Int returns int type value.
func (k *Key) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64 returns int64 type value.
func (k *Key) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint returns uint type valued.
func (k *Key) Uint() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint64 returns uint64 type value.
func (k *Key) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Duration returns time.Duration type value.
func (k *Key) Duration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// TimeFormat parses with given format and returns time.Time type value.
func (k *Key) TimeFormat(format string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Time parses with RFC3339 format and returns time.Time type value.
func (k *Key) Time() (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }

// MustString returns default value if key value is empty.
func (k *Key) MustString(defaultVal string) string { _ = "STUB: not implemented"; return "" }

// MustBool always returns value without error,
// it returns false if error occurs.
func (k *Key) MustBool(defaultVal ...bool) bool { _ = "STUB: not implemented"; return false }

// MustFloat64 always returns value without error,
// it returns 0.0 if error occurs.
func (k *Key) MustFloat64(defaultVal ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// MustInt always returns value without error,
// it returns 0 if error occurs.
func (k *Key) MustInt(defaultVal ...int) int { _ = "STUB: not implemented"; return 0 }

// MustInt64 always returns value without error,
// it returns 0 if error occurs.
func (k *Key) MustInt64(defaultVal ...int64) int64 { _ = "STUB: not implemented"; return 0 }

// MustUint always returns value without error,
// it returns 0 if error occurs.
func (k *Key) MustUint(defaultVal ...uint) uint { _ = "STUB: not implemented"; return 0 }

// MustUint64 always returns value without error,
// it returns 0 if error occurs.
func (k *Key) MustUint64(defaultVal ...uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// MustDuration always returns value without error,
// it returns zero value if error occurs.
func (k *Key) MustDuration(defaultVal ...time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// MustTimeFormat always parses with given format and returns value without error,
// it returns zero value if error occurs.
func (k *Key) MustTimeFormat(format string, defaultVal ...time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// MustTime always parses with RFC3339 format and returns value without error,
// it returns zero value if error occurs.
func (k *Key) MustTime(defaultVal ...time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// In always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) In(defaultVal string, candidates []string) string {
	_ = "STUB: not implemented"
	return ""
}

// InFloat64 always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InFloat64(defaultVal float64, candidates []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// InInt always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InInt(defaultVal int, candidates []int) int { _ = "STUB: not implemented"; return 0 }

// InInt64 always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InInt64(defaultVal int64, candidates []int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// InUint always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InUint(defaultVal uint, candidates []uint) uint { _ = "STUB: not implemented"; return 0 }

// InUint64 always returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InUint64(defaultVal uint64, candidates []uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// InTimeFormat always parses with given format and returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InTimeFormat(format string, defaultVal time.Time, candidates []time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// InTime always parses with RFC3339 format and returns value without error,
// it returns default value if error occurs or doesn't fit into candidates.
func (k *Key) InTime(defaultVal time.Time, candidates []time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// RangeFloat64 checks if value is in given range inclusively,
// and returns default value if it's not.
func (k *Key) RangeFloat64(defaultVal, min, max float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// RangeInt checks if value is in given range inclusively,
// and returns default value if it's not.
func (k *Key) RangeInt(defaultVal, min, max int) int { _ = "STUB: not implemented"; return 0 }

// RangeInt64 checks if value is in given range inclusively,
// and returns default value if it's not.
func (k *Key) RangeInt64(defaultVal, min, max int64) int64 { _ = "STUB: not implemented"; return 0 }

// RangeTimeFormat checks if value with given format is in given range inclusively,
// and returns default value if it's not.
func (k *Key) RangeTimeFormat(format string, defaultVal, min, max time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// RangeTime checks if value with RFC3339 format is in given range inclusively,
// and returns default value if it's not.
func (k *Key) RangeTime(defaultVal, min, max time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Strings returns list of string divided by given delimiter.
func (k *Key) Strings(delim string) []string { _ = "STUB: not implemented"; return nil }

// StringsWithShadows returns list of string divided by given delimiter.
// Shadows will also be appended if any.
func (k *Key) StringsWithShadows(delim string) []string { _ = "STUB: not implemented"; return nil }

// Float64s returns list of float64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Float64s(delim string) []float64 { _ = "STUB: not implemented"; return nil }

// Ints returns list of int divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Ints(delim string) []int { _ = "STUB: not implemented"; return nil }

// Int64s returns list of int64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Int64s(delim string) []int64 { _ = "STUB: not implemented"; return nil }

// Uints returns list of uint divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Uints(delim string) []uint { _ = "STUB: not implemented"; return nil }

// Uint64s returns list of uint64 divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Uint64s(delim string) []uint64 { _ = "STUB: not implemented"; return nil }

// Bools returns list of bool divided by given delimiter. Any invalid input will be treated as zero value.
func (k *Key) Bools(delim string) []bool { _ = "STUB: not implemented"; return nil }

// TimesFormat parses with given format and returns list of time.Time divided by given delimiter.
// Any invalid input will be treated as zero value (0001-01-01 00:00:00 +0000 UTC).
func (k *Key) TimesFormat(format, delim string) []time.Time { _ = "STUB: not implemented"; return nil }

// Times parses with RFC3339 format and returns list of time.Time divided by given delimiter.
// Any invalid input will be treated as zero value (0001-01-01 00:00:00 +0000 UTC).
func (k *Key) Times(delim string) []time.Time { _ = "STUB: not implemented"; return nil }

// ValidFloat64s returns list of float64 divided by given delimiter. If some value is not float, then
// it will not be included to result list.
func (k *Key) ValidFloat64s(delim string) []float64 { _ = "STUB: not implemented"; return nil }

// ValidInts returns list of int divided by given delimiter. If some value is not integer, then it will
// not be included to result list.
func (k *Key) ValidInts(delim string) []int { _ = "STUB: not implemented"; return nil }

// ValidInt64s returns list of int64 divided by given delimiter. If some value is not 64-bit integer,
// then it will not be included to result list.
func (k *Key) ValidInt64s(delim string) []int64 { _ = "STUB: not implemented"; return nil }

// ValidUints returns list of uint divided by given delimiter. If some value is not unsigned integer,
// then it will not be included to result list.
func (k *Key) ValidUints(delim string) []uint { _ = "STUB: not implemented"; return nil }

// ValidUint64s returns list of uint64 divided by given delimiter. If some value is not 64-bit unsigned
// integer, then it will not be included to result list.
func (k *Key) ValidUint64s(delim string) []uint64 { _ = "STUB: not implemented"; return nil }

// ValidBools returns list of bool divided by given delimiter. If some value is not 64-bit unsigned
// integer, then it will not be included to result list.
func (k *Key) ValidBools(delim string) []bool { _ = "STUB: not implemented"; return nil }

// ValidTimesFormat parses with given format and returns list of time.Time divided by given delimiter.
func (k *Key) ValidTimesFormat(format, delim string) []time.Time {
	_ = "STUB: not implemented"
	return nil
}

// ValidTimes parses with RFC3339 format and returns list of time.Time divided by given delimiter.
func (k *Key) ValidTimes(delim string) []time.Time { _ = "STUB: not implemented"; return nil }

// StrictFloat64s returns list of float64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictFloat64s(delim string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StrictInts returns list of int divided by given delimiter or error on first invalid input.
func (k *Key) StrictInts(delim string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

// StrictInt64s returns list of int64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictInt64s(delim string) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StrictUints returns list of uint divided by given delimiter or error on first invalid input.
func (k *Key) StrictUints(delim string) ([]uint, error) { _ = "STUB: not implemented"; return nil, nil }

// StrictUint64s returns list of uint64 divided by given delimiter or error on first invalid input.
func (k *Key) StrictUint64s(delim string) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StrictBools returns list of bool divided by given delimiter or error on first invalid input.
func (k *Key) StrictBools(delim string) ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// StrictTimesFormat parses with given format and returns list of time.Time divided by given delimiter
// or error on first invalid input.
func (k *Key) StrictTimesFormat(format, delim string) ([]time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StrictTimes parses with RFC3339 format and returns list of time.Time divided by given delimiter
// or error on first invalid input.
func (k *Key) StrictTimes(delim string) ([]time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseBools transforms strings to bools.
func (k *Key) parseBools(strs []string, addInvalid, returnOnInvalid bool) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseFloat64s transforms strings to float64s.
func (k *Key) parseFloat64s(strs []string, addInvalid, returnOnInvalid bool) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseInts transforms strings to ints.
func (k *Key) parseInts(strs []string, addInvalid, returnOnInvalid bool) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseInt64s transforms strings to int64s.
func (k *Key) parseInt64s(strs []string, addInvalid, returnOnInvalid bool) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseUints transforms strings to uints.
func (k *Key) parseUints(strs []string, addInvalid, returnOnInvalid bool) ([]uint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseUint64s transforms strings to uint64s.
func (k *Key) parseUint64s(strs []string, addInvalid, returnOnInvalid bool) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Parser func(str string) (interface{}, error)

// parseTimesFormat transforms strings to times in given format.
func (k *Key) parseTimesFormat(format string, strs []string, addInvalid, returnOnInvalid bool) ([]time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// doParse transforms strings to different types
func (k *Key) doParse(strs []string, addInvalid, returnOnInvalid bool, parser Parser) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetValue changes key value.
func (k *Key) SetValue(v string) { _ = "STUB: not implemented"; return }
