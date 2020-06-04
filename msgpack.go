// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !nomsgpack

package binding

import (
	"bytes"
	"io"
	"net/http"

	"github.com/ugorji/go/codec"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var MsgPack = msgpackBinding{}

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj interface{}) error {
	return bindMsgPack(req.Body, obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return bindMsgPack(bytes.NewReader(body), obj)
}

func (msgpackBinding) BindReader(reader io.Reader, obj interface{}) error {
	return bindMsgPack(reader, obj)
}

func (msgpackBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeMsgPack(r.Body, obj)
}

func (msgpackBinding) DecodeBody(body []byte, obj interface{}) error {
	return decodeMsgPack(bytes.NewReader(body), obj)
}

func (msgpackBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeMsgPack(reader, obj)
}

func bindMsgPack(r io.Reader, obj interface{}) error {
	if err := decodeMsgPack(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeMsgPack(r io.Reader, obj interface{}) error {
	return codec.NewDecoder(r, new(codec.MsgpackHandle)).Decode(obj)
}
