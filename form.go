// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"
)

const defaultMemory = 32 << 20

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Form          = formBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
)

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (b formBinding) Bind(req *http.Request, obj interface{}) error {
	if err := b.Decode(req, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formBinding) Decode(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	return mapForm(obj, req.Form)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (b formPostBinding) Bind(req *http.Request, obj interface{}) error {
	if err := b.Decode(req, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) Decode(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	return mapForm(obj, req.PostForm)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (b formMultipartBinding) Bind(req *http.Request, obj interface{}) error {
	if err := b.Decode(req, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formMultipartBinding) Decode(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	return mappingByPtr(obj, (*multipartRequest)(req), "form")
}
