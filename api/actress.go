package api

import (
    "fmt"
    "net/url"
    "strconv"
    "github.com/mitchellh/mapstructure"
)

const (
    DEFAULT_ACTRESS_API_LENGTH = 20
    DEFAULT_ACTRESS_MAX_LENGTH = 100
)

type ActressService struct {
    ApiId        string `mapstructure:"api_id"`
    AffiliateId  string `mapstructure:"affiliate_id"`
    Length       int64  `mapstructure:"hits"`
    Offset       int64  `mapstructure:"offset"`
    Sort         string `mapstructure:"sort"`
    Initial      string `mapstructure:"initial"`
    Keyword      string `mapstructure:"keyword"`
    Bust         string `mapstructure:"bust"`
    Waist        string `mapstructure:"waist"`
    Hip          string `mapstructure:"hip"`
    Height       string `mapstructure:"height"`
    Birthday     string `mapstructure:"birthday"`
}

type ActressRawResponse struct {
    Request ActressService  `mapstructure:"request"`
    Result  ActressResponse `mapstructure:"result"`
}

type ActressResponse struct {
    Id          string             `mapstructure:"id"`
    Name        string             `mapstructure:"name"`
    Ruby        string             `mapstructure:"ruby"`
    Bust        string             `mapstructure:"bust"`
    Waist       string             `mapstructure:"waist"`
    Hip         string             `mapstructure:"hip"`
    Height      string             `mapstructure:"height"`
    Birthday    string             `mapstructure:"birthday"`
    Blood_type  string             `mapstructure:"blood_type"`
    Hobby       string             `mapstructure:"hobby"`
    Prefectures string             `mapstructure:"prefectures"`
    ListURL     ActressProductList `mapstructure:"listURL"`
}

type ActressProductList struct {
    Digital string `mapstructure:"digital"`
    Mono    string `mapstructure:"mono"`
    Monthly string `mapstructure:"monthly"`
    Ppm     string `mapstructure:"ppm"`
    Rental  string `mapstructure:"rental"`
}

func NewActressService(affiliateId, apiId string) *ActressService {
    return &ActressService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
        Length:      DEFAULT_ACTRESS_API_LENGTH,
        Offset:      DEFAULT_API_OFFSET,
        Sort:        "",
        Initial:     "",
        Keyword:     "",
        Bust:        "",
        Waist:       "",
        Hip:         "",
        Height:      "",
        Birthday:    "",

    }
}

func (srv *ActressService) Execute() (*ActressResponse, error) {
    result, err := srv.ExecuteWeak()
    if err != nil {
        return nil, err
    }
    var raw ActressRawResponse
    if err = mapstructure.WeakDecode(result, &raw); err != nil {
        return nil, err
    }
    return &raw.Result, nil
}

func (srv *ActressService) ExecuteWeak() (interface{}, error) {
    reqUrl, err := srv.BuildRequestUrl()
    if err != nil {
        return nil, err
    }

    return RequestJson(reqUrl)
}

func (srv *ActressService) SetLength(length int64) *ActressService {
    srv.Length = length
    return srv
}

func (srv *ActressService) SetHits(length int64) *ActressService {
    srv.SetLength(length)
    return srv
}

func (srv *ActressService) SetOffset(offset int64) *ActressService {
    srv.Offset = offset
    return srv
}

func (srv *ActressService) SetKeyword(keyword string) *ActressService {
    srv.Keyword = keyword
    return srv
}

func (srv *ActressService) SetBirthday(birthday string) *ActressService {
    srv.Birthday = birthday
    return srv
}

func (srv *ActressService) SetBust(bust string) *ActressService {
    srv.Bust = bust
    return srv
}

func (srv *ActressService) SetWaist(waist string) *ActressService {
    srv.Waist = waist
    return srv
}

func (srv *ActressService) SetHip(hip string) *ActressService {
    srv.Hip = hip
    return srv
}

func (srv *ActressService) SetHeight(height string) *ActressService {
    srv.Height = height
    return srv
}

func (srv *ActressService) ValidateLength() bool {
    return ValidateRange(srv.Length, 1, DEFAULT_ACTRESS_MAX_LENGTH)
}

func (srv *ActressService) ValidateOffset() bool {
    return srv.Offset > 1
}

func (srv *ActressService) BuildRequestUrl() (string, error) {
    if srv.ApiId == "" {
        return "", fmt.Errorf("set invalid ApiId parameter.")
    }

    if !ValidateAffiliateId(srv.AffiliateId) {
        return "", fmt.Errorf("set invalid AffiliateId parameter.")
    }

    queries := url.Values{}
    queries.Set("api_id", srv.ApiId)
    queries.Set("affiliate_id", srv.AffiliateId)

    if !srv.ValidateLength() {
        return "", fmt.Errorf("length out of range: %d", srv.Length)
    }
    queries.Set("hits", strconv.FormatInt(srv.Length, 10))

    if !srv.ValidateOffset() {
        return "", fmt.Errorf("offset out of range: %d", srv.Offset)
    }
    queries.Set("offset", strconv.FormatInt(srv.Offset, 10))

    if (srv.Initial != "") {
        queries.Set("initial", srv.Initial)
    }
    if (srv.Sort != "") {
        queries.Set("sort", srv.Sort)
    }
    if (srv.Keyword != "") {
        queries.Set("keyword", srv.Keyword)
    }
    if (srv.Birthday != "") {
        queries.Set("birthday", srv.Birthday)
    }
    if (srv.Bust != "") {
        queries.Set("bust", srv.Bust)
    }
    if (srv.Waist != "") {
        queries.Set("waist", srv.Waist)
    }
    if (srv.Hip != "") {
        queries.Set("hip", srv.Hip)
    }
    if (srv.Height != "") {
        queries.Set("height", srv.Height)
    }

    return API_BASE_URL + "/ActressSearch?" + queries.Encode(), nil
}