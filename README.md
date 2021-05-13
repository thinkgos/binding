# binding this is gin binding package, but with option tags

[![GoDoc](https://godoc.org/github.com/things-go/binding?status.svg)](https://godoc.org/github.com/things-go/binding)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/binding?tab=doc)
[![Build Status](https://www.travis-ci.com/things-go/binding.svg?branch=master)](https://www.travis-ci.com/thinkgos/binding)
[![codecov](https://codecov.io/gh/things-go/binding/branch/master/graph/badge.svg)](https://codecov.io/gh/things-go/binding)
![Action Status](https://github.com/things-go/binding/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/binding)](https://goreportcard.com/report/github.com/things-go/binding)
[![Licence](https://img.shields.io/github/license/things-go/binding)](https://raw.githubusercontent.com/things-go/binding/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/binding)](https://github.com/things-go/binding/tags)

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
*** binding tag is `binding` ***

## References
- [gin](https://github.com/gin-gonic/gin)