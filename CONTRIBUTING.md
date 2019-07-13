# How to contribute

1. Create an issue where you tell me that you want to contribute (and what, Runtime, Parser or API, or combinations of them)
2. I will add you to the repository and the respective team(s)

Alternatively, you can

1. Fork the repository
2. Make changes
3. Open a Pull Request against `develop`

# Code requirements

* No test cases can be removed and/or changed (if a test fails, but ran green in the past, this is a regression)
  * If you think a test is incorrect, open an issue, stating why the test should be changed/removed
* To be consistent, new objects are created with constructor methods (`foo := NewFoo()` instead of `foo := &Foo{...}`)
  * Within constructor functions, the keyword `new` shall be used (`f := new(Foo); f.x = y` instead of `f := &Foo{x: y}`)
* [![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)
* Readability > Performance (there **are** exceptions, but don't optimize prematurely)

# Getting started

* Checkout the repository
* Run `go test -short ./...` to ensure everything works
* To test everything, run `go test -v ./...` (includes parser and conformance tests, takes a while)
* To build the application, run `go build -o gojisvm .` and it will build the executable file `./gojisvm`