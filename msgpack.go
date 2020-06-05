// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !nomsgpack

package binding

import (
	"io"
	"net/http"

	"github.com/ugorji/go/codec"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var MsgPack = msgpackBinding{}

var (
	_ Binding     = (*msgpackBinding)(nil)
	_ BindingBody = (*msgpackBinding)(nil)
	_ Decoder     = (*msgpackBinding)(nil)
	_ DecoderBody = (*msgpackBinding)(nil)
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (msgpackBinding) Bind(req *http.Request, obj interface{}) error {
	return bindMsgPack(req.Body, obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return bindMsgPackBody(body, obj)
}

func (msgpackBinding) BindReader(reader io.Reader, obj interface{}) error {
	return bindMsgPack(reader, obj)
}

func (msgpackBinding) Decode(r *http.Request, obj interface{}) error {
	return decodeMsgPack(r.Body, obj)
}

func (msgpackBinding) DecodeBody(body []byte, obj interface{}) error {
	return decodeMsgPackBody(body, obj)
}

func (msgpackBinding) DecodeReader(reader io.Reader, obj interface{}) error {
	return decodeMsgPack(reader, obj)
}

func bindMsgPackBody(in []byte, obj interface{}) error {
	if err := decodeMsgPackBody(in, obj); err != nil {
		return err
	}
	return validate(obj)
}

func bindMsgPack(r io.Reader, obj interface{}) error {
	if err := decodeMsgPack(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeMsgPackBody(in []byte, obj interface{}) error {
	return codec.NewDecoderBytes(in, new(codec.MsgpackHandle)).Decode(obj)
}

func decodeMsgPack(r io.Reader, obj interface{}) error {
	return codec.NewDecoder(r, new(codec.MsgpackHandle)).Decode(obj)
}
