# Programming Project Tracker - Reference Implementation Go Library

[![Go Tests](https://github.com/otg996/libptrack-go/actions/workflows/validation.yml/badge.svg)](https://github.com/otg996/libptrack-go/actions/workflows/validation.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/otg996/libptrack-go)](https://goreportcard.com/report/github.com/otg996/libptrack-go)
[![Go.Dev Reference](https://pkg.go.dev/badge/github.com/otg996/libptrack-go.svg)](https://pkg.go.dev/github.com/otg996/libptrack-go)

> A Go library for discovering software development projects in a filesystem according to the Programming Project Tracker Specification.

This library is the official Go reference implementation for the [Programming Project Tracker Specification](https://github.com/otg996/programming-project-tracker-spec). It provides the core scanning logic for finding projects.

## Installation

```shell
go get github.com/otg996/libptrack-go@latest
```

## Usage

Basic example of how to use the library to scan a directory:

```go
package main

import (
 "fmt"
 "log"

 ptrack "github.com/otg996/libptrack-go"
)

func main() {
 // Scan the current working directory.
 projects, err := ptrack.ScanDirectory(".")
 if err != nil {
  log.Fatalf("Error scanning directory: %v", err)
 }

 fmt.Println("Found projects:")
 for _, projectPath := range projects {
  fmt.Println("- ", projectPath)
 }
}
```

## Documentation

Full API documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/otg996/libptrack-go).

## Contributing

This project follows the same philosophy and process as the main specification. Please see the [main contributing guide](https://github.com/otg996/programming-project-tracker-spec/blob/main/CONTRIBUTING.md) for details on our workflow. The technical setup for this repository is outlined below.

## License

This library is licensed under the [LGPL-2.1 License](LICENSE).
