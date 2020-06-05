package binding

import (
	"net/http"
	"net/textproto"
	"reflect"
)

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var Header = headerBinding{}

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (b headerBinding) Bind(req *http.Request, obj interface{}) error {
	return b.BindHeader(req.Header, obj)
}

func (headerBinding) BindHeader(h map[string][]string, obj interface{}) error {
	if err := mapHeader(obj, h); err != nil {
		return err
	}
	return validate(obj)
}

func (headerBinding) Decode(req *http.Request, obj interface{}) error {
	return mapHeader(obj, req.Header)
}

func (headerBinding) DecodeHeader(h map[string][]string, obj interface{}) error {
	return mapHeader(obj, h)
}

func mapHeader(ptr interface{}, h map[string][]string) error {
	return mappingByPtr(ptr, headerSource(h), "header")
}

type headerSource map[string][]string

var _ setter = headerSource(nil)

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (isSetted bool, err error) {
	return setByForm(value, field, hs, textproto.CanonicalMIMEHeaderKey(tagValue), opt)
}
