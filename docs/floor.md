## フロアAPI
#### 使用例

```go
package (  
  "fmt"  
  "github.com/DMMcomLabo/dmm-go-sdk"  
)  

client := dmm.New("foobarbazbuzz", "dummy-990")
api := client.Floor
result, err := api.Execute()
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(result)
}
```

もしくは以下のように1行で書くこともできます。

```go
package (
  "fmt"
  "github.com/DMMcomLabo/dmm-go-sdk/api"
)

rst, err := NewFloorService("foobarbazbuzz", "dummy-999").Execute()
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(rst)
}
```

#### リクエストパラメータ
APIのパラメータとSDKのパラメータの関連について

| 論理名 | API (物理名) | 必須 | SDK | データ型 |
|---|---|:---:|---|---|
| API ID | api_id | ◯ | ApiId | string |
| アフィリエイトID | affiliate_id | ◯ | AffiliateId | string |

