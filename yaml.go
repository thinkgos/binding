// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noyamlpack

package binding

import (
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var YAML = yamlBinding{}

var (
	_ Binding     = (*yamlBinding)(nil)
	_ BindingBody = (*yamlBinding)(nil)
	_ Decoder     = (*yamlBinding)(nil)
	_ DecoderBody = (*yamlBinding)(nil)
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (b yamlBinding) Bind(r *http.Request, obj interface{}) error {
	return b.BindReader(r.Body, obj)
}

func (b yamlBinding) BindBody(body []byte, obj interface{}) error {
	if err := b.DecodeBody(body, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b yamlBinding) BindReader(r io.Reader, obj interface{}) error {
	if err := b.DecodeReader(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b yamlBinding) Decode(r *http.Request, obj interface{}) error {
	return b.DecodeReader(r.Body, obj)
}

func (yamlBinding) DecodeBody(body []byte, obj interface{}) error {
	return yaml.Unmarshal(body, obj)
}

func (yamlBinding) DecodeReader(r io.Reader, obj interface{}) error {
	return yaml.NewDecoder(r).Decode(obj)
}
