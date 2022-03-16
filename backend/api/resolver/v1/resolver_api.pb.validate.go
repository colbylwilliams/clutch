// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resolver/v1/resolver_api.proto

package resolverv1

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

// Validate checks the field values on AutocompleteResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AutocompleteResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AutocompleteResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AutocompleteResultMultiError, or nil if none found.
func (m *AutocompleteResult) ValidateAll() error {
	return m.validate(true)
}

func (m *AutocompleteResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Label

	if len(errors) > 0 {
		return AutocompleteResultMultiError(errors)
	}

	return nil
}

// AutocompleteResultMultiError is an error wrapping multiple validation errors
// returned by AutocompleteResult.ValidateAll() if the designated constraints
// aren't met.
type AutocompleteResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AutocompleteResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AutocompleteResultMultiError) AllErrors() []error { return m }

// AutocompleteResultValidationError is the validation error returned by
// AutocompleteResult.Validate if the designated constraints aren't met.
type AutocompleteResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AutocompleteResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AutocompleteResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AutocompleteResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AutocompleteResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AutocompleteResultValidationError) ErrorName() string {
	return "AutocompleteResultValidationError"
}

// Error satisfies the builtin error interface
func (e AutocompleteResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAutocompleteResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AutocompleteResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AutocompleteResultValidationError{}

// Validate checks the field values on AutocompleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AutocompleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AutocompleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AutocompleteRequestMultiError, or nil if none found.
func (m *AutocompleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AutocompleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetWant()) < 1 {
		err := AutocompleteRequestValidationError{
			field:  "Want",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetSearch()) < 1 {
		err := AutocompleteRequestValidationError{
			field:  "Search",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Limit

	if len(errors) > 0 {
		return AutocompleteRequestMultiError(errors)
	}

	return nil
}

// AutocompleteRequestMultiError is an error wrapping multiple validation
// errors returned by AutocompleteRequest.ValidateAll() if the designated
// constraints aren't met.
type AutocompleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AutocompleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AutocompleteRequestMultiError) AllErrors() []error { return m }

// AutocompleteRequestValidationError is the validation error returned by
// AutocompleteRequest.Validate if the designated constraints aren't met.
type AutocompleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AutocompleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AutocompleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AutocompleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AutocompleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AutocompleteRequestValidationError) ErrorName() string {
	return "AutocompleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AutocompleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAutocompleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AutocompleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AutocompleteRequestValidationError{}

// Validate checks the field values on AutocompleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AutocompleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AutocompleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AutocompleteResponseMultiError, or nil if none found.
func (m *AutocompleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AutocompleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AutocompleteResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AutocompleteResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AutocompleteResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AutocompleteResponseMultiError(errors)
	}

	return nil
}

// AutocompleteResponseMultiError is an error wrapping multiple validation
// errors returned by AutocompleteResponse.ValidateAll() if the designated
// constraints aren't met.
type AutocompleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AutocompleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AutocompleteResponseMultiError) AllErrors() []error { return m }

