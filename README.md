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

## 使用例

```
api := New("foobarbazbuzz", "dummy-990")
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
rst, err := New("foobarbazbuzz", "dummy-999").SetSite(SITE_ADULT).SetLength(1).Execute()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(rst)
}
```

## リクエストパラメータ

>r = *dmm.request

| API | ライブラリ | 説明 | 例 | 設定方法 |
|---|---|---|---|---|
| api_id | ApiId | API ID | "KcZ2ymn6VPufm4XjxFu6" | r := New("KcZ2ymn6VPufm4XjxFu6", "dummy-999") |
| affiliate_id | AffiliateId | アフィリエイト iD | dummy-999 | r := New("foobarbazbuzz", "dummy-999") |
| operation | Operation | API メソッド名 | ItemList | なし (operation は ItemList しかありません) |
| site | Site | サイト名 (DMM.com or DMM.co.jp) | DMM.co.jp | r.SetSite("DMM.com") |
| service | Service | サービス名 | mono | r.SetService("mono") |
| floor | Floor | フロア名 | dvd | r.SetFloor("dvd") |
| hits | Length | 取得件数 | 100 | r.SetLength(100) |
| offset | Offset | 検索開始位置 | 0 | r.SetOffset(0) |
| sort | Sort | 並び替え | rank | r.SetSort("rank") |
| keyword | Keyword | 検索キーワード | social network | r.SetKeyword("social network") |