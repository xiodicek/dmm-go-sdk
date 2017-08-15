package api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const (
	// DefaultActressAPILength is Default length for a Actress API request
	DefaultActressAPILength = 20
	// DefaultActressMaxLength is MAX length for a Actress API request
	DefaultActressMaxLength = 100
)

type ActressService struct {
	ApiID       string `mapstructure:"api_id"`
	AffiliateID string `mapstructure:"affiliate_id"`
	Length      int64  `mapstructure:"hits"`
	Offset      int64  `mapstructure:"offset"`
	Sort        string `mapstructure:"sort"`
	Initial     string `mapstructure:"initial"`
	Keyword     string `mapstructure:"keyword"`
	Bust        string `mapstructure:"bust"`
	Waist       string `mapstructure:"waist"`
	Hip         string `mapstructure:"hip"`
	Height      string `mapstructure:"height"`
	Birthday    string `mapstructure:"birthday"`
}

type ActressRawResponse struct {
	Request ActressService  `mapstructure:"request"`
	Result  ActressResponse `mapstructure:"result"`
}

type ActressResponse struct {
	ResultCount   int64     `mapstructure:"result_count"`
	TotalCount    int64     `mapstructure:"total_count"`
	FirstPosition int64     `mapstructure:"first_position"`
	Actresses     []Actress `mapstructure:"actress"`
}

type Actress struct {
	ID          string             `mapstructure:"id"`
	Name        string             `mapstructure:"name"`
	Ruby        string             `mapstructure:"ruby"`
	Bust        string             `mapstructure:"bust"`
	Cup         string             `mapstructure:"cup"`
	Waist       string             `mapstructure:"waist"`
	Hip         string             `mapstructure:"hip"`
	Height      string             `mapstructure:"height"`
	Birthday    string             `mapstructure:"birthday"`
	BloodType   string             `mapstructure:"blood_type"`
	Hobby       string             `mapstructure:"hobby"`
	Prefectures string             `mapstructure:"prefectures"`
	ImageURL    ActressImageList   `mapstructure:"imageURL"`
	ListURL     ActressProductList `mapstructure:"listURL"`
}

type ActressImageList struct {
	Small string `mapstructure:"small"`
	Large string `mapstructure:"large"`
}

type ActressProductList struct {
	Digital string `mapstructure:"digital"`
	Mono    string `mapstructure:"mono"`
	Monthly string `mapstructure:"monthly"`
	Ppm     string `mapstructure:"ppm"`
	Rental  string `mapstructure:"rental"`
}

