// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/list_view_item.proto

package experimentationv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListViewItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListViewItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListViewItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListViewItemMultiError, or
// nil if none found.
func (m *ListViewItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListViewItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListViewItemValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListViewItemValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListViewItemValidationError{
				field:  "Properties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListViewItemMultiError(errors)
	}

	return nil
}

// ListViewItemMultiError is an error wrapping multiple validation errors
// returned by ListViewItem.ValidateAll() if the designated constraints aren't met.
type ListViewItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListViewItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListViewItemMultiError) AllErrors() []error { return m }

// ListViewItemValidationError is the validation error returned by
// ListViewItem.Validate if the designated constraints aren't met.
type ListViewItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListViewItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListViewItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListViewItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListViewItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListViewItemValidationError) ErrorName() string { return "ListViewItemValidationError" }

// Error satisfies the builtin error interface
func (e ListViewItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListViewItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListViewItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListViewItemValidationError{}
