package api

import (
	"net/url"
	"strconv"
	"testing"
)

func TestNewGenreService(t *testing.T) {
	affiliate_id := Dummy_Affliate_Id
	api_id := Dummy_Api_Id

	srv := NewGenreService(affiliate_id, api_id)
	if srv.AffiliateId != affiliate_id {
		t.Fatalf("GenreService.AffiliateId is expected to equal the input value(affiliate_id)")
	}

	if srv.ApiId != api_id {
		t.Fatalf("GenreService.ApiId is expected to equal the input value(api_id)")
	}
}

func TestSetLengthInGenreService(t *testing.T) {
	srv := dummyGenreService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("GenreService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInGenreService(t *testing.T) {
	srv := dummyGenreService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("GenreService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInGenreService(t *testing.T) {
	srv := dummyGenreService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("GenreService.Offset is expected to equal the input value(offset)")
	}
}

func TestValidateLengthInGenreService(t *testing.T) {
	srv := dummyGenreService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("GenreService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_API_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("GenreService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("GenreService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("GenreService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("GenreService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("GenreService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInGenreService(t *testing.T) {
	srv := dummyGenreService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("GenreService.ValidateOffset is expected TRUE.")
	}

	target = 100
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("GenreService.ValidateOffset is expected TRUE.")
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("GenreService.ValidateOffset is expected FALSE.")
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("GenreService.ValidateOffset is expected FALSE.")
	}
}

func TestBuildRequestUrlInGenreService(t *testing.T) {
	var srv *GenreService
	var u string
	var err error
	var expected string

	srv = dummyGenreService()
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/GenreSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40" + "&hits=" + strconv.FormatInt(DEFAULT_API_LENGTH, 10) + "&offset=" + strconv.FormatInt(DEFAULT_API_OFFSET, 10)
	if u != expected {
		t.Fatalf("GenreService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("GenreService.BuildRequestUrl is not expected to have error")
	}

	srv = dummyGenreService()
	srv.SetLength(0)
	srv.SetOffset(0)
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/GenreSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40"
	expected_base := expected
	if u != expected {
		t.Fatalf("GenreService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("GenreService.BuildRequestUrl is not expected to have error")
	}

	srv.SetLength(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("GenreService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("GenreService.BuildRequestUrl is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("GenreService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("GenreService.BuildRequestUrl is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetFloorId("")
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("GenreService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("GenreService.BuildRequestUrl is expected to return error.")
	}
	srv.SetFloorId("40")

	srv.SetInitial("あ")
	expected = expected_base + "&initial=" + url.QueryEscape("あ")
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("GenreService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("GenreService.BuildRequestUrl is not expected to have error")
	}
	srv.SetInitial("")
}

func TestBuildRequestUrlWithoutApiIdInGenreService(t *testing.T) {
	srv := dummyGenreService()
	srv.ApiId = ""
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("GenreService.BuildRequestUrl is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("GenreService.BuildRequestUrl is expected to return error.")
	}
}

func TestBuildRequestUrlWithWrongAffiliateIdInGenreService(t *testing.T) {
	srv := dummyGenreService()
	srv.AffiliateId = "fizzbizz-100"
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("GenreService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("GenreService.BuildRequestUrl is expected to return error.")
	}
}

func dummyGenreService() *GenreService {
	return NewGenreService(Dummy_Affliate_Id, Dummy_Api_Id)
}
