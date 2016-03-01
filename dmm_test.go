package dmm

import (
    "testing"
)

func TestNew(t *testing.T) {
    affiliate_id := "foobar-999"
    api_id       := "TXpEZ5D4T2xB3J5cuSLf"
    
    c := New(affiliate_id, api_id)

    if c.AffiliateId != affiliate_id {
        t.Fatalf("Client.AffiliateId is expected to equal the affiliate_id parameter")
    }
    if c.ApiId != api_id {
        t.Fatalf("Client.ApiId is expected to equal the api_id parameter")
    }

    if c.Actress.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Actress.AffiliateId")
    }
    if c.Actress.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Actress.ApiId")
    }

    if c.Author.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Author.AffiliateId")
    }
    if c.Author.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Author.ApiId")
    }

    if c.Floor.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Floor.AffiliateId")
    }
    if c.Floor.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Floor.ApiId")
    }

    if c.Genre.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Genre.AffiliateId")
    }
    if c.Genre.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Genre.ApiId")
    }

    if c.Maker.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Maker.AffiliateId")
    }
    if c.Maker.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Maker.ApiId")
    }

    if c.Product.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Product.AffiliateId")
    }
    if c.Product.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Product.ApiId")
    }

    if c.Series.AffiliateId != c.AffiliateId {
        t.Fatalf("Client.AffiliateId is expected to equal Client.Series.AffiliateId")
    }
    if c.Series.ApiId != c.ApiId {
        t.Fatalf("Client.ApiId is expected to equal Client.Series.ApiId")
    }
}
