package api

import (
    "fmt"
    "net/url"
    "strconv"
    "github.com/mitchellh/mapstructure"
)

type AuthorService struct {
    ApiId        string `mapstructure:"api_id"`
    AffiliateId  string `mapstructure:"affiliate_id"`
    FloorId      string `mapstructure:"floor_id"`
    Initial      string `mapstructure:"initial"`
    Length       int64  `mapstructure:"hits"`
    Offset       int64  `mapstructure:"offset"`
}

type AuthorRawResponse struct {
    Request AuthorService  `mapstructure:"request"`
    Result  AuthorResponse `mapstructure:"result"`
}

type AuthorResponse struct {
    ResultCount   int64    `mapstructure:"result_count"`
    TotalCount    int64    `mapstructure:"total_count"`
    FirstPosition int64    `mapstructure:"first_position"`
    SiteName      string   `mapstructure:"site_name"`
    SiteCode      string   `mapstructure:"site_code"`
    ServiceName   string   `mapstructure:"service_name"`
    ServiceCode   string   `mapstructure:"service_code"`
    FloorId       string   `mapstructure:"floor_id"`
    FloorName     string   `mapstructure:"floor_name"`
    FloorCode     string   `mapstructure:"floor_code"`
    AuthorList    []Author `mapstructure:"author"`
}

type Author struct {
    AuthorId  string `mapstructure:"author_id"`
    Name      string `mapstructure:"name"`
    Ruby      string `mapstructure:"ruby"`
    ListURL   string `mapstructure:"list_url"`
}

// NewAuthorService returns a new service for the given affiliate ID and API ID.
//
// NewAuthorServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewAuthorService(affiliateId, apiId string) *AuthorService {
    return &AuthorService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
        FloorId:     "",
        Initial:     "",
        Length:      DEFAULT_API_LENGTH,
        Offset:      DEFAULT_API_OFFSET,
    }
}

// Execute requests a url is created by BuildRequestUrl.
// Use ExecuteWeak If you want get this response in interface{}.
//
// BuildRequestUrlで生成したURLにリクエストします。
// もし interface{} でこのレスポンスを取得したい場合は ExecuteWeak を使用してください。
func (srv *AuthorService) Execute() (*AuthorResponse, error) {
    result, err := srv.ExecuteWeak()
    if err != nil {
        return nil, err
    }
    var raw AuthorRawResponse
    if err = mapstructure.WeakDecode(result, &raw); err != nil {
        return nil, err
    }
    return &raw.Result, nil
}

// ExecuteWeak requests a url is created by BuildRequestUrl.
//
// BuildRequestUrlで生成したURLにリクエストします。
func (srv *AuthorService) ExecuteWeak() (interface{}, error) {
    reqUrl, err := srv.BuildRequestUrl()
    if err != nil {
        return nil, err
    }

    return RequestJson(reqUrl)
}

// SetLength set the specified argument to AuthorService.Length
//
// SetLengthはLengthパラメータを設定します。
func (srv *AuthorService) SetLength(length int64) *AuthorService {
    srv.Length = length
    return srv
}

// SetHits sets the specified argument to AuthorService.Length
//  SetHits is the alias for SetLength
//
// SetHitsはLengthパラメータを設定します。
func (srv *AuthorService) SetHits(length int64) *AuthorService {
    srv.SetLength(length)
    return srv
}

// SetOffset sets the specified argument to AuthorService.Offset
//
// SetOffsetはOffsetパラメータを設定します。
func (srv *AuthorService) SetOffset(offset int64) *AuthorService {
    srv.Offset = offset
    return srv
}

// SetInitial sets the specified argument to AuthorService.Initial.
// This argment is author name's initial and you can use only hiragana.
//  e.g. srv.SetInitial("な") -> Soseki Natsume(なつめ そうせき, 夏目漱石)
//
// SetInitialはInitalパラメータに検索したい作者の頭文字をひらがなで設定します。
func (srv *AuthorService) SetInitial(initial string) *AuthorService {
    srv.Initial = TrimString(initial)
    return srv
}

// SetFloorId sets the specified argument to AuthorService.FloorId.
// You can retrieve Floor IDs from floor API.
//
// SetFloorIdはFloorIdパラメータを設定します。
// フロアIDはフロアAPIから取得できます。
func (srv *AuthorService) SetFloorId(floor_id string) *AuthorService {
    srv.FloorId = TrimString(floor_id)
    return srv
}

// ValidateLength validates AuthorService.Length within the range (1 <= value <= DEFAULT_MAX_LENGTH).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateLengthはAuthorService.Lengthが範囲内(1 <= value <= DEFAULT_MAX_LENGTH)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *AuthorService) ValidateLength() bool {
    return ValidateRange(srv.Length, 1, DEFAULT_MAX_LENGTH)
}

// ValidateOffset validates AuthorService.Offset within the range (1 <= value).
//
// ValidateOffsetはAuthorService.Offsetが範囲内(1 <= value)にあるか検証します。
func (srv *AuthorService) ValidateOffset() bool {
    return srv.Offset >= 1
}

// BuildRequestUrl creates url to request author API.
//
// BuildRequestUrlは作者検索APIにリクエストするためのURLを作成します。
func (srv *AuthorService) BuildRequestUrl() (string, error) {
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
    return API_BASE_URL + "/AuthorSearch?" + queries.Encode(), nil
}