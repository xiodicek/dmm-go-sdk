package api

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"strconv"
)

const (
	DEFAULT_PRODUCT_API_LENGTH = 20
	DEFAULT_PRODUCT_MAX_LENGTH = 100

	DEFAULT_PRODUCT_MAX_OFFSET = 50000
)

type ProductService struct {
	ApiId       string `mapstructure:"api_id"`
	AffiliateId string `mapstructure:"affiliate_id"`
	Site        string `mapstructure:"site"`
	Service     string `mapstructure:"service"`
	Floor       string `mapstructure:"floor"`
	Length      int64  `mapstructure:"hits"`
	Offset      int64  `mapstructure:"offset"`
	Sort        string `mapstructure:"sort"`
	Keyword     string `mapstructure:"keyword"`
	Article     string `mapstructure:"article"`
	ArticleId   string `mapstructure:"article_id"`
	Stock       string `mapstructure:"mono_stock"`
}

type ProductRawResponse struct {
	Request ProductService  `mapstructure:"request"`
	Result  ProductResponse `mapstructure:"result"`
}

type ProductResponse struct {
	ResultCount   int64  `mapstructure:"result_count"`
	TotalCount    int64  `mapstructure:"total_count"`
	FirstPosition int64  `mapstructure:"first_position"`
	Items         []Item `mapstructure:"items"`
}

type Item struct {
	AffiliateUrl       string             `mapstructure:"affiliateURL"`
	AffiliateUrlMobile string             `mapstructure:"affiliateURLsp"`
	CategoryName       string             `mapstructure:"category_name"`
	Comment            string             `mapstructure:"comment"`
	ContentId          string             `mapstructure:"content_id"`
	Date               string             `mapstructure:"date"`
	FloorName          string             `mapstructure:"floor_name"`
	FloorCode          string             `mapstructure:"floor_code"`
	ISBN               string             `mapstructure:"isbn"`
	JANCode            string             `mapstructure:"jancode"`
	ProductCode        string             `mapstructure:"maker_product"`
	ProductId          string             `mapstructure:"product_id"`
	ServiceName        string             `mapstructure:"service_name"`
	ServiceCode        string             `mapstructure:"service_code"`
	Stock              string             `mapstructure:"stock"`
	Title              string             `mapstructure:"title"`
	Url                string             `mapstructure:"URL"`
	UrlMoble           string             `mapstructure:"URLsp"`
	Volume             string             `mapstructure:"volume"`
	ImageUrl           ImageUrlList       `mapstructure:"imageURL"`
	SampleImageUrl     SampleImageUrlList `mapstructure:"sampleImageURL"`
	SampleMovieUrl     SampleMovieUrlList `mapstructure:"sampleMovieURL"`
	Review             ReviewInformation  `mapstructure:"review"`
	PriceInformation   PriceInformation   `mapstructure:"prices"`
	ItemInformation    ItemInformation    `mapstructure:"iteminfo"`
	BandaiInformation  BandaiInformation  `mapstructure:"bandaiinfo"`
	CdInformation      CdInformation      `mapstructure:"cdinfo"`
}

type ImageUrlList struct {
	List  string `mapstructure:"list"`
	Small string `mapstructure:"small"`
	Large string `mapstructure:"large"`
}

type SampleImageUrlList struct {
	Sample_s SmallSampleList `mapstructure:"sample_s"`
}

type SmallSampleList struct {
	Image []string `mapstructure:"image"`
}

type SampleMovieUrlList struct {
	Size_476_306 string `mapstructure:"size_476_306"`
	Size_560_360 string `mapstructure:"size_560_360"`
	Size_644_414 string `mapstructure:"size_644_414"`
	Size_720_480 string `mapstructure:"size_720_480"`
	PC_flag      bool   `mapstructure:"pc_flag"`
	SP_flag      bool   `mapstructure:"sp_flag"`
}

