package api

import (
	"net/url"
	"strconv"
	"testing"
)

func TestNewAuthorService(t *testing.T) {
	affiliate_id := Dummy_Affliate_Id
	api_id := Dummy_Api_Id

	srv := NewAuthorService(affiliate_id, api_id)
	if srv.AffiliateId != affiliate_id {
		t.Fatalf("AuthorService.AffiliateId is expected to equal the input value(affiliate_id)")
	}

	if srv.ApiId != api_id {
		t.Fatalf("AuthorService.ApiId is expected to equal the input value(api_id)")
	}
}

func TestSetLengthInAuthorService(t *testing.T) {
	srv := dummyAuthorService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("AuthorService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInAuthorService(t *testing.T) {
	srv := dummyAuthorService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("AuthorService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInAuthorService(t *testing.T) {
	srv := dummyAuthorService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("AuthorService.Offset is expected to equal the input value(offset)")
	}
}

func TestValidateLengthInAuthorService(t *testing.T) {
	srv := dummyAuthorService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("AuthorService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_API_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("AuthorService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("AuthorService.ValidateLength is expected TRUE.")
	}

	target = DEFAULT_MAX_LENGTH + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("AuthorService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("AuthorService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("AuthorService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInAuthorService(t *testing.T) {
	srv := dummyAuthorService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("AuthorService.ValidateOffset is expected TRUE.")
	}

	target = 100
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("AuthorService.ValidateOffset is expected TRUE.")
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("AuthorService.ValidateOffset is expected FALSE.")
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("AuthorService.ValidateOffset is expected FALSE.")
	}
}

func TestBuildRequestUrlInAuthorService(t *testing.T) {
	var srv *AuthorService
	var u string
	var err error
	var expected string

	srv = dummyAuthorService()
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/AuthorSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40" + "&hits=" + strconv.FormatInt(DEFAULT_API_LENGTH, 10) + "&offset=" + strconv.FormatInt(DEFAULT_API_OFFSET, 10)
	if u != expected {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("AuthorService.BuildRequestUrl is not expected to have error")
	}

	srv = dummyAuthorService()
	srv.SetLength(0)
	srv.SetOffset(0)
	srv.SetFloorId("40")
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/AuthorSearch?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id + "&floor_id=40"
	expected_base := expected
	if u != expected {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("AuthorService.BuildRequestUrl is not expected to have error")
	}

	srv.SetLength(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("AuthorService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("AuthorService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetFloorId("")
	u, err = srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("AuthorService.BuildRequestUrl is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to return error.")
	}
	srv.SetFloorId("40")

	srv.SetInitial("あ")
	expected = expected_base + "&initial=" + url.QueryEscape("あ")
	u, err = srv.BuildRequestUrl()
	if u != expected {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("AuthorService.BuildRequestUrl is not expected to have error")
	}
	srv.SetInitial("")
}

func TestBuildRequestUrlWithoutApiIdInAuthorService(t *testing.T) {
	srv := dummyAuthorService()
	srv.ApiId = ""
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("AuthorService.BuildRequestUrl is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to return error.")
	}
}

func TestBuildRequestUrlWithWrongAffiliateIdInAuthorService(t *testing.T) {
	srv := dummyAuthorService()
	srv.AffiliateId = "fizzbizz-100"
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("AuthorService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("AuthorService.BuildRequestUrl is expected to return error.")
	}
}

func dummyAuthorService() *AuthorService {
	return NewAuthorService(Dummy_Affliate_Id, Dummy_Api_Id)
}
