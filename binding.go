// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
)

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with Bind,
// but it reads the body or io.Reader from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
	BindReader(io.Reader, interface{}) error
}

// BindingUri adds BindUri method to Binding. BindUri is similar with Bind,
// but it read the Params.
type BindingUri interface {
	Name() string
	BindUri(map[string][]string, interface{}) error
}

//  DecoderUri is similar with Decoder, but it read the Params.
type DecoderUri interface {
	Name() string
	DecodeUri(map[string][]string, interface{}) error
}

// Decoder decode data present in the request such as JSON request body, query parameters or
// the form POST.
type Decoder interface {
	Name() string
	Decode(*http.Request, interface{}) error
}

// DecoderBody is similar with Decoder,
// but it reads the body from supplied bytes or io.Reader instead of req.Body.
type DecoderBody interface {
	Decoder
	DecodeBody([]byte, interface{}) error
	DecodeReader(io.Reader, interface{}) error
}

// StructValidator is the minimal interface which needs to be implemented in
// order for it to be used as the validator engine for ensuring the correctness
// of the request. Gin provides a default implementation for this using
// https://github.com/go-playground/validator/tree/v8.18.2.
type StructValidator interface {
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is not a struct, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
	ValidateStruct(interface{}) error
	// ValidateVar validates a single variable using tag style validation.
	ValidateVar(field interface{}, tag string) error
	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
	Engine() *validator.Validate
}

// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v8.18.2
// under the hood.
var Validator StructValidator = &defaultValidator{}

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}

// DisableBindValidation closes the default validator.
func DisableBindValidation() {
	Validator = nil
}

// EnableJsonDecoderUseNumber sets true for EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.Which default value false.
func EnableJsonDecoderUseNumber() {
	EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.Which default value false.
func EnableJsonDecoderDisallowUnknownFields() {
	EnableDecoderDisallowUnknownFields = true
}
