package api

import (
    "testing"
    "net/url"
    "strconv"
)

func TestNewProductService(t *testing.T) {
    affiliate_id := Dummy_Affliate_Id
    api_id       := Dummy_Api_Id

    srv := NewProductService(affiliate_id, api_id)
    if srv.AffiliateId != affiliate_id {
        t.Fatalf("ProductService.AffiliateId is expected to equal the input value(affiliate_id)")
    }

    if srv.ApiId != api_id {
        t.Fatalf("ProductService.ApiId is expected to equal the input value(api_id)")
    }
}

func TestSetLengthInProductService(t *testing.T) {
    srv := dummyProductService()
    var length int64 = 10
    srv.SetLength(length)

    if srv.Length != length {
        t.Fatalf("ProductService.Length is expected to equal the input value(length)")
    }
}

func TestSetHitsInProductService(t *testing.T) {
    srv := dummyProductService()
    var hits int64 = 10
    srv.SetHits(hits)

    if srv.Length != hits {
        t.Fatalf("ProductService.Length is expected to equal the input value(hits)")
    }
}

func TestSetOffsetInProductService(t *testing.T) {
    srv := dummyProductService()
    var offset int64 = 10
    srv.SetOffset(offset)

    if srv.Offset != offset {
        t.Fatalf("ProductService.Offset is expected to equal the input value(offset)")
    }
}

func TestSetKeywordInProductService(t *testing.T) {
    srv := dummyProductService()

    keyword1 := "abcdefghijkelmnopqrstuvwxyzABCDEFGHIJKELMNOPQRSTUVWXYZ0123456789"
    srv.SetKeyword(keyword1)
    if srv.Keyword != keyword1 {
        t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword1)")
    }

    keyword2 := ""
    srv.SetKeyword(keyword2)
    if srv.Keyword != keyword2 {
        t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword2)")
    }

    keyword3 := "つれづれなるまゝに、日暮らし、硯にむかひて、心にうつりゆくよしなし事を、そこはかとなく書きつくれば、あやしうこそものぐるほしけれ。"
    srv.SetKeyword(keyword3)
    if srv.Keyword != keyword3 {
        t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword3)")
    }

    keyword4 := " a b c d 0 "
    keyword4_expected := "a b c d 0"
    srv.SetKeyword(keyword4)
    if srv.Keyword != keyword4_expected {
        t.Fatalf("ProductService.Keyword is expected to equal keyword4_expected")
    }

    keyword5 := "　あ ア　化Ａ "
    keyword5_expected := "あ ア　化Ａ"
    srv.SetKeyword(keyword5)
    if srv.Keyword != keyword5_expected {
        t.Fatalf("ProductService.Keyword is expected to equal keyword5_expected")
    }
}

func TestSetSiteInProductService(t *testing.T) {
    srv := dummyProductService()

    var site string

    site = SITE_ALLAGES
    srv.SetSite(site)
    if srv.Site != site {
        t.Fatalf("ProductService.Site is expected to equal the input value. value:%s", site)
    }

    site = SITE_ADULT
    srv.SetSite(site)
    if srv.Site != site {
        t.Fatalf("ProductService.Site is expected to equal the input value. value:%s", site)
    }
}

func TestSetServiceInProductService(t *testing.T) {
    srv := dummyProductService()

    service := "digital"
    srv.SetService(service)
    if srv.Service != service {
        t.Fatalf("ProductService.Service is expected to equal the input value(service)")
    }
}

func TestSetFloorInProductService(t *testing.T) {
    srv := dummyProductService()

    floor := "videoa"
    srv.SetFloor(floor)
    if srv.Floor != floor {
        t.Fatalf("ProductService.Floor is expected to equal the input value(floor)")
    }
}

func TestValidateLengthInProductService(t *testing.T) {
    srv := dummyProductService()

    var target int64

    target = 1
    srv.SetLength(target)
    if srv.ValidateLength() == false {
        t.Fatalf("ProductService.ValidateLength is expected TRUE.")
    }

    target = DEFAULT_PRODUCT_API_LENGTH
    srv.SetLength(target)
    if srv.ValidateLength() == false {
        t.Fatalf("ProductService.ValidateLength is expected TRUE.")
    }

    target = DEFAULT_PRODUCT_MAX_LENGTH
    srv.SetLength(target)
    if srv.ValidateLength() == false {
        t.Fatalf("ProductService.ValidateLength is expected TRUE.")
    }

    target = DEFAULT_PRODUCT_MAX_LENGTH + 1
    srv.SetLength(target)
    if srv.ValidateLength() == true {
        t.Fatalf("ProductService.ValidateLength is expected FALSE.")
    }

    target = 0
    srv.SetLength(target)
    if srv.ValidateLength() == true {
        t.Fatalf("ProductService.ValidateLength is expected FALSE.")
    }

    target = -1
    srv.SetLength(target)
    if srv.ValidateLength() == true {
        t.Fatalf("ProductService.ValidateLength is expected FALSE.")
    }
}

