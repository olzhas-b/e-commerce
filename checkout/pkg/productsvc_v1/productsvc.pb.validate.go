// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: productsvc.proto

package productsvc_v1

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

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductRequestMultiError, or nil if none found.
func (m *GetProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Sku

	if len(errors) > 0 {
		return GetProductRequestMultiError(errors)
	}

	return nil
}

// GetProductRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductRequestMultiError) AllErrors() []error { return m }

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on GetProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductResponseMultiError, or nil if none found.
func (m *GetProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Price

	if len(errors) > 0 {
		return GetProductResponseMultiError(errors)
	}

	return nil
}

// GetProductResponseMultiError is an error wrapping multiple validation errors
// returned by GetProductResponse.ValidateAll() if the designated constraints
// aren't met.
type GetProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductResponseMultiError) AllErrors() []error { return m }

// GetProductResponseValidationError is the validation error returned by
// GetProductResponse.Validate if the designated constraints aren't met.
type GetProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductResponseValidationError) ErrorName() string {
	return "GetProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductResponseValidationError{}

// Validate checks the field values on ListSkusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSkusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSkusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSkusRequestMultiError, or nil if none found.
func (m *ListSkusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSkusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for StartAfterSku

	// no validation rules for Count

	if len(errors) > 0 {
		return ListSkusRequestMultiError(errors)
	}

	return nil
}

// ListSkusRequestMultiError is an error wrapping multiple validation errors
// returned by ListSkusRequest.ValidateAll() if the designated constraints
// aren't met.
type ListSkusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSkusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSkusRequestMultiError) AllErrors() []error { return m }

// ListSkusRequestValidationError is the validation error returned by
// ListSkusRequest.Validate if the designated constraints aren't met.
type ListSkusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkusRequestValidationError) ErrorName() string { return "ListSkusRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListSkusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkusRequestValidationError{}

// Validate checks the field values on ListSkusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSkusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSkusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSkusResponseMultiError, or nil if none found.
func (m *ListSkusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSkusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListSkusResponseMultiError(errors)
	}

	return nil
}

// ListSkusResponseMultiError is an error wrapping multiple validation errors
// returned by ListSkusResponse.ValidateAll() if the designated constraints
// aren't met.
type ListSkusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSkusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSkusResponseMultiError) AllErrors() []error { return m }

// ListSkusResponseValidationError is the validation error returned by
// ListSkusResponse.Validate if the designated constraints aren't met.
type ListSkusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkusResponseValidationError) ErrorName() string { return "ListSkusResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListSkusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkusResponseValidationError{}
