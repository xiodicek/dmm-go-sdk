# DMM SDK for Go (v3)
[![GoDoc](https://img.shields.io/badge/go-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/DMMcomLabo/dmm-go-sdk/blob/master/LICENSE)
[![Build Status](http://img.shields.io/travis/usk81/base65536.svg?style=flat-square)](https://travis-ci.org/DMMcomLabo/dmm-go-sdk)
[![Coverage Status](https://img.shields.io/coveralls/usk81/base65536.svg?style=flat-square)](https://coveralls.io/github/DMMcomLabo/dmm-go-sdk?branch=master)

DMM Web API version.3 Client for Go

see: [DMM Affiliate](https://affiliate.dmm.com/)

## Installation

Standard `go get`:

```
$ go get github.com/DMMcomLabo/dmm-go-sdk
```

## Usage

For usage and examples see the [Godoc](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk).

## Example (e.g Product API)

```
package (
    "fmt"
    "github.com/DMMcomLabo/dmm-go-sdk"
)

client := dmm.New("foobarbazbuzz", "dummy-990")
api := client.Product
api.SetSite(SITE_ALLAGES)
api.SetService("mono")
api.SetFloor("dvd")
api.SetSort("date")
api.SetLength(1)
result, err := api.Execute()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}
```

OR

```
package (
    "fmt"
    "github.com/DMMcomLabo/dmm-go-sdk/api"
)

rst, err := NewProductService("foobarbazbuzz", "dummy-999").SetSite(SITE_ADULT).SetLength(1).Execute()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(rst)
}
```

# Details

see [Godoc](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk/api) (English and Japanese) or [our documentation](https://github.com/DMMcomLabo/dmm-go-sdk/blob/master/docs/README.md) (Japanese ONLY)