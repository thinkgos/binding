// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var XML = xmlBinding{}

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) Bind(req *http.Request, obj interface{}) error {
	return bindXML(req.Body, obj)
}

func (xmlBinding) BindBody(body []byte, obj interface{}) error {
	return bindXML(bytes.NewReader(body), obj)
}

func (xmlBinding) BindReader(r io.Reader, obj interface{}) error {
	return bindXML(r, obj)
}

func (xmlBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeXML(r.Body, obj)
}

func (xmlBinding) DecodeBody(body []byte, obj interface{}) error {
	return decodeXML(bytes.NewReader(body), obj)
}

func (xmlBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeXML(reader, obj)
}

func decodeXML(r io.Reader, obj interface{}) error {
	return xml.NewDecoder(r).Decode(obj)
}

func bindXML(r io.Reader, obj interface{}) error {
	if err := decodeXML(r, obj); err != nil {
		return err
	}
	return validate(obj)
}
