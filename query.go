// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "net/http"

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var Query = queryBinding{}

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	if err := mapForm(obj, req.URL.Query()); err != nil {
		return err
	}
	return validate(obj)
}

func (queryBinding) Decode(req *http.Request, obj interface{}) error {
	return mapForm(obj, req.URL.Query())
}
