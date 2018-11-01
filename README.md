# Menoh Go

[![Build Status](https://travis-ci.org/pfnet-research/go-menoh.svg?branch=master)](https://travis-ci.org/pfnet-research/go-menoh)
[![Build status](https://ci.appveyor.com/api/projects/status/29w9dkt4noorr7rl/branch/master?svg=true)](https://ci.appveyor.com/project/disktnk/go-menoh-27309/branch/master)
[![codecov](https://codecov.io/gh/pfnet-research/go-menoh/branch/master/graph/badge.svg)](https://codecov.io/gh/pfnet-research/go-menoh)
[![GoDoc](https://godoc.org/github.com/pfnet-research/go-menoh?status.svg)](http://godoc.org/github.com/pfnet-research/go-menoh)
[![Go Report Card](https://goreportcard.com/badge/github.com/pfnet-research/go-menoh)](https://goreportcard.com/report/github.com/pfnet-research/go-menoh)

Golang binding for [Menoh](https://github.com/pfnet-research/menoh)

## Requirements

- Go 1.10+
- [Menoh](https://github.com/pfnet-research/menoh) 1.1.1+

## Install

After install Menoh, then

```bash
$ go get -u github.com/pfnet-research/go-menoh
```

### Linux/Mac

Add a path to library to `LD_LIBRARY_PATH` environment variable. Menoh libraries are installed to `/usr/local/lib` on default.

```bash
$ export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
```

### Windows

Add a path to DLLs distributed by Menoh to local Path environment.

```
\path\to\menoh\bin
  |- libiomp5md.dll
  |- menoh.dll
  |- mkldnn.dll
  |- mklml.dll
```

```cmd
set PATH=\path\to\menoh\bin;%PATH%
```

## Usage

- [example/vgg16](example/vgg16) is a tutorial for this package.
- [example/mnist](example/mnist) is an example using MNIST dataset and model.

## Development

### Test

Download ONNX file, using in menoh-rust test, before testing.

```bash
$ wget https://github.com/pfnet-research/menoh-rs/releases/download/assets/MLP.onnx -P test_data
$ go test ./...
```

Additionally go-menoh follows `gofmt` with simplify option (`-s`), `go vet` and `golint`.

## Note

- At first name of this repository is "menoh-go", and renamed to "**go-menoh**" to follow [recommended naming rule](https://github.com/golang/go/wiki/PackagePublishing)

## License

MIT License (see [LICENSE](/LICENSE) file).
