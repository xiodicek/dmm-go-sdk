package api

import (
    "testing"
)

func TestNewFloorService(t *testing.T) {
    affiliate_id := "foobar-999"
    api_id       := "TXpEZ5D4T2xB3J5cuSLf"

    srv := NewFloorService(affiliate_id, api_id)
    if srv.AffiliateId != affiliate_id {
        t.Fatalf("FloorService.AffiliateId is expected to equal the input value(affiliate_id)")
    }

    if srv.ApiId != api_id {
        t.Fatalf("FloorService.ApiId is expected to equal the input value(api_id)")
    }
}
