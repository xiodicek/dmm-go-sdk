package api

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
)

type FloorService struct {
	ApiId       string
	AffiliateId string
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
	Id   int64  `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Code string `mapstructure:"code"`
}

// NewFloorService returns a new service for the given affiliate ID and API ID.
//
// NewFloorServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewFloorService(affiliateId, apiId string) *FloorService {
	return &FloorService{
		ApiId:       apiId,
		AffiliateId: affiliateId,
	}
}

// Execute requests a url is created by BuildRequestUrl.
// Use ExecuteWeak If you want get this response in interface{}.
//
// BuildRequestUrlで生成したURLにリクエストします。
// もし interface{} でこのレスポンスを取得したい場合は ExecuteWeak を使用してください。
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

// ExecuteWeak requests a url is created by BuildRequestUrl.
//
// BuildRequestUrlで生成したURLにリクエストします。
func (srv *FloorService) ExecuteWeak() (interface{}, error) {
	reqUrl, err := srv.BuildRequestUrl()
	if err != nil {
		return nil, err
	}

	return RequestJson(reqUrl)
}

// BuildRequestUrl creates url to request floor API.
//
// BuildRequestUrlはフロアAPIにリクエストするためのURLを作成します。
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

	return API_BASE_URL + "/FloorList?" + queries.Encode(), nil
}
