// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/thinkgos/binding/internal/json"
)

// EnableDecoderUseNumber is used to call the UseNumber method on the JSON
// Decoder instance. UseNumber causes the Decoder to unmarshal a number into an
// interface{} as a Number instead of as a float64.
var EnableDecoderUseNumber = false

// EnableDecoderDisallowUnknownFields is used to call the DisallowUnknownFields method
// on the JSON Decoder instance. DisallowUnknownFields causes the Decoder to
// return an error when the destination is a struct and the input contains object
// keys which do not match any non-ignored, exported fields in the destination.
var EnableDecoderDisallowUnknownFields = false

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var JSON = jsonBinding{}

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return bindJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj interface{}) error {
	return bindJSON(bytes.NewReader(body), obj)
}

func (jsonBinding) BindReader(reader io.Reader, obj interface{}) error {
	return bindJSON(reader, obj)
}

func (jsonBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeJSON(r.Body, obj)
}

func (jsonBinding) DecodeBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func (jsonBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeJSON(reader, obj)
}

func bindJSON(r io.Reader, obj interface{}) error {
	if err := decodeJSON(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if EnableDecoderDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	return decoder.Decode(obj)
}
