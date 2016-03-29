# DMM SDK for Go (v3)
[![GoDoc](https://img.shields.io/badge/go-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/DMMcomLabo/dmm-go-sdk/blob/master/LICENSE)
[![Build Status](http://img.shields.io/travis/DMMcomLabo/dmm-go-sdk.svg?style=flat-square)](https://travis-ci.org/DMMcomLabo/dmm-go-sdk)
[![Coverage Status](https://img.shields.io/coveralls/DMMcomLabo/dmm-go-sdk.svg?style=flat-square)](https://coveralls.io/github/DMMcomLabo/dmm-go-sdk?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/DMMcomLabo/dmm-go-sdk)](https://goreportcard.com/report/github.com/DMMcomLabo/dmm-go-sdk)

DMM Web API version.3 クライアント

参照: [DMM Affiliate](https://affiliate.dmm.com/)

## インストール

`go get` の場合:

```
$ go get github.com/DMMcomLabo/dmm-go-sdk
```

## 使い方

使い方や使用例はこちらを参照してください。 [Godoc](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk).

## 使用例 (商品検索APIの場合)

```
package (
    "fmt"
    "github.com/DMMcomLabo/dmm-go-sdk"
)

client := dmm.New("foobarbazbuzz", "dummy-990")
api := client.Product
api.SetSite(SiteGeneral)
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

もしくは

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

# 詳細

[Godoc](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk/api) もしくは [our documentation](https://github.com/DMMcomLabo/dmm-go-sdk/blob/master/docs/README.md) を参照してください
