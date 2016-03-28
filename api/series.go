package api

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"strconv"
)

type SeriesService struct {
	ApiId       string `mapstructure:"api_id"`
	AffiliateId string `mapstructure:"affiliate_id"`
	FloorId     string `mapstructure:"floor_id"`
	Initial     string `mapstructure:"initial"`
	Length      int64  `mapstructure:"hits"`
	Offset      int64  `mapstructure:"offset"`
}

type SeriesRawResponse struct {
	Request SeriesService  `mapstructure:"request"`
	Result  SeriesResponse `mapstructure:"result"`
}

type SeriesResponse struct {
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
	SeriesList    []Series `mapstructure:"series"`
}

type Series struct {
	SeriesId string `mapstructure:"series_id"`
	Name     string `mapstructure:"name"`
	Ruby     string `mapstructure:"ruby"`
	ListURL  string `mapstructure:"list_url"`
}

// NewSeriesService returns a new service for the given affiliate ID and API ID.
//
// NewSeriesServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewSeriesService(affiliateId, apiId string) *SeriesService {
	return &SeriesService{
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
func (srv *SeriesService) Execute() (*SeriesResponse, error) {
	result, err := srv.ExecuteWeak()
	if err != nil {
		return nil, err
	}
	var raw SeriesRawResponse
	if err = mapstructure.WeakDecode(result, &raw); err != nil {
		return nil, err
	}
	return &raw.Result, nil
}

// ExecuteWeak requests a url is created by BuildRequestUrl.
//
// BuildRequestUrlで生成したURLにリクエストします。
func (srv *SeriesService) ExecuteWeak() (interface{}, error) {
	reqUrl, err := srv.BuildRequestUrl()
	if err != nil {
		return nil, err
	}

	return RequestJson(reqUrl)
}

// SetLength set the specified argument to SeriesService.Length
//
// SetLengthはLengthパラメータを設定します。
func (srv *SeriesService) SetLength(length int64) *SeriesService {
	srv.Length = length
	return srv
}

// SetHits set the specified argument to SeriesService.Length
//  SetHits is the alias for SetLength
//
// SetHitsはLengthパラメータを設定します。
func (srv *SeriesService) SetHits(length int64) *SeriesService {
	srv.SetLength(length)
	return srv
}

// SetOffset set the specified argument to SeriesService.Offset
//
// SetOffsetはOffsetパラメータを設定します。
func (srv *SeriesService) SetOffset(offset int64) *SeriesService {
	srv.Offset = offset
	return srv
}

// SetInitial sets the specified argument to SeriesService.Initial.
// This argment is series' initial and you can use only hiragana.
//  e.g. srv.SetInitial("あ") -> ARIA(ありあ)
//
// SetInitialはInitalパラメータに検索したいシリーズの頭文字をひらがなで設定します。
func (srv *SeriesService) SetInitial(initial string) *SeriesService {
	srv.Initial = TrimString(initial)
	return srv
}

// SetFloorId sets the specified argument to SeriesService.FloorId.
// You can retrieve Floor IDs from floor API.
//
// SetFloorIdはFloorIdパラメータを設定します。
// フロアIDはフロアAPIから取得できます。
func (srv *SeriesService) SetFloorId(floor_id string) *SeriesService {
	srv.FloorId = TrimString(floor_id)
	return srv
}

// ValidateLength validates SeriesService.Length within the range (1 <= value <= DEFAULT_MAX_LENGTH).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateLengthはSeriesService.Lengthが範囲内(1 <= value <= DEFAULT_MAX_LENGTH)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *SeriesService) ValidateLength() bool {
	return ValidateRange(srv.Length, 1, DEFAULT_MAX_LENGTH)
}

// ValidateOffset validates SeriesService.Offset within the range (1 <= value).
//
// ValidateOffsetはSeriesService.Offsetが範囲内(1 <= value)にあるか検証します。
func (srv *SeriesService) ValidateOffset() bool {
	return srv.Offset >= 1
}

// BuildRequestUrl creates url to request series API.
//
// BuildRequestUrlはシリーズ検索APIにリクエストするためのURLを作成します。
func (srv *SeriesService) BuildRequestUrl() (string, error) {
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

	if srv.Initial != "" {
		queries.Set("initial", srv.Initial)
	}
	return API_BASE_URL + "/SeriesSearch?" + queries.Encode(), nil
}
