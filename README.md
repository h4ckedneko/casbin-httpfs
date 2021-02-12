# Casbin HTTP FileSystem Adapter

[![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/h4ckedneko/casbin-httpfs)
[![Latest Version](https://img.shields.io/github/v/release/h4ckedneko/casbin-httpfs?label=latest)](https://github.com/h4ckedneko/casbin-httpfs/releases)
[![License Name](https://img.shields.io/github/license/h4ckedneko/casbin-httpfs?color=blue)](https://github.com/h4ckedneko/casbin-httpfs/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/h4ckedneko/casbin-httpfs/Testing)](https://github.com/h4ckedneko/casbin-httpfs/actions?query=workflow:Testing)
[![Coverage Status](https://gocover.io/_badge/github.com/h4ckedneko/casbin-httpfs)](https://gocover.io/github.com/h4ckedneko/casbin-httpfs)
[![Report Card Status](https://goreportcard.com/badge/github.com/h4ckedneko/casbin-httpfs)](https://goreportcard.com/report/github.com/h4ckedneko/casbin-httpfs)

Package casbin-httpfs is a convenient http.FileSystem adapter for [Casbin](https://github.com/casbin/casbin) v2. It enables Casbin to load policy and model from anything that implements the http.FileSystem interface.

## Installation

Make sure you have a working [Go](https://golang.org/doc/install) workspace, then:

```
go get github.com/h4ckedneko/casbin-httpfs
```

For updating to latest stable release, do:

```
go get -u github.com/h4ckedneko/casbin-httpfs
```

## Usage

Here is a basic example for this package:

```go
package main

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	casbinfs "github.com/h4ckedneko/casbin-httpfs"
)

var fs http.FileSystem = http.Dir("/")

func main() {
	// Initialize a new adapter instance.
	adapter := casbinfs.NewAdapter(fs, "/path/to/policy.csv")

	// Initialize a new enforcer instance by passing the adapter.
	enforcer, _ := casbin.NewEnforcer("/path/to/model.conf", adapter)

	// Start checking for permissions.
	enforcer.Enforce("alice", "data1", "read")
}
```

## License

MIT © Lyntor Paul Figueroa. See [LICENSE](https://github.com/h4ckedneko/casbin-httpfs/blob/master/LICENSE) for full license text.
