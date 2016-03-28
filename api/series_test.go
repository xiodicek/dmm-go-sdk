package api

import (
	"net/url"
	"strconv"
	"testing"
)

func TestNewSeriesService(t *testing.T) {
	affiliate_id := Dummy_Affliate_Id
	api_id := Dummy_Api_Id

	srv := NewSeriesService(affiliate_id, api_id)
	if srv.AffiliateId != affiliate_id {
		t.Fatalf("SeriesService.AffiliateId is expected to equal the input value(affiliate_id)")
	}

	if srv.ApiId != api_id {
		t.Fatalf("SeriesService.ApiId is expected to equal the input value(api_id)")
	}
}

func TestSetLengthInSeriesService(t *testing.T) {
	srv := dummySeriesService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("SeriesService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInSeriesService(t *testing.T) {
	srv := dummySeriesService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("SeriesService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInSeriesService(t *testing.T) {
	srv := dummySeriesService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("SeriesService.Offset is expected to equal the input value(offset)")
	}
}

func TestValidateLengthInSeriesService(t *testing.T) {
	srv := dummySeriesService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("SeriesService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_API_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("SeriesService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("SeriesService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("SeriesService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("SeriesService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("SeriesService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInSeriesService(t *testing.T) {
	srv := dummySeriesService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("SeriesService.ValidateOffset is expected TRUE.")
	}

	target = 100
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("SeriesService.ValidateOffset is expected TRUE.")
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("SeriesService.ValidateOffset is expected FALSE.")
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("SeriesService.ValidateOffset is expected FALSE.")
	}
}

func TestBuildRequestUrlInSeriesService(t *testing.T) {
	var srv *SeriesService
	var u string
	var err error
	var expected string

	srv = dummySeriesService()
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/SeriesSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40" + "&hits=" + strconv.FormatInt(DEFAULT_API_LENGTH, 10) + "&offset=" + strconv.FormatInt(DEFAULT_API_OFFSET, 10)
	if u != expected {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("SeriesService.BuildRequestUrl is not expected to have error")
	}

	srv = dummySeriesService()
	srv.SetLength(0)
	srv.SetOffset(0)
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/SeriesSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40"
	expected_base := expected
	if u != expected {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("SeriesService.BuildRequestUrl is not expected to have error")
	}

	srv.SetLength(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("SeriesService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("SeriesService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetFloorId("")
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("SeriesService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to return error.")
	}
	srv.SetFloorId("40")

	srv.SetInitial("あ")
	expected = expected_base + "&initial=" + url.QueryEscape("あ")
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("SeriesService.BuildRequestUrl is not expected to have error")
	}
	srv.SetInitial("")
}

func TestBuildRequestUrlWithoutApiIdInSeriesService(t *testing.T) {
	srv := dummySeriesService()
	srv.ApiId = ""
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("SeriesService.BuildRequestUrl is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to return error.")
	}
}

func TestBuildRequestUrlWithWrongAffiliateIdInSeriesService(t *testing.T) {
	srv := dummySeriesService()
	srv.AffiliateId = "fizzbizz-100"
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("SeriesService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("SeriesService.BuildRequestUrl is expected to return error.")
	}
}

func dummySeriesService() *SeriesService {
	return NewSeriesService(Dummy_Affliate_Id, Dummy_Api_Id)
}
