# binding this is gin binding package, but with option tags

[![GoDoc](https://godoc.org/github.com/thinkgos/binding?status.svg)](https://godoc.org/github.com/thinkgos/binding)
[![Build Status](https://www.travis-ci.org/thinkgos/binding.svg?branch=master)](https://www.travis-ci.org/thinkgos/binding)
[![codecov](https://codecov.io/gh/thinkgos/binding/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/binding)
![Action Status](https://github.com/thinkgos/binding/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/binding)](https://goreportcard.com/report/github.com/thinkgos/binding)
[![Licence](https://img.shields.io/github/license/thinkgos/binding)](https://raw.githubusercontent.com/thinkgos/binding/master/LICENSE)

## Build with [jsoniter](https://github.com/json-iterator/go)

binding uses `encoding/json` as default json package, but you can change to [jsoniter](https://github.com/json-iterator/go) by build from other tags.

```sh
$ go build -tags=jsoniter .
```

## Build without 
- [msg](github.com/ugorji/go)
- [protobuf](github.com/golang/protobuf/proto)
- [yaml](https://github.com/go-yaml/yaml)
   
binding uses `msg`,`protobuf`,`yaml`, you can build without them with tags.
   
```sh
   $ go build -tags=noprotopack,noyamlpack,nomsgpack .
```

## note 
*** binding tag same as [validator](github.com/go-playground/validator/v10) which is `validate` not `binding` ***

## References
- [gin](https://github.com/gin-gonic/gin)