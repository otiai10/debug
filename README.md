# debug printer

[![Build Status](https://travis-ci.org/otiai10/debug.svg?branch=master)](https://travis-ci.org/otiai10/debug)
[![codecov](https://codecov.io/gh/otiai10/debug/branch/master/graph/badge.svg)](https://codecov.io/gh/otiai10/debug)
[![GoDoc](https://godoc.org/github.com/otiai10/debug?status.png)](https://godoc.org/github.com/otiai10/debug)

```go
package main

import "github.com/otiai10/debug"

func main() {
  debug.Fatalln("Print line here!!")
}

// 2018/08/13 23:19:00 /your/path/main.go:6 Print line here!!
```