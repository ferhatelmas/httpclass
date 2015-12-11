## httpclass

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/httpclass)
[![Build Status](https://travis-ci.org/ferhatelmas/httpclass.png?branch=master)](https://travis-ci.org/ferhatelmas/httpclass)

> Returns the class of a HTTP status code.

### Install

```
go get github.com/ferhatelmas/httpclass
```

### Usage

```go
import "github.com/ferhatelmas/httpclass"

httpclass.Get(http.StatusOK)
//=> httpclass.Success, nil

httpclass.Get(http.StatusBadRequest)
//=> httpclass.ClientError, nil

httpclass.Get(0)
//=> Informational, "Wrong status code 0"
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
