package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

const (
    API_BASE_URL = "https://api.dmm.com/affiliate/v3/"
    API_VERSION  = "3"
    SITE_ALLAGES = "DMM.com"
    SITE_ADULT   = "DMM.R18"
)

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

// validates affiliate_id
// example value: dummy-999
// affiliate number range: 990 ~ 999
func ValidateAffiliateId(affiliate_id string) bool {
    if affiliate_id == "" {
        return false
    }
    return regexp.MustCompile(`^.+-99[0-9]$`).Match([]byte(affiliate_id))
}

// validates site parameter
func ValidateSite(site string) bool {
    if site == "" {
        return false
    }
    if site != SITE_ALLAGES && site != SITE_ADULT {
        return false
    }
    return true
}

func ValidateRange(target, min, max int64) bool {
    return target >= min && target <= max
}

func GetApiVersion() string {
    return API_VERSION
}