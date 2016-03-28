package api

import (
	"testing"
)

func TestNewFloorService(t *testing.T) {
	affiliate_id := "foobar-999"
	api_id := "TXpEZ5D4T2xB3J5cuSLf"

	srv := NewFloorService(affiliate_id, api_id)
	if srv.AffiliateId != affiliate_id {
		t.Fatalf("FloorService.AffiliateId is expected to equal the input value(affiliate_id)")
	}

	if srv.ApiId != api_id {
		t.Fatalf("FloorService.ApiId is expected to equal the input value(api_id)")
	}
}

func TestBuildRequestUrlInFloorService(t *testing.T) {
	var srv *FloorService
	var u string
	var err error
	var expected string

	srv = dummyFloorService()
	u, err = srv.BuildRequestUrl()
	expected = API_BASE_URL + "/FloorList?affiliate_id=" + Dummy_Affliate_Id + "&api_id=" + Dummy_Api_Id
	if u != expected {
		t.Fatalf("FloorService.BuildRequestUrl is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("FloorService.BuildRequestUrl is not expected to have error")
	}
}

func TestBuildRequestUrlWithoutApiIdInInFloorService(t *testing.T) {
	srv := dummyFloorService()
	srv.ApiId = ""
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("FloorService.BuildRequestUrl is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("FloorService.BuildRequestUrl is expected to return error.")
	}
}

func TestBuildRequestUrlWithWrongAffiliateIdInFloorService(t *testing.T) {
	srv := dummyFloorService()
	srv.AffiliateId = "fizzbizz-100"
	u, err := srv.BuildRequestUrl()
	if u != "" {
		t.Fatalf("FloorService.BuildRequestUrl is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("FloorService.BuildRequestUrl is expected to return error.")
	}
}

func dummyFloorService() *FloorService {
	return NewFloorService(Dummy_Affliate_Id, Dummy_Api_Id)
}