type PriceInformation struct {
	Price         string           `mapstructure:"price"`
	PriceAll      string           `mapstructure:"price_all"`
	RetailPrice   string           `mapstructure:"list_price"`
	Distributions DistributionList `mapstructure:"deliveries"`
}

type DistributionList struct {
	Distribution []Distribution `mapstructure:"delivery"`
}

type Distribution struct {
	Type  string `mapstructure:"type"`
	Price string `mapstructure:"price"`
}

type ItemInformation struct {
	Maker     []ItemComponent `mapstructure:"maker"`
	Label     []ItemComponent `mapstructure:"label"`
	Series    []ItemComponent `mapstructure:"series"`
	Keywords  []ItemComponent `mapstructure:"keyword"`
	Genres    []ItemComponent `mapstructure:"genre"`
	Actors    []ItemComponent `mapstructure:"actor"`
	Artists   []ItemComponent `mapstructure:"artist"`
	Authors   []ItemComponent `mapstructure:"author"`
	Directors []ItemComponent `mapstructure:"director"`
	Fighters  []ItemComponent `mapstructure:"fighter"`
	Colors    []ItemComponent `mapstructure:"color"`
	Sizes     []ItemComponent `mapstructure:"size"`
}

type ItemComponent struct {
	Id   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

type BandaiInformation struct {
	TitleCode string `mapstructure:"titlecode"`
}

type CdInformation struct {
	Kind string `mapstructure:"kind"`
}

type ReviewInformation struct {
	Count   int64   `mapstructure:"count"`
	Average float64 `mapstructure:"average"`
}

// NewProductService returns a new service for the given affiliate ID and API ID.
//
// NewProductServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewProductService(affiliateId, apiId string) *ProductService {
	return &ProductService{
		ApiId:       apiId,
		AffiliateId: affiliateId,
		Site:        "",
		Service:     "",
		Floor:       "",
		Length:      DEFAULT_PRODUCT_API_LENGTH,
		Offset:      DEFAULT_API_OFFSET,
		Sort:        "",
		Keyword:     "",
		Article:     "",
		ArticleId:   "",
		Stock:       "",
	}
}

// Execute requests a url is created by BuildRequestUrl.
// Use ExecuteWeak If you want get this response in interface{}.
//
// BuildRequestUrlで生成したURLにリクエストします。
// もし interface{} でこのレスポンスを取得したい場合は ExecuteWeak を使用してください。
func (srv *ProductService) Execute() (*ProductResponse, error) {
	result, err := srv.ExecuteWeak()
	if err != nil {
		return nil, err
	}
	var raw ProductRawResponse
	if err = mapstructure.WeakDecode(result, &raw); err != nil {
		return nil, err
	}
	return &raw.Result, nil
}

// ExecuteWeak requests a url is created by BuildRequestUrl.
//
// BuildRequestUrlで生成したURLにリクエストします。
func (srv *ProductService) ExecuteWeak() (interface{}, error) {
	reqUrl, err := srv.BuildRequestUrl()
	if err != nil {
		return nil, err
	}

	return RequestJson(reqUrl)
}

// SetLength set the specified argument to ProductService.Length
//
// SetLengthはLengthパラメータを設定します。
func (srv *ProductService) SetLength(length int64) *ProductService {
	srv.Length = length
	return srv
}

// SetHits set the specified argument to ProductService.Length
//  SetHits is the alias for SetLength
//
// SetHitsはLengthパラメータを設定します。
func (srv *ProductService) SetHits(length int64) *ProductService {
	srv.SetLength(length)
	return srv
}

// SetOffset set the specified argument to ProductService.Offset
//
// SetOffsetはOffsetパラメータを設定します。
func (srv *ProductService) SetOffset(offset int64) *ProductService {
	srv.Offset = offset
	return srv
}

// SetKeyword set the specified argument to ProductService.Keyword
//
// SetKeywordはKeywordパラメータを設定します。
func (srv *ProductService) SetKeyword(keyword string) *ProductService {
	srv.Keyword = TrimString(keyword)
	return srv
}

