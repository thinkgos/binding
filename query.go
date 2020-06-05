// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "net/http"

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var Query = queryBinding{}

var (
	_ Binding = (*queryBinding)(nil)
	_ Decoder = (*queryBinding)(nil)
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (b queryBinding) Bind(req *http.Request, obj interface{}) error {
	return b.BindForm(req.URL.Query(), obj)
}

func (queryBinding) BindForm(m map[string][]string, obj interface{}) error {
	if err := mapForm(obj, m); err != nil {
		return err
	}
	return validate(obj)
}

func (queryBinding) Decode(req *http.Request, obj interface{}) error {
	return mapForm(obj, req.URL.Query())
}

func (queryBinding) DecodeForm(m map[string][]string, obj interface{}) error {
	return mapForm(obj, m)
}
