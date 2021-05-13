// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !noprotopack

package binding

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var ProtoBuf = protobufBinding{}

var (
	_ Binding     = (*protobufBinding)(nil)
	_ BindingBody = (*protobufBinding)(nil)
	_ Decoder     = (*protobufBinding)(nil)
	_ DecoderBody = (*protobufBinding)(nil)
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(r *http.Request, obj interface{}) error {
	return b.BindReader(r.Body, obj)
}

func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}

func (b protobufBinding) BindReader(r io.Reader, obj interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}

func (b protobufBinding) Decode(r *http.Request, obj interface{}) error {
	return b.DecodeReader(r.Body, obj)
}

func (protobufBinding) DecodeBody(body []byte, obj interface{}) error {
	return proto.Unmarshal(body, obj.(proto.Message))
}

func (b protobufBinding) DecodeReader(r io.Reader, obj interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(buf, obj.(proto.Message))
}