// NewActressService returns a new service for the given affiliate ID and API ID.
//
// NewActressServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewActressService(affiliateID, apiID string) *ActressService {
	return &ActressService{
		ApiID:       apiID,
		AffiliateID: affiliateID,
		Length:      DefaultActressAPILength,
		Offset:      DefaultAPIOffset,
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

// Execute requests a url is created by BuildRequestURL.
// Use ExecuteWeak If you want get this response in interface{}.
//
// BuildRequestURLで生成したURLにリクエストします。
// もし interface{} でこのレスポンスを取得したい場合は ExecuteWeak を使用してください。
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

// ExecuteWeak requests a url is created by BuildRequestURL.
//
// BuildRequestURLで生成したURLにリクエストします。
func (srv *ActressService) ExecuteWeak() (interface{}, error) {
	reqURL, err := srv.BuildRequestURL()
	if err != nil {
		return nil, err
	}

	return RequestJSON(reqURL)
}

// SetLength set the specified argument to ProductService.Length
//
// SetLengthはLengthパラメータを設定します。
func (srv *ActressService) SetLength(length int64) *ActressService {
	srv.Length = length
	return srv
}

// SetHits set the specified argument to ProductService.Length
//  SetHits is the alias for SetLength
//
// SetHitsはLengthパラメータを設定します。
func (srv *ActressService) SetHits(length int64) *ActressService {
	srv.SetLength(length)
	return srv
}

// SetOffset set the specified argument to ProductService.Offset
//
// SetOffsetはOffsetパラメータを設定します。
func (srv *ActressService) SetOffset(offset int64) *ActressService {
	srv.Offset = offset
	return srv
}

// SetKeyword set the specified argument to ActressService.Keyword
//
// SetKeywordはKeywordパラメータを設定します。
func (srv *ActressService) SetKeyword(keyword string) *ActressService {
	srv.Keyword = TrimString(keyword)
	return srv
}

// SetSort set the specified argument to ActressService.Sort
//
// SetSortはsortパラメータを設定します。
func (srv *ActressService) SetSort(sort string) *ActressService {
	srv.Sort = TrimString(sort)
	return srv
}

// SetInitial sets the specified argument to ActressService.Initial.
// This argment is actress name's initial and you can use only hiragana.
//  e.g. srv.SetInitial("あ") -> Sora Aoi(あおい そら, 蒼井そら)
//
// SetInitialはInitalパラメータに検索したい女優の頭文字をひらがなで設定します。
func (srv *ActressService) SetInitial(initial string) *ActressService {
	srv.Initial = TrimString(initial)
	return srv
}

// SetBirthday sets the specified argument to ActressService.Birthday.
//  format YYYYMMDD
//  e.g. 1999/01/01 -> 19990101
//
// SetBirthdayはBirthdayパラメータに女優の誕生日を設定します。
func (srv *ActressService) SetBirthday(birthday string) *ActressService {
	srv.Birthday = TrimString(birthday)
	return srv
}

// SetBust sets the specified argument (numeric format string) to ActressService.Bust. unit: centimeter.
//
// SetBustはBustパラメータに女優のバストサイズを設定します。
func (srv *ActressService) SetBust(bust string) *ActressService {
	srv.Bust = TrimString(bust)
	return srv
}

// SetWaist sets the specified argument (numeric format string) to ActressService.Waist. unit: centimeter.
//
// SetWaistはBirthdayパラメータに女優のウエストサイズを設定します。
func (srv *ActressService) SetWaist(waist string) *ActressService {
	srv.Waist = TrimString(waist)
	return srv
}

// SetHip sets the specified argument (numeric format string) to ActressService.Hip. unit: centimeter.
//
// SetHipはBirthdayパラメータに女優のヒップサイズを設定します。
func (srv *ActressService) SetHip(hip string) *ActressService {
	srv.Hip = TrimString(hip)
	return srv
}

// SetHeight sets the specified argument (numeric format string) to ActressService.Height. unit: centimeter.
//
// SetHeightはBirthdayパラメータに女優の身長を設定します。
func (srv *ActressService) SetHeight(height string) *ActressService {
	srv.Height = TrimString(height)
	return srv
}

// ValidateLength validates ProductService.Length within the range (1 <= value <= DEFAULT_ACTRESS_MAX_LENGTH).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateLengthはProductService.Lengthが範囲内(1 <= value <= DEFAULT_ACTRESS_MAX_LENGTH)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *ActressService) ValidateLength() bool {
	return ValidateRange(srv.Length, 1, DefaultActressMaxLength)
}

// ValidateOffset validates ActressService.Offset within the range (1 <= value).
//
// ValidateOffsetはActressService.Offsetが範囲内(1 <= value)にあるか検証します。
func (srv *ActressService) ValidateOffset() bool {
	return srv.Offset >= 1
}

// BuildRequestURL creates url to request actress API.
//
// BuildRequestURLは女優検索APIにリクエストするためのURLを作成します。
func (srv *ActressService) BuildRequestURL() (string, error) {
	if srv.ApiID == "" {
		return "", fmt.Errorf("set invalid ApiID parameter")
	}

	if !ValidateAffiliateID(srv.AffiliateID) {
		return "", fmt.Errorf("set invalid AffiliateID parameter")
	}

	queries := url.Values{}
	queries.Set("api_id", srv.ApiID)
	queries.Set("affiliate_id", srv.AffiliateID)

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
	if srv.Sort != "" {
		queries.Set("sort", srv.Sort)
	}
	if srv.Keyword != "" {
		queries.Set("keyword", srv.Keyword)
	}
	if srv.Birthday != "" {
		queries.Set("birthday", srv.Birthday)
	}
	if srv.Bust != "" {
		queries.Set("bust", srv.Bust)
	}
	if srv.Waist != "" {
		queries.Set("waist", srv.Waist)
	}
	if srv.Hip != "" {
		queries.Set("hip", srv.Hip)
	}
	if srv.Height != "" {
		queries.Set("height", srv.Height)
	}

	u, err := buildAPIEndpoint("ActressSearch")
	if err != nil {
		return "", err
	}
	u.RawQuery = queries.Encode()

	return u.String(), nil
}
