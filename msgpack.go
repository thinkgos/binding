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

func (b msgpackBinding) Bind(req *http.Request, obj interface{}) error {
	return b.BindReader(req.Body, obj)
}

func (b msgpackBinding) BindBody(body []byte, obj interface{}) error {
	if err := b.DecodeBody(body, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b msgpackBinding) BindReader(r io.Reader, obj interface{}) error {
	if err := b.DecodeReader(r, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (b msgpackBinding) Decode(r *http.Request, obj interface{}) error {
	return b.DecodeReader(r.Body, obj)
}

func (msgpackBinding) DecodeBody(body []byte, obj interface{}) error {
	return codec.NewDecoderBytes(body, new(codec.MsgpackHandle)).Decode(obj)
}

func (msgpackBinding) DecodeReader(r io.Reader, obj interface{}) error {
	return codec.NewDecoder(r, new(codec.MsgpackHandle)).Decode(obj)
}