// AutocompleteResponseValidationError is the validation error returned by
// AutocompleteResponse.Validate if the designated constraints aren't met.
type AutocompleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AutocompleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AutocompleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AutocompleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AutocompleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AutocompleteResponseValidationError) ErrorName() string {
	return "AutocompleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AutocompleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAutocompleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AutocompleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AutocompleteResponseValidationError{}

// Validate checks the field values on ResolveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResolveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResolveRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResolveRequestMultiError,
// or nil if none found.
func (m *ResolveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ResolveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetWant()) < 1 {
		err := ResolveRequestValidationError{
			field:  "Want",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetHave() == nil {
		err := ResolveRequestValidationError{
			field:  "Have",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetHave(); a != nil {

	}

	// no validation rules for Limit

	if len(errors) > 0 {
		return ResolveRequestMultiError(errors)
	}

	return nil
}

// ResolveRequestMultiError is an error wrapping multiple validation errors
// returned by ResolveRequest.ValidateAll() if the designated constraints
// aren't met.
type ResolveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResolveRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResolveRequestMultiError) AllErrors() []error { return m }

// ResolveRequestValidationError is the validation error returned by
// ResolveRequest.Validate if the designated constraints aren't met.
type ResolveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResolveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResolveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResolveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResolveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResolveRequestValidationError) ErrorName() string { return "ResolveRequestValidationError" }

// Error satisfies the builtin error interface
func (e ResolveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResolveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResolveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResolveRequestValidationError{}

// Validate checks the field values on ResolveResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ResolveResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResolveResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResolveResponseMultiError, or nil if none found.
func (m *ResolveResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ResolveResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResolveResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResolveResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResolveResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPartialFailures() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResolveResponseValidationError{
						field:  fmt.Sprintf("PartialFailures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResolveResponseValidationError{
						field:  fmt.Sprintf("PartialFailures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResolveResponseValidationError{
					field:  fmt.Sprintf("PartialFailures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ResolveResponseMultiError(errors)
	}

	return nil
}

// ResolveResponseMultiError is an error wrapping multiple validation errors
// returned by ResolveResponse.ValidateAll() if the designated constraints
// aren't met.
type ResolveResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResolveResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResolveResponseMultiError) AllErrors() []error { return m }

// ResolveResponseValidationError is the validation error returned by
// ResolveResponse.Validate if the designated constraints aren't met.
type ResolveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResolveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResolveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResolveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResolveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResolveResponseValidationError) ErrorName() string { return "ResolveResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResolveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResolveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResolveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResolveResponseValidationError{}

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchRequestMultiError, or
// nil if none found.
func (m *SearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetWant()) < 1 {
		err := SearchRequestValidationError{
			field:  "Want",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetQuery()) < 1 {
		err := SearchRequestValidationError{
			field:  "Query",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Limit

	if len(errors) > 0 {
		return SearchRequestMultiError(errors)
	}

	return nil
}

// SearchRequestMultiError is an error wrapping multiple validation errors
// returned by SearchRequest.ValidateAll() if the designated constraints
// aren't met.
type SearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchRequestMultiError) AllErrors() []error { return m }

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

// Validate checks the field values on SearchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchResponseMultiError,
// or nil if none found.
func (m *SearchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPartialFailures() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("PartialFailures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("PartialFailures[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  fmt.Sprintf("PartialFailures[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SearchResponseMultiError(errors)
	}

	return nil
}

// SearchResponseMultiError is an error wrapping multiple validation errors
// returned by SearchResponse.ValidateAll() if the designated constraints
// aren't met.
type SearchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchResponseMultiError) AllErrors() []error { return m }

// SearchResponseValidationError is the validation error returned by
// SearchResponse.Validate if the designated constraints aren't met.
type SearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchResponseValidationError) ErrorName() string { return "SearchResponseValidationError" }

// Error satisfies the builtin error interface
func (e SearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchResponseValidationError{}

// Validate checks the field values on GetObjectSchemasRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetObjectSchemasRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetObjectSchemasRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetObjectSchemasRequestMultiError, or nil if none found.
func (m *GetObjectSchemasRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetObjectSchemasRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTypeUrl()) < 1 {
		err := GetObjectSchemasRequestValidationError{
			field:  "TypeUrl",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetObjectSchemasRequestMultiError(errors)
	}

	return nil
}

// GetObjectSchemasRequestMultiError is an error wrapping multiple validation
// errors returned by GetObjectSchemasRequest.ValidateAll() if the designated
// constraints aren't met.
type GetObjectSchemasRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetObjectSchemasRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetObjectSchemasRequestMultiError) AllErrors() []error { return m }

// GetObjectSchemasRequestValidationError is the validation error returned by
// GetObjectSchemasRequest.Validate if the designated constraints aren't met.
type GetObjectSchemasRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectSchemasRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectSchemasRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectSchemasRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectSchemasRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectSchemasRequestValidationError) ErrorName() string {
	return "GetObjectSchemasRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectSchemasRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectSchemasRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectSchemasRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectSchemasRequestValidationError{}

// Validate checks the field values on GetObjectSchemasResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetObjectSchemasResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetObjectSchemasResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetObjectSchemasResponseMultiError, or nil if none found.
func (m *GetObjectSchemasResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetObjectSchemasResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TypeUrl

	for idx, item := range m.GetSchemas() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetObjectSchemasResponseValidationError{
						field:  fmt.Sprintf("Schemas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetObjectSchemasResponseValidationError{
						field:  fmt.Sprintf("Schemas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetObjectSchemasResponseValidationError{
					field:  fmt.Sprintf("Schemas[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetObjectSchemasResponseMultiError(errors)
	}

	return nil
}

// GetObjectSchemasResponseMultiError is an error wrapping multiple validation
// errors returned by GetObjectSchemasResponse.ValidateAll() if the designated
// constraints aren't met.
type GetObjectSchemasResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetObjectSchemasResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetObjectSchemasResponseMultiError) AllErrors() []error { return m }

// GetObjectSchemasResponseValidationError is the validation error returned by
// GetObjectSchemasResponse.Validate if the designated constraints aren't met.
type GetObjectSchemasResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectSchemasResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectSchemasResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectSchemasResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectSchemasResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectSchemasResponseValidationError) ErrorName() string {
	return "GetObjectSchemasResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectSchemasResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectSchemasResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectSchemasResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectSchemasResponseValidationError{}
