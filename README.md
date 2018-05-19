# flagvar 

[![](https://godoc.org/github.com/sgreben/flagvar?status.svg)](http://godoc.org/github.com/sgreben/flagvar) [![](https://goreportcard.com/badge/github.com/sgreben/flagvar/goreportcard)](https://goreportcard.com/report/github.com/sgreben/flagvar) [![](https://gocover.io/_badge/github.com/sgreben/flagvar)](https://gocover.io/github.com/sgreben/flagvar) [![Build Status](https://travis-ci.org/sgreben/flagvar.svg?branch=master)](https://travis-ci.org/sgreben/flagvar)

A collection of CLI argument types for the `flag` package. 

```go
import "github.com/sgreben/flagvar"
```

Or just copy & paste what you need. It's public domain.

<!-- TOC -->

- [Example](#example)
- [Conventions](#conventions)
- [Types](#types)

<!-- /TOC -->

## Example

```go
package main

import (
    "flag"
    "fmt"
    "github.com/sgreben/flagvar"
)

var (
    fruit    = flagvar.Enum{Choices: []string{"apple","banana"}}
    urls     flagvar.URLs
    settings flagvar.Assignments
)

func main() {
    flag.Var(&fruit, "fruit", fmt.Sprintf("set a fruit %v", fruit.Choices))
    flag.Var(&urls, "url", "add a URL")
    flag.Var(&settings, "set", "set key=value")
    flag.Parse()
}
```

```sh
$ go run main.go -set abc=xyz -url https://github.com 
# no error

$ go run main.go -set abc=xyz -url ://github.com
invalid value "://github.com" for flag -url: parse ://github.com: missing protocol scheme

$ go run /tmp/main.go -fruit kiwi
invalid value "kiwi" for flag -fruit: "kiwi" must be one of [apple banana]
```

## Conventions

- Pluralized argument types (e.g. `Strings`, `Assignments`) can be specified repeatedly, the values are collected in a slice.
- -Set types (`EnumSet`, `StringSet`) de-duplicate provided values.
- The resulting value is stored in `.Value` for singular types and in `.Values` for plural types
- The original argument string is stored in `.Text` for singular types and in `.Texts` for plural types

## Types

Here's a compact overview:

| `flagvar` type | example CLI arg    | type of resulting Go value           |
|----------------|--------------------|--------------------------------------|
| [Alternative](https://godoc.org/github.com/sgreben/flagvar#Alternative)  |           |  |
| [Assignment](https://godoc.org/github.com/sgreben/flagvar#Assignment)  | KEY=VALUE          | KV{Key:"Key", Value:"Value"} |
| [Assignments](https://godoc.org/github.com/sgreben/flagvar#Assignments) | KEY=VALUE          | []KV                         |
| [AssignmentsMap](https://godoc.org/github.com/sgreben/flagvar#AssignmentsMap) | KEY=VALUE          | map[string]string                         |
| [Enum](https://godoc.org/github.com/sgreben/flagvar#Enum)        | apple              | string                               |
| [Enums](https://godoc.org/github.com/sgreben/flagvar#Enums)       | apple              | []string                             |
| [EnumSet](https://godoc.org/github.com/sgreben/flagvar#EnumSet)     | apple              | []string                             |
| [File](https://godoc.org/github.com/sgreben/flagvar#File)        | ./README.md        | string                               |
| [Files](https://godoc.org/github.com/sgreben/flagvar#Files)        | ./README.md        | string                               |
| [Floats](https://godoc.org/github.com/sgreben/flagvar#Floats)      | 1.234              | []float64                            |
| [Glob](https://godoc.org/github.com/sgreben/flagvar#Glob)        | src/**.js          | glob.Glob                            |
| [Globs](https://godoc.org/github.com/sgreben/flagvar#Globs)       | src/**.js          | glob.Glob                            |
| [Ints](https://godoc.org/github.com/sgreben/flagvar#Ints)        | 1002               | []int64                              |
| [JSON](https://godoc.org/github.com/sgreben/flagvar#JSON)        | '{"a":1}'          | interface{}                          |
| [JSONs](https://godoc.org/github.com/sgreben/flagvar#JSONs)       | '{"a":1}'          | []interface{}                        |
| [Strings](https://godoc.org/github.com/sgreben/flagvar#Strings)     | "xyxy"             | []string                             |
| [StringSet](https://godoc.org/github.com/sgreben/flagvar#StringSet)  | "xyxy"             | []string                             |
| [Template](https://godoc.org/github.com/sgreben/flagvar#Template)    | "{{.Size}}"        | *template.Template                   |
| [Templates](https://godoc.org/github.com/sgreben/flagvar#Templates)   | "{{.Size}}"        | []*template.Template                 |
| [Time](https://godoc.org/github.com/sgreben/flagvar#Time)        | "10:30 AM"         | time.Time                            |
| [Times](https://godoc.org/github.com/sgreben/flagvar#Times)       | "10:30 AM"         | []time.Time                          |
| [URL](https://godoc.org/github.com/sgreben/flagvar#URL)         | https://github.com | *url.URL                             |
| [URLs](https://godoc.org/github.com/sgreben/flagvar#URLs)        | https://github.com | []*url.URL                           |
| [Wrap](https://godoc.org/github.com/sgreben/flagvar#Wrap)        |                    |                                      |
| [WrapFunc](https://godoc.org/github.com/sgreben/flagvar#WrapFunc)    |                    |                                      |
| [WrapPointer](https://godoc.org/github.com/sgreben/flagvar#WrapPointer)    |                    |                                      |