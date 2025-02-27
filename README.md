# Argus

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/argus.svg)](https://pkg.go.dev/github.com/aethiopicuschan/argus)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/argus)](https://goreportcard.com/report/github.com/aethiopicuschan/argus)
[![CI](https://github.com/aethiopicuschan/argus/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/argus/actions/workflows/ci.yaml)

Argus is a library that provides lightweight, minimal structured logging.

## Installation

```bash
go get -u github.com/aethiopicuschan/argus
```

## Usage

```go
package main

import (
	"os"

	"github.com/aethiopicuschan/argus"
)

func main() {
	logger := argus.NewLogger(os.Stdout)
	logger.Info().Add("key", "value").Print()
	// Output: {"level":"INFO","time":"2006-01-02T15:04:05.000000+09:00","key":"value"}
}
```
