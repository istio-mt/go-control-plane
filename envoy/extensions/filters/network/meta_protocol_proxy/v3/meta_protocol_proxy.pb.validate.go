// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/meta_protocol_proxy/v3/meta_protocol_proxy.proto

package meta_protocol_proxyv3

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

// Validate checks the field values on MetaProtocolProxy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MetaProtocolProxy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetaProtocolProxy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetaProtocolProxyMultiError, or nil if none found.
func (m *MetaProtocolProxy) ValidateAll() error {
	return m.validate(true)
}

func (m *MetaProtocolProxy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := MetaProtocolProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetApplicationProtocol() == nil {
		err := MetaProtocolProxyValidationError{
			field:  "ApplicationProtocol",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetApplicationProtocol()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaProtocolProxyValidationError{
					field:  "ApplicationProtocol",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaProtocolProxyValidationError{
					field:  "ApplicationProtocol",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApplicationProtocol()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaProtocolProxyValidationError{
				field:  "ApplicationProtocol",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetMetaProtocolFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  fmt.Sprintf("MetaProtocolFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  fmt.Sprintf("MetaProtocolFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetaProtocolProxyValidationError{
					field:  fmt.Sprintf("MetaProtocolFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.RouteSpecifier.(type) {

	case *MetaProtocolProxy_Rds:

		if all {
			switch v := interface{}(m.GetRds()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  "Rds",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  "Rds",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRds()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetaProtocolProxyValidationError{
					field:  "Rds",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MetaProtocolProxy_RouteConfig:

		if all {
			switch v := interface{}(m.GetRouteConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  "RouteConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetaProtocolProxyValidationError{
						field:  "RouteConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetaProtocolProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := MetaProtocolProxyValidationError{
			field:  "RouteSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return MetaProtocolProxyMultiError(errors)
	}
	return nil
}

// MetaProtocolProxyMultiError is an error wrapping multiple validation errors
// returned by MetaProtocolProxy.ValidateAll() if the designated constraints
// aren't met.
type MetaProtocolProxyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetaProtocolProxyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetaProtocolProxyMultiError) AllErrors() []error { return m }

// MetaProtocolProxyValidationError is the validation error returned by
// MetaProtocolProxy.Validate if the designated constraints aren't met.
type MetaProtocolProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaProtocolProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaProtocolProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaProtocolProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaProtocolProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaProtocolProxyValidationError) ErrorName() string {
	return "MetaProtocolProxyValidationError"
}

// Error satisfies the builtin error interface
func (e MetaProtocolProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetaProtocolProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaProtocolProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaProtocolProxyValidationError{}

// Validate checks the field values on ApplicationProtocol with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApplicationProtocol) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplicationProtocol with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplicationProtocolMultiError, or nil if none found.
func (m *ApplicationProtocol) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplicationProtocol) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ApplicationProtocolValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCodec() == nil {
		err := ApplicationProtocolValidationError{
			field:  "Codec",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCodec()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationProtocolValidationError{
					field:  "Codec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationProtocolValidationError{
					field:  "Codec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCodec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationProtocolValidationError{
				field:  "Codec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ApplicationProtocolMultiError(errors)
	}
	return nil
}

// ApplicationProtocolMultiError is an error wrapping multiple validation
// errors returned by ApplicationProtocol.ValidateAll() if the designated
// constraints aren't met.
type ApplicationProtocolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationProtocolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationProtocolMultiError) AllErrors() []error { return m }

// ApplicationProtocolValidationError is the validation error returned by
// ApplicationProtocol.Validate if the designated constraints aren't met.
type ApplicationProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationProtocolValidationError) ErrorName() string {
	return "ApplicationProtocolValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationProtocolValidationError{}

// Validate checks the field values on MetaRds with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetaRds) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetaRds with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MetaRdsMultiError, or nil if none found.
func (m *MetaRds) ValidateAll() error {
	return m.validate(true)
}

func (m *MetaRds) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetConfigSource() == nil {
		err := MetaRdsValidationError{
			field:  "ConfigSource",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConfigSource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetaRdsValidationError{
					field:  "ConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetaRdsValidationError{
					field:  "ConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfigSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetaRdsValidationError{
				field:  "ConfigSource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetRouteConfigName()) < 1 {
		err := MetaRdsValidationError{
			field:  "RouteConfigName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetaRdsMultiError(errors)
	}
	return nil
}

// MetaRdsMultiError is an error wrapping multiple validation errors returned
// by MetaRds.ValidateAll() if the designated constraints aren't met.
type MetaRdsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetaRdsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetaRdsMultiError) AllErrors() []error { return m }

// MetaRdsValidationError is the validation error returned by MetaRds.Validate
// if the designated constraints aren't met.
type MetaRdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaRdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaRdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaRdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaRdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaRdsValidationError) ErrorName() string { return "MetaRdsValidationError" }

// Error satisfies the builtin error interface
func (e MetaRdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetaRds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaRdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaRdsValidationError{}
