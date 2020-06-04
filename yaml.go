// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noyamlpack

package binding

import (
	"bytes"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var YAML = yamlBinding{}

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(r *http.Request, obj interface{}) error {
	return bindYAML(r.Body, obj)
}

func (yamlBinding) BindBody(body []byte, obj interface{}) error {
	return bindYAML(bytes.NewReader(body), obj)
}

func (yamlBinding) BindReader(reader io.Reader, obj interface{}) error {
	return bindYAML(reader, obj)
}

func (yamlBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeYAML(r.Body, obj)
}

func (yamlBinding) DecodeBody(body []byte, obj interface{}) error {
	return decodeYAML(bytes.NewReader(body), obj)
}

func (yamlBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeYAML(reader, obj)
}

func bindYAML(r io.Reader, obj interface{}) error {
	if err := decodeYAML(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeYAML(r io.Reader, obj interface{}) error {
	return yaml.NewDecoder(r).Decode(obj)
}
