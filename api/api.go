package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)

const (
    API_BASE_URL       = "https://api.dmm.com/affiliate/v3"
    API_VERSION        = "3"
    SITE_ALLAGES       = "DMM.com"
    SITE_ADULT         = "DMM.R18"
    DEFAULT_API_OFFSET = 1
    DEFAULT_API_LENGTH = 100
    DEFAULT_MAX_LENGTH = 500
)

// RequestJson requests a retirived url and returns the response is parsed JSON-encoded data
//
// RequestJsonは指定されたURLにリクエストしJSONで返ってきたレスポンスをパースしたデータを返します。
func RequestJson(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("Error at API request:%#v", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    var result interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// TrimString wraps up strings.TrimString
//
// TrimStringはstrings.TrimStringのラップ関数です。
func TrimString(str string) (string) {
    return strings.TrimSpace(str)
}

// ValidateAffiliateId validates affiliate ID.
// (affiliate number range: 990 ~ 999)
//  e.g. dummy-999
//
// ValidateAffiliateIdはアフィリエイトID(例: dummy-999)のバリデーションを行います。
//（アフィリエイトの数値の範囲は 990〜999です）
func ValidateAffiliateId(affiliate_id string) bool {
    if affiliate_id == "" {
        return false
    }
    return regexp.MustCompile(`^.+-99[0-9]$`).Match([]byte(affiliate_id))
}

// ValidateSite validates site parameter.
//
// ValidateSiteはsiteパラメータのバリデーションを行います
func ValidateSite(site string) bool {
    if site == "" {
        return false
    }
    if site != SITE_ALLAGES && site != SITE_ADULT {
        return false
    }
    return true
}

// ValidateRange validates a retrieved number within the range ( number >= min && number <= max).
//
// ValidateRangeは指定された数値が最小値と最大値の範囲内にあるかどうか判定します。
func ValidateRange(target, min, max int64) bool {
    return target >= min && target <= max
}

// GetApiVersion returns API version.
//
// GetApiVersionはAPIのバージョンを返します。
func GetApiVersion() string {
    return API_VERSION
}