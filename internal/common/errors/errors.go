package errors

import (
	"encoding/json"
	"fmt"
	"strings"

	goerrors "errors"

	"github.com/badbud-backend-v2/internal/common/validate"
	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
)

var (
	// There is such error?
	ErrorTypeUnknown = ErrorType{}
	// For all database related errors
	ErrorTypeDB = ErrorType{"db"}
	// For error calling 3rd party API
	ErrorTypeUpstream = ErrorType{"upstream"}
	// When public user access private resource
	ErrorTypeUnauthorized = ErrorType{"unauthorized"}
	// When user does not have permission to do action on private resource
	ErrorTypeForbidden = ErrorType{"forbidden"}
	// Generial user input validation error
	ErrorTypeNotFound = ErrorType{"not-found"}
	// Generial user input validation error
	ErrorTypeValidation = ErrorType{"validation"}
)

// New returns service error
func New(internal error, t ErrorType, msg string, args ...any) *Error {
	return &Error{
		Type:     t,
		Message:  fmt.Sprintf(msg, args...),
		Internal: internal,
	}
}

// FromValidator returns service error from validator errors
func FromValidator(errs validator.ValidationErrors, msgAndArgs ...any) *Error {
	var msg string
	if len(msgAndArgs) > 0 {
		msg = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	e := Validation(nil, msg)

	for _, v := range errs {
		// take namespace to avoid confusion in nested field
		// and cut of the root struct
		_, fn, _ := strings.Cut(v.Namespace(), ".")
		e.Errors = append(e.Errors, &ErrorDetail{
			// workaround to match the field name in json tag
			Path:    strcase.ToSnakeWithIgnore(fn, "."),
			Message: v.Translate(validate.Translator()),
		})
	}

	return e
}

// Unknown quickly creates unknown error.
func Unknown(internal error, msg string) *Error {
	return New(internal, ErrorTypeUnknown, msg)
}

// DB quickly creates db related error.
func DB(internal error, msg string) *Error {
	return New(internal, ErrorTypeDB, msg)
}

// Upstream quickly creates 3rd party related error.
func Upstream(internal error, msg string) *Error {
	return New(internal, ErrorTypeUpstream, msg)
}

// Unauthorized quickly creates unauthorized error.
//
// Do pass internal error or addition information for tracing purpose
func Unauthorized(errOrInfo any, msg string) *Error {
	return New(errFrom(errOrInfo), ErrorTypeUnauthorized, msg)
}

// Forbidden quickly creates permission error.
//
// Do pass internal error or addition information for tracing purpose
func Forbidden(errOrInfo any, msg string) *Error {
	return New(errFrom(errOrInfo), ErrorTypeForbidden, msg)
}

// NotFound quickly creates resource not found error.
//
// Do pass internal error or addition information for tracing purpose
func NotFound(errOrInfo any, msg string) *Error {
	return New(errFrom(errOrInfo), ErrorTypeNotFound, msg)
}

// Validation quickly creates validation error.
//
// Do not put error details to the first param, use chained method AddError or AddErrors instead.
func Validation(errOrInfo any, msg string) *Error {
	if msg == "" {
		msg = "Some required information is missing or incorrect"
	}
	return New(errFrom(errOrInfo), ErrorTypeValidation, msg)
}

// Alias functions to avoid duplicate imports
var (
	// Is reports whether any error in err's tree matches target.
	//
	// Alias of standard errors package.
	Is = goerrors.Is
	// As finds the first error in err's tree that matches target, and if one is found, sets target to that error value and returns true. Otherwise, it returns false.
	//
	// Alias of standard errors package.
	As = goerrors.As
	// Join returns an error that wraps the given errors. Any nil error values are discarded.
	//
	// Alias of standard errors package.
	// Join = goerrors.Join
	// Unwrap returns the result of calling the Unwrap method on err, if err's type contains an Unwrap method returning error. Otherwise, Unwrap returns nil.
	//
	// Alias of standard errors package.
	Unwrap = goerrors.Unwrap
)

// ErrorType represents type of a service error
// swagger:strfmt string
type ErrorType struct {
	// Struct-based Enums, lets us work with code secure by design.
	// the field is unexported, it’s not possible to fill it from outside the package
	t string
}

func (et ErrorType) String() string {
	if et.t == "" {
		return "unknown"
	}
	return et.t
}

// MarshalJSON interface
func (et ErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(et.String())
}

// ErrorDetail represents the error details
type ErrorDetail struct {
	// Path to the error field
	// example: name
	Path string `json:"path,omitempty"`
	// Error message for the given path
	// example: Name is required
	Message string `json:"message,omitempty"`

	// ? add validation tag ?
}

func (ed *ErrorDetail) String() string {
	return ed.Path + ": " + ed.Message
}

// ErrorDetail represents a list of error detail
type ErrorDetails []*ErrorDetail

func (eds *ErrorDetails) String() string {
	if eds == nil {
		return ""
	}

	var s strings.Builder
	for i, v := range *eds {
		if i > 0 {
			s.WriteString("\n")
		}
		s.WriteString(v.String())
	}
	return s.String()
}

// Error represents a service error
type Error struct {
	// Generic type of the error
	// example: validation
	Type ErrorType `json:"type,omitempty"`
	// Error message
	// example: Some required information is missing or incorrect
	Message string `json:"message,omitempty"`
	// Error details, in validation error for example
	Errors ErrorDetails `json:"errors,omitempty"`
	// Internal error cause this
	Internal error `json:"-"`

	// wrapped error to satisfy Unwrap interface
	root error
}

// String satisfies Stringer interface
func (e *Error) String() string {
	var s strings.Builder
	s.WriteString("type=")
	s.WriteString(e.Type.t)
	s.WriteString(", message=")
	s.WriteString(e.Message)
	if len(e.Errors) > 0 {
		s.WriteString(", errors=")
		s.WriteString(e.Errors.String())
	}
	if e.Internal != nil {
		s.WriteString(", internal=")
		s.WriteString(e.Internal.Error())
	}
	return s.String()
}

// Error satisfies error interface
func (e *Error) Error() string {
	return e.String()
}

// to custom json ouput
type (
	alias      Error
	customJson struct {
		*alias
		Internal string `json:"internal,omitempty"`
	}
)

// MarshalJSON modification
func (e *Error) MarshalJSON() ([]byte, error) {
	output := customJson{alias: (*alias)(e)}
	if e.Internal != nil {
		output.Internal = e.Internal.Error()
	}
	return json.Marshal(output)
}

// WithInternal returns clone of Error with given error.
// Note: The Errors field is a shallow copy.
//
// The original service error as well as the initial internal error can still be examined by errors.Is
func (e *Error) WithInternal(err error) *Error {
	if err == nil {
		return e
	}
	if e.Internal != nil {
		// err = goerrors.Join(e.Internal, err)
	}
	return &Error{
		Type:     e.Type,
		Message:  e.Message,
		Errors:   e.Errors,
		Internal: err,
		root:     e,
	}
}

// WithInternalf allows to set custom message. Should be used only when there is no internal error!
func (e *Error) WithInternalf(format string, a ...any) *Error {
	return e.WithInternal(fmt.Errorf(format, a...))
}

// SetErrors sets error details.
//
// Do NOT set internal error with this. Use WithInternal instead.
func (e *Error) SetErrors(errs ErrorDetails) *Error {
	e.Errors = errs
	return e
}

// AddErrors appends given error details into errors.
//
// Do NOT set internal error with this. Use WithInternal instead.
func (e *Error) AddErrors(errs ...*ErrorDetail) *Error {
	e.Errors = append(e.Errors, errs...)
	return e
}

// AddError add error detail into errors.
//
// Do NOT set internal error with this. Use WithInternal instead.
func (e *Error) AddError(path, msg string) *Error {
	e.Errors = append(e.Errors, &ErrorDetail{Path: path, Message: msg})
	return e
}

// Unwrap satisfies the Go 1.13 error wrapper interface.
func (e *Error) Unwrap() error {
	return e.root
}

func errFrom(in any) error {
	err, ok := in.(error)
	if !ok && in != nil {
		err = fmt.Errorf("%+v", in)
	}
	return err
}
