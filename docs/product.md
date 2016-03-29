# 商品検索API
## 使用例

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

もしくは以下のように1行で書くこともできます。

```
package (
  "fmt"
  "github.com/DMMcomLabo/dmm-go-sdk/api"
)

rst, err := NewProductService("foobarbazbuzz", "dummy-999").SetSite(SiteAdult).SetLength(1).Execute()
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(rst)
}
```

## リクエストパラメータ
APIのパラメータとSDKのパラメータの関連について

| 論理名 | API (物理名) | 必須 | SDK | データ型 |
|---|---|:---:|---|---|
| API ID | api_id | ◯ | ApiID | string |
| アフィリエイトID | affiliate_id | ◯ | AffiliateId | string |
| サイト | site | ◯ | Site | string |
| サービス | service | | Service | string |
| フロア | floor | | Floor | string |
| 取得件数 | hits | | Length | int64 |
| 検索開始位置 | offset | | Offset | int64 |
| ソート順 | sort | | Sort | string |
| キーワード | keyword | | Keyword | string |
| 絞りこみ項目 | article | | Article | string |
| 絞り込みID | article_id | | ArticleID | string |
| 在庫絞り込み | mono_stock | | Stock | string |
