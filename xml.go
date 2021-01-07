// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"encoding/xml"
	"io"
	"net/http"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var XML = xmlBinding{}

var (
	_ Binding     = (*xmlBinding)(nil)
	_ BindingBody = (*xmlBinding)(nil)
	_ Decoder     = (*xmlBinding)(nil)
	_ DecoderBody = (*xmlBinding)(nil)
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (b xmlBinding) Bind(r *http.Request, obj interface{}) error {
	return b.BindReader(r.Body, obj)
}

func (b xmlBinding) BindBody(body []byte, obj interface{}) error {
	if err := b.DecodeBody(body, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b xmlBinding) BindReader(r io.Reader, obj interface{}) error {
	if err := b.DecodeReader(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b xmlBinding) Decode(r *http.Request, obj interface{}) error {
	return b.DecodeReader(r.Body, obj)
}

func (xmlBinding) DecodeBody(body []byte, obj interface{}) error {
	return xml.Unmarshal(body, obj)
}

func (xmlBinding) DecodeReader(r io.Reader, obj interface{}) error {
	return xml.NewDecoder(r).Decode(obj)
}