func TestValidateOffsetInProductService(t *testing.T) {
    srv := dummyProductService()

    var target int64

    target = 1
    srv.SetOffset(target)
    if srv.ValidateOffset() == false {
        t.Fatalf("ProductService.ValidateOffset is expected TRUE. target:%d", target)
    }

    target = DEFAULT_PRODUCT_MAX_OFFSET
    srv.SetOffset(target)
    if srv.ValidateOffset() == false {
        t.Fatalf("ProductService.ValidateOffset is expected TRUE. target:%d", target)
    }

    target = DEFAULT_PRODUCT_MAX_OFFSET + 1
    srv.SetOffset(target)
    if srv.ValidateOffset() == true {
        t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
    }

    target = 0
    srv.SetOffset(target)
    if srv.ValidateOffset() == true {
        t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
    }

    target = -1
    srv.SetOffset(target)
    if srv.ValidateOffset() == true {
        t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
    }
}

func TestBuildRequestUrlInProductService(t *testing.T) {
    var srv *ProductService
    var u string
    var err error
    var expected string

    srv = dummyProductService()
    srv.SetSite(SITE_ADULT)
    u, err = srv.BuildRequestUrl()
    expected = API_BASE_URL + "/ItemList?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&hits=" + strconv.FormatInt(DEFAULT_ACTRESS_API_LENGTH, 10) + "&offset=" + strconv.FormatInt(DEFAULT_API_OFFSET, 10) + "&site=" + SITE_ADULT
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }

    srv = dummyProductService()
    srv.SetSite(SITE_ADULT)
    srv.SetLength(0)
    srv.SetOffset(0)
    u, err = srv.BuildRequestUrl()
    expected = API_BASE_URL + "/ItemList?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id
    expected_base := expected
    expected = expected_base + "&site=" + SITE_ADULT
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }

    srv.SetSite("")
    u, err = srv.BuildRequestUrl()
    if u != "" {
        t.Fatalf("ProductService.BuildRequestUrl is expected empty if error occurs.")
    }
    if err == nil {
        t.Fatalf("ProductService.BuildRequestUrl is expected to return error.")
    }
    srv.SetSite(SITE_ADULT)

    srv.SetLength(-1)
    u, err = srv.BuildRequestUrl()
    if u != "" {
        t.Fatalf("ProductService.BuildRequestUrl is expected empty if error occurs.")
    }
    if err == nil {
        t.Fatalf("ProductService.BuildRequestUrl is expected to return error.")
    }
    srv.SetLength(0)

    srv.SetOffset(-1)
    u, err = srv.BuildRequestUrl()
    if u != "" {
        t.Fatalf("ProductService.BuildRequestUrl is expected empty if error occurs.")
    }
    if err == nil {
        t.Fatalf("ProductService.BuildRequestUrl is expected to return error.")
    }
    srv.SetOffset(0)

    srv.SetSort("rank")
    expected = expected_base + "&site=" + SITE_ADULT + "&sort=rank"
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
    }
    srv.SetSort("")

    srv.SetKeyword("上原亜衣")
    expected = expected_base + "&keyword=" + url.QueryEscape("上原亜衣") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetKeyword("")

    srv.SetService("digital")
    expected = expected_base + "&service=" + url.QueryEscape("digital") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetService("")

    srv.SetFloor("videoa")
    expected = expected_base + "&floor=" + url.QueryEscape("videoa") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetFloor("")

    srv.SetArticle("actress")
    expected = expected_base + "&article=" + url.QueryEscape("actress") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetArticle("")

    srv.SetArticleId("1011199")
    expected = expected_base + "&article_id=" + url.QueryEscape("1011199") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetArticleId("")

    srv.SetStock("mono")
    expected = expected_base + "&mono_stock=" + url.QueryEscape("mono") + "&site=" + SITE_ADULT
    u, err = srv.BuildRequestUrl()
    if u != expected {
        t.Fatalf("ProductService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
    }
    if err != nil {
        t.Fatalf("ProductService.BuildRequestUrl is not expected to have error")
    }
    srv.SetStock("")
}

func TestBuildRequestUrlWithoutApiIdInInProductService(t *testing.T) {
    srv := dummyProductService()
    srv.ApiId = ""
    u, err := srv.BuildRequestUrl()
    if u != "" {
        t.Fatalf("ProductService.BuildRequestUrl is expected empty if API ID is not set.")
    }
    if err == nil {
        t.Fatalf("ProductService.BuildRequestUrl is expected to return error.")
    }
}

func TestBuildRequestUrlWithWrongAffiliateIdInProductService(t *testing.T) {
    srv := dummyProductService()
    srv.AffiliateId = "fizzbizz-100"
    u, err := srv.BuildRequestUrl()
    if u != "" {
        t.Fatalf("ProductService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
    }
    if err == nil {
        t.Fatalf("ProductService.BuildRequestUrl is expected to return error.")
    }
}

func dummyProductService() *ProductService {
    return NewProductService(Dummy_Affliate_Id, Dummy_Api_Id)
}