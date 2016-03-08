package api

import (
    "testing"
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

func dummyProductService() *ProductService {
    return NewProductService(Dummy_Affliate_Id, Dummy_Api_Id)
}