package api

import (
    "fmt"
    "net/url"
    "strconv"
    "github.com/mitchellh/mapstructure"
)

type MakerService struct {
    ApiId        string `mapstructure:"api_id"`
    AffiliateId  string `mapstructure:"affiliate_id"`
    FloorId      string `mapstructure:"floor_id"`
    Initial      string `mapstructure:"initial"`
    Length       int64  `mapstructure:"hits"`
    Offset       int64  `mapstructure:"offset"`
}

type MakerRawResponse struct {
    Request MakerService  `mapstructure:"request"`
    Result  MakerResponse `mapstructure:"result"`
}

type MakerResponse struct {
    ResultCount   int64   `mapstructure:"result_count"`
    TotalCount    int64   `mapstructure:"total_count"`
    FirstPosition int64   `mapstructure:"first_position"`
    SiteName      string  `mapstructure:"site_name"`
    SiteCode      string  `mapstructure:"site_code"`
    ServiceName   string  `mapstructure:"service_name"`
    ServiceCode   string  `mapstructure:"service_code"`
    FloorId       string  `mapstructure:"floor_id"`
    FloorName     string  `mapstructure:"floor_name"`
    FloorCode     string  `mapstructure:"floor_code"`
    MakerList     []Maker `mapstructure:"maker"`
}

type Maker struct {
    MakerId  string `mapstructure:"maker_id"`
    Name     string `mapstructure:"name"`
    Ruby     string `mapstructure:"ruby"`
    ListURL  string `mapstructure:"list_url"`
}

func NewMakerService(affiliateId, apiId string) *MakerService {
    return &MakerService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
        FloorId:     "",
        Initial:     "",
        Length:      DEFAULT_API_LENGTH,
        Offset:      DEFAULT_API_OFFSET,
    }
}

func (srv *MakerService) Execute() (*MakerResponse, error) {
    result, err := srv.ExecuteWeak()
    if err != nil {
        return nil, err
    }
    var raw MakerRawResponse
    if err = mapstructure.WeakDecode(result, &raw); err != nil {
        return nil, err
    }
    return &raw.Result, nil
}

func (srv *MakerService) ExecuteWeak() (interface{}, error) {
    reqUrl, err := srv.BuildRequestUrl()
    if err != nil {
        return nil, err
    }

    return RequestJson(reqUrl)
}

func (srv *MakerService) SetLength(length int64) *MakerService {
    srv.Length = length
    return srv
}

func (srv *MakerService) SetHits(length int64) *MakerService {
    srv.SetLength(length)
    return srv
}

func (srv *MakerService) SetOffset(offset int64) *MakerService {
    srv.Offset = offset
    return srv
}

func (srv *MakerService) SetInitial(initial string) *MakerService {
    srv.Initial = TrimString(initial)
    return srv
}

func (srv *MakerService) SetFloorId(floor_id string) *MakerService {
    srv.FloorId = TrimString(floor_id)
    return srv
}

func (srv *MakerService) ValidateLength() bool {
    return ValidateRange(srv.Length, 1, DEFAULT_MAX_LENGTH)
}

func (srv *MakerService) ValidateOffset() bool {
    return srv.Offset >= 1
}

func (srv *MakerService) BuildRequestUrl() (string, error) {
    if srv.ApiId == "" {
        return "", fmt.Errorf("set invalid ApiId parameter.")
    }
    if !ValidateAffiliateId(srv.AffiliateId) {
        return "", fmt.Errorf("set invalid AffiliateId parameter.")
    }
    if srv.FloorId == "" {
        return "", fmt.Errorf("set invalid FloorId parameter.")
    }

    queries := url.Values{}
    queries.Set("api_id", srv.ApiId)
    queries.Set("affiliate_id", srv.AffiliateId)
    queries.Set("floor_id", srv.FloorId)

    if srv.Length != 0 {
        if !srv.ValidateLength() {
            return "", fmt.Errorf("length out of range: %d", srv.Length)
        }
        queries.Set("hits", strconv.FormatInt(srv.Length, 10))
    }

    if srv.Offset != 0 {
        if !srv.ValidateOffset() {
            return "", fmt.Errorf("offset out of range: %d", srv.Offset)
        }
        queries.Set("offset", strconv.FormatInt(srv.Offset, 10))
    }

    if (srv.Initial != "") {
        queries.Set("initial", srv.Initial)
    }
    return API_BASE_URL + "/MakerSearch?" + queries.Encode(), nil
}