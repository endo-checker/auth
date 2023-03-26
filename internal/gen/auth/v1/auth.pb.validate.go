// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/auth.proto

package authv1

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

// Validate checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAccountRequestMultiError, or nil if none found.
func (m *CreateAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRegisterAuthUser() == nil {
		err := CreateAccountRequestValidationError{
			field:  "RegisterAuthUser",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRegisterAuthUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAccountRequestValidationError{
					field:  "RegisterAuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAccountRequestValidationError{
					field:  "RegisterAuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegisterAuthUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAccountRequestValidationError{
				field:  "RegisterAuthUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateAccountRequestMultiError(errors)
	}

	return nil
}

// CreateAccountRequestMultiError is an error wrapping multiple validation
// errors returned by CreateAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAccountRequestMultiError) AllErrors() []error { return m }

// CreateAccountRequestValidationError is the validation error returned by
// CreateAccountRequest.Validate if the designated constraints aren't met.
type CreateAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountRequestValidationError) ErrorName() string {
	return "CreateAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountRequestValidationError{}

// Validate checks the field values on CreateAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAccountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAccountResponseMultiError, or nil if none found.
func (m *CreateAccountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAccountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if all {
		switch v := interface{}(m.GetRegisterAuthUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAccountResponseValidationError{
					field:  "RegisterAuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAccountResponseValidationError{
					field:  "RegisterAuthUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegisterAuthUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAccountResponseValidationError{
				field:  "RegisterAuthUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateAccountResponseMultiError(errors)
	}

	return nil
}

// CreateAccountResponseMultiError is an error wrapping multiple validation
// errors returned by CreateAccountResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateAccountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAccountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAccountResponseMultiError) AllErrors() []error { return m }

// CreateAccountResponseValidationError is the validation error returned by
// CreateAccountResponse.Validate if the designated constraints aren't met.
type CreateAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountResponseValidationError) ErrorName() string {
	return "CreateAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountResponseValidationError{}

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInRequestMultiError, or
// nil if none found.
func (m *SignInRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAuthUserSignIn() == nil {
		err := SignInRequestValidationError{
			field:  "AuthUserSignIn",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAuthUserSignIn()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SignInRequestValidationError{
					field:  "AuthUserSignIn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SignInRequestValidationError{
					field:  "AuthUserSignIn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthUserSignIn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignInRequestValidationError{
				field:  "AuthUserSignIn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SignInRequestMultiError(errors)
	}

	return nil
}

// SignInRequestMultiError is an error wrapping multiple validation errors
// returned by SignInRequest.ValidateAll() if the designated constraints
// aren't met.
type SignInRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInRequestMultiError) AllErrors() []error { return m }

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

// Validate checks the field values on SignInResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInResponseMultiError,
// or nil if none found.
func (m *SignInResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for IdToken

	// no validation rules for Scope

	// no validation rules for TokenType

	// no validation rules for ExpiresIn

	if len(errors) > 0 {
		return SignInResponseMultiError(errors)
	}

	return nil
}

// SignInResponseMultiError is an error wrapping multiple validation errors
// returned by SignInResponse.ValidateAll() if the designated constraints
// aren't met.
type SignInResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInResponseMultiError) AllErrors() []error { return m }

// SignInResponseValidationError is the validation error returned by
// SignInResponse.Validate if the designated constraints aren't met.
type SignInResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInResponseValidationError) ErrorName() string { return "SignInResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignInResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInResponseValidationError{}

// Validate checks the field values on GetAccountRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountRequestMultiError, or nil if none found.
func (m *GetAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return GetAccountRequestMultiError(errors)
	}

	return nil
}

// GetAccountRequestMultiError is an error wrapping multiple validation errors
// returned by GetAccountRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountRequestMultiError) AllErrors() []error { return m }

// GetAccountRequestValidationError is the validation error returned by
// GetAccountRequest.Validate if the designated constraints aren't met.
type GetAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountRequestValidationError) ErrorName() string {
	return "GetAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountRequestValidationError{}

// Validate checks the field values on GetAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAccountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountResponseMultiError, or nil if none found.
func (m *GetAccountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAccountResponseValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAccountResponseValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAccountResponseValidationError{
				field:  "UserInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetAccountResponseMultiError(errors)
	}

	return nil
}

// GetAccountResponseMultiError is an error wrapping multiple validation errors
// returned by GetAccountResponse.ValidateAll() if the designated constraints
// aren't met.
type GetAccountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountResponseMultiError) AllErrors() []error { return m }

// GetAccountResponseValidationError is the validation error returned by
// GetAccountResponse.Validate if the designated constraints aren't met.
type GetAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountResponseValidationError) ErrorName() string {
	return "GetAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountResponseValidationError{}

// Validate checks the field values on SignOutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignOutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignOutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignOutRequestMultiError,
// or nil if none found.
func (m *SignOutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignOutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SignOutRequestMultiError(errors)
	}

	return nil
}

// SignOutRequestMultiError is an error wrapping multiple validation errors
// returned by SignOutRequest.ValidateAll() if the designated constraints
// aren't met.
type SignOutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignOutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignOutRequestMultiError) AllErrors() []error { return m }

// SignOutRequestValidationError is the validation error returned by
// SignOutRequest.Validate if the designated constraints aren't met.
type SignOutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignOutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignOutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignOutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignOutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignOutRequestValidationError) ErrorName() string { return "SignOutRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignOutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignOutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignOutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignOutRequestValidationError{}

// Validate checks the field values on SignOutResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SignOutResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignOutResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SignOutResponseMultiError, or nil if none found.
func (m *SignOutResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignOutResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return SignOutResponseMultiError(errors)
	}

	return nil
}

// SignOutResponseMultiError is an error wrapping multiple validation errors
// returned by SignOutResponse.ValidateAll() if the designated constraints
// aren't met.
type SignOutResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignOutResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignOutResponseMultiError) AllErrors() []error { return m }

// SignOutResponseValidationError is the validation error returned by
// SignOutResponse.Validate if the designated constraints aren't met.
type SignOutResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignOutResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignOutResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignOutResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignOutResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignOutResponseValidationError) ErrorName() string { return "SignOutResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignOutResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignOutResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignOutResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignOutResponseValidationError{}

// Validate checks the field values on AuthUserSignIn with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthUserSignIn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthUserSignIn with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthUserSignInMultiError,
// or nil if none found.
func (m *AuthUserSignIn) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthUserSignIn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for GrantType

	// no validation rules for Audience

	// no validation rules for ClientId

	// no validation rules for ClientSecret

	if len(errors) > 0 {
		return AuthUserSignInMultiError(errors)
	}

	return nil
}

// AuthUserSignInMultiError is an error wrapping multiple validation errors
// returned by AuthUserSignIn.ValidateAll() if the designated constraints
// aren't met.
type AuthUserSignInMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthUserSignInMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthUserSignInMultiError) AllErrors() []error { return m }

// AuthUserSignInValidationError is the validation error returned by
// AuthUserSignIn.Validate if the designated constraints aren't met.
type AuthUserSignInValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthUserSignInValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthUserSignInValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthUserSignInValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthUserSignInValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthUserSignInValidationError) ErrorName() string { return "AuthUserSignInValidationError" }

// Error satisfies the builtin error interface
func (e AuthUserSignInValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthUserSignIn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthUserSignInValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthUserSignInValidationError{}

// Validate checks the field values on RegisterAuthUser with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterAuthUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterAuthUser with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterAuthUserMultiError, or nil if none found.
func (m *RegisterAuthUser) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterAuthUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GivenName

	// no validation rules for FamilyName

	// no validation rules for Email

	// no validation rules for Nickname

	// no validation rules for Password

	// no validation rules for Connection

	// no validation rules for ClientId

	if len(errors) > 0 {
		return RegisterAuthUserMultiError(errors)
	}

	return nil
}

// RegisterAuthUserMultiError is an error wrapping multiple validation errors
// returned by RegisterAuthUser.ValidateAll() if the designated constraints
// aren't met.
type RegisterAuthUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterAuthUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterAuthUserMultiError) AllErrors() []error { return m }

// RegisterAuthUserValidationError is the validation error returned by
// RegisterAuthUser.Validate if the designated constraints aren't met.
type RegisterAuthUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterAuthUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterAuthUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterAuthUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterAuthUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterAuthUserValidationError) ErrorName() string { return "RegisterAuthUserValidationError" }

// Error satisfies the builtin error interface
func (e RegisterAuthUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterAuthUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterAuthUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterAuthUserValidationError{}

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserInfoMultiError, or nil
// if none found.
func (m *UserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for EmailVerified

	// no validation rules for Name

	// no validation rules for Nickname

	// no validation rules for Picture

	// no validation rules for Sub

	// no validation rules for UpdatedAt

	// no validation rules for UserId

	// no validation rules for Username

	if len(errors) > 0 {
		return UserInfoMultiError(errors)
	}

	return nil
}

// UserInfoMultiError is an error wrapping multiple validation errors returned
// by UserInfo.ValidateAll() if the designated constraints aren't met.
type UserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoMultiError) AllErrors() []error { return m }

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}