// SetSort set the specified argument to ProductService.Sort
//
// SetSortはSortパラメータを設定します。
func (srv *ProductService) SetSort(sort string) *ProductService {
	srv.Sort = TrimString(sort)
	return srv
}

// SetSite set the specified argument to ProductService.Site
//
// SetSiteはSiteパラメータを設定します。
func (srv *ProductService) SetSite(site string) *ProductService {
	srv.Site = TrimString(site)
	return srv
}

// SetService set the specified argument to ProductService.Service
//
// SetServiceはServiceパラメータを設定します。
func (srv *ProductService) SetService(service string) *ProductService {
	srv.Service = TrimString(service)
	return srv
}

// SetFloor set the specified argument to ProductService.Floor
//
// SetFloorはFloorパラメータを設定します。
func (srv *ProductService) SetFloor(floor string) *ProductService {
	srv.Floor = TrimString(floor)
	return srv
}

// SetArticle set the specified argument to ProductService.Article
//
// SetArticleはArticleパラメータを設定します。
func (srv *ProductService) SetArticle(stock string) *ProductService {
	srv.Article = TrimString(stock)
	return srv
}

// SetArticleId set the specified argument to ProductService.ArticleId
//
// SetArticleIdはArticleIdパラメータを設定します。
func (srv *ProductService) SetArticleId(stock string) *ProductService {
	srv.ArticleId = TrimString(stock)
	return srv
}

// SetStock set the specified argument to ProductService.Stock
//
// SetStockはStockパラメータを設定します。
func (srv *ProductService) SetStock(stock string) *ProductService {
	srv.Stock = TrimString(stock)
	return srv
}

// ValidateLength validates ProductService.Length within the range (1 <= value <= DEFAULT_PRODUCT_MAX_LENGTH).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateLengthはProductService.Lengthが範囲内(1 <= value <= DEFAULT_PRODUCT_MAX_LENGTH)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *ProductService) ValidateLength() bool {
	return ValidateRange(srv.Length, 1, DEFAULT_PRODUCT_MAX_LENGTH)
}

// ValidateOffset validates ProductService.Offset within the range (1 <= value <= DEFAULT_PRODUCT_MAX_OFFSET).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateOffsetはProductService.Offsetが範囲内(1 <= value <= DEFAULT_PRODUCT_MAX_OFFSET)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *ProductService) ValidateOffset() bool {
	return ValidateRange(srv.Offset, 1, DEFAULT_PRODUCT_MAX_OFFSET)
}

// BuildRequestUrl creates url to request product API.
//
// BuildRequestUrlは商品検索APIにリクエストするためのURLを作成します。
func (srv *ProductService) BuildRequestUrl() (string, error) {
	if srv.ApiId == "" {
		return "", fmt.Errorf("set invalid ApiId parameter.")
	}
	if !ValidateAffiliateId(srv.AffiliateId) {
		return "", fmt.Errorf("set invalid AffiliateId parameter.")
	}

	if !ValidateSite(srv.Site) {
		return "", fmt.Errorf("set invalid Site parameter.")
	}

	queries := url.Values{}
	queries.Set("api_id", srv.ApiId)
	queries.Set("affiliate_id", srv.AffiliateId)
	queries.Set("site", srv.Site)

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

	if srv.Service != "" {
		queries.Set("service", srv.Service)
	}
	if srv.Floor != "" {
		queries.Set("floor", srv.Floor)
	}
	if srv.Sort != "" {
		queries.Set("sort", srv.Sort)
	}
	if srv.Keyword != "" {
		queries.Set("keyword", srv.Keyword)
	}
	if srv.Article != "" {
		queries.Set("article", srv.Article)
	}
	if srv.ArticleId != "" {
		queries.Set("article_id", srv.ArticleId)
	}
	if srv.Stock != "" {
		queries.Set("mono_stock", srv.Stock)
	}
	return API_BASE_URL + "/ItemList?" + queries.Encode(), nil
}
