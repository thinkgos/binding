// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var Uri = uriBinding{}

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) BindUri(m map[string][]string, obj interface{}) error {
	if err := mapUri(obj, m); err != nil {
		return err
	}
	return validate(obj)
}

func (uriBinding) DecodeUri(m map[string][]string, obj interface{}) error {
	return mapUri(obj, m)
}
