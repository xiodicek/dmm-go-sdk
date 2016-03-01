package api

import (
    "fmt"
    "net/url"
    "github.com/mitchellh/mapstructure"
)

type FloorService struct {
    ApiId        string
    AffiliateId  string
}

type FloorRawResponse struct {
    Request FloorService  `mapstructure:"request"`
    Result  FloorResponse `mapstructure:"result"`
}

type FloorResponse struct {
    Site []Site
}

type Site struct {
    Name     string       `mapstructure:"name"`
    Code     string       `mapstructure:"code"`
    Services []DMMService `mapstructure:"service"`
}

type DMMService struct {
    Name   string     `mapstructure:"name"`
    Code   string     `mapstructure:"code"`
    Floors []DMMFloor `mapstructure:"floor"`
}

type DMMFloor struct {
    Id     int64  `mapstructure:"id"`
    Name   string `mapstructure:"name"`
    Code   string `mapstructure:"code"`
}

func NewFloorService(affiliateId, apiId string) *FloorService {
    return &FloorService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}

func (srv *FloorService) Execute() (*FloorResponse, error) {
    result, err := srv.ExecuteWeak()
    if err != nil {
        return nil, err
    }
    var raw FloorRawResponse
    if err = mapstructure.WeakDecode(result, &raw); err != nil {
        return nil, err
    }
    return &raw.Result, nil
}

func (srv *FloorService) ExecuteWeak() (interface{}, error) {
    reqUrl, err := srv.BuildRequestUrl()
    if err != nil {
        return nil, err
    }

    return RequestJson(reqUrl)
}

func (srv *FloorService) BuildRequestUrl() (string, error) {
    if srv.ApiId == "" {
        return "", fmt.Errorf("set invalid ApiId parameter.")
    }
    if !ValidateAffiliateId(srv.AffiliateId) {
        return "", fmt.Errorf("set invalid AffiliateId parameter.")
    }

    queries := url.Values{}
    queries.Set("api_id", srv.ApiId)
    queries.Set("affiliate_id", srv.AffiliateId)

    return API_BASE_URL + "FloorList?" + queries.Encode(), nil
}