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

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(req *http.Request, obj interface{}) error {
	return bindProtobuf(req.Body, obj)
}

func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `validate:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}

func (b protobufBinding) BindReader(r io.Reader, obj interface{}) error {
	return bindProtobuf(r, obj)
}

func (b protobufBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeProtobuf(r.Body, obj)
}

func (protobufBinding) DecodeBody(body []byte, obj interface{}) error {
	return proto.Unmarshal(body, obj.(proto.Message))
}

func (b protobufBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeProtobuf(reader, obj)
}

func decodeProtobuf(reader io.Reader, obj interface{}) error {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return proto.Unmarshal(buf, obj.(proto.Message))
}

func bindProtobuf(reader io.Reader, obj interface{}) error {
	if err := decodeProtobuf(reader, obj); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `validate:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}
