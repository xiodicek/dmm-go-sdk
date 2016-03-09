# DMM SDK for Go (v3) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk)
DMM Web API version.3 クライアント

参照: [DMM Affiliate](https://affiliate.dmm.com/)

## インストール

`go get` の場合:

```
$ go get github.com/DMMcomLabo/dmm-go-sdk
```

## 使い方

使い方や使用例はこちらを参考にしてください。 [Godoc](https://godoc.org/github.com/DMMcomLabo/dmm-go-sdk).

## 使用例 (商品検索APIの場合)

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