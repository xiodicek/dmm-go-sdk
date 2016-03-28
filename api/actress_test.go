package api

import (
	"net/url"
	"strconv"
	"testing"
)

func TestNewActressService(t *testing.T) {
	affiliate_id := Dummy_Affliate_Id
	api_id := Dummy_Api_Id

	srv := NewActressService(affiliate_id, api_id)
	if srv.AffiliateId != affiliate_id {
		t.Fatalf("ActressService.AffiliateId is expected to equal the input value(affiliate_id)")
	}

	if srv.ApiId != api_id {
		t.Fatalf("ActressService.ApiId is expected to equal the input value(api_id)")
	}
}

func TestSetLengthInActressService(t *testing.T) {
	srv := dummyActressService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("ActressService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInActressService(t *testing.T) {
	srv := dummyActressService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("ActressService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInActressService(t *testing.T) {
	srv := dummyActressService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("ActressService.Offset is expected to equal the input value(offset)")
	}
}

func TestSetKeywordInActressService(t *testing.T) {
	srv := dummyActressService()

	keyword1 := "abcdefghijkelmnopqrstuvwxyzABCDEFGHIJKELMNOPQRSTUVWXYZ0123456789"
	srv.SetKeyword(keyword1)
	if srv.Keyword != keyword1 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword1)")
	}

	keyword2 := ""
	srv.SetKeyword(keyword2)
	if srv.Keyword != keyword2 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword2)")
	}

	keyword3 := "つれづれなるまゝに、日暮らし、硯にむかひて、心にうつりゆくよしなし事を、そこはかとなく書きつくれば、あやしうこそものぐるほしけれ。"
	srv.SetKeyword(keyword3)
	if srv.Keyword != keyword3 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword3)")
	}

	keyword4 := " a b c d 0 "
	keyword4_expected := "a b c d 0"
	srv.SetKeyword(keyword4)
	if srv.Keyword != keyword4_expected {
		t.Fatalf("ActressService.Keyword is expected to equal keyword4_expected")
	}

	keyword5 := "　あ ア　化Ａ "
	keyword5_expected := "あ ア　化Ａ"
	srv.SetKeyword(keyword5)
	if srv.Keyword != keyword5_expected {
		t.Fatalf("ActressService.Keyword is expected to equal keyword5_expected")
	}
}

func TestSetBirthdayInActressService(t *testing.T) {
	srv := dummyActressService()

	date1 := "19840723"
	srv.SetBirthday(date1)
	if srv.Birthday != date1 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date1)")
	}

	date2 := ""
	srv.SetBirthday(date2)
	if srv.Birthday != date2 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date2)")
	}

	date3 := "19000101-20160101"
	srv.SetBirthday(date3)
	if srv.Birthday != date3 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date3)")
	}
}

func TestSetBustInActressService(t *testing.T) {
	srv := dummyActressService()

	bust1 := "D"
	srv.SetBust(bust1)
	if srv.Bust != bust1 {
		t.Fatalf("ActressService.Bust is expected to equal the input value(bust1)")
	}

	bust2 := ""
	srv.SetBust(bust2)
	if srv.Bust != bust2 {
		t.Fatalf("ActressService.Bust is expected to equal the input value(bust2)")
	}
}

func TestSetWaistInActressService(t *testing.T) {
	srv := dummyActressService()

	waist1 := "60"
	srv.SetWaist(waist1)
	if srv.Waist != waist1 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist1)")
	}

	waist2 := ""
	srv.SetWaist(waist2)
	if srv.Waist != waist2 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist2)")
	}

	waist3 := "50-60"
	srv.SetWaist(waist3)
	if srv.Waist != waist3 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist3)")
	}
}

func TestSetHipInActressService(t *testing.T) {
	srv := dummyActressService()

	hip1 := "88"
	srv.SetHip(hip1)
	if srv.Hip != hip1 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip1)")
	}

	hip2 := ""
	srv.SetHip(hip2)
	if srv.Hip != hip2 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip2)")
	}

	hip3 := "80-90"
	srv.SetHip(hip3)
	if srv.Hip != hip3 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip3)")
	}
}

func TestSetHeightInActressService(t *testing.T) {
	srv := dummyActressService()

	height1 := "155"
	srv.SetHeight(height1)
	if srv.Height != height1 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height1)")
	}

	height2 := ""
	srv.SetHeight(height2)
	if srv.Height != height2 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height2)")
	}

	height3 := "140-160"
	srv.SetHeight(height3)
	if srv.Height != height3 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height3)")
	}
}

func TestValidateLengthInActressService(t *testing.T) {
	srv := dummyActressService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_ACTRESS_API_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_ACTRESS_MAX_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_ACTRESS_MAX_LENGTH + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInActressService(t *testing.T) {
	srv := dummyActressService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ActressService.ValidateOffset is expected TRUE.")
	}

	target = 100
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ActressService.ValidateOffset is expected TRUE.")
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ActressService.ValidateOffset is expected FALSE.")
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ActressService.ValidateOffset is expected FALSE.")
	}
}

func TestBuildRequestUrlInActressService(t *testing.T) {
	var srv *ActressService
	var u string
	var err error
	var expected string

	srv = dummyActressService()
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/ActressSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&hits=" + strconv.FormatInt(DEFAULT_ACTRESS_API_LENGTH, 10) + "&offset=" + strconv.FormatInt(DEFAULT_API_OFFSET, 10)
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}

	srv = dummyActressService()
	srv.SetLength(0)
	srv.SetOffset(0)
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/ActressSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id
	expected_base := expected
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}

	srv.SetLength(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestUrl is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestUrl is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetInitial("あ")
	expected = expected_base + "&initial=" + url.QueryEscape("あ")
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetInitial("")

	srv.SetSort("name")
	expected = expected_base + "&sort=name"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetSort("")

	srv.SetKeyword("天使もえ")
	expected = expected_base + "&keyword=" + url.QueryEscape("天使もえ")
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetKeyword("")

	srv.SetBirthday("19940710")
	expected = expected_base + "&birthday=19940710"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetBirthday("")

	srv.SetBust("84")
	expected = expected_base + "&bust=84"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetBust("")

	srv.SetWaist("57")
	expected = expected_base + "&waist=57"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetWaist("")

	srv.SetHip("82")
	expected = expected_base + "&hip=82"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
	srv.SetHip("")

	srv.SetHeight("155")
	expected = expected_base + "&height=155"
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestUrl is not expected to have error")
	}
}

func TestBuildRequestUrlWithoutApiIdInActressService(t *testing.T) {
	srv := dummyActressService()
	srv.ApiId = ""
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestUrl is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestUrl is expected to return error.")
	}
}

func TestBuildRequestUrlWithWrongAffiliateIdInActressService(t *testing.T) {
	srv := dummyActressService()
	srv.AffiliateId = "fizzbizz-100"
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestUrl is expected to return error.")
	}
}

func dummyActressService() *ActressService {
	return NewActressService(Dummy_Affliate_Id, Dummy_Api_Id)
}
