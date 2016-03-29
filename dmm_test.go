package dmm

import (
	"testing"
)

func TestNew(t *testing.T) {
	affiliateID := "foobar-999"
	apiID := "TXpEZ5D4T2xB3J5cuSLf"

	c := New(affiliateID, apiID)

	if c.AffiliateID != affiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal the affiliateID parameter")
	}
	if c.ApiID != apiID {
		t.Fatalf("Client.ApiID is expected to equal the apiID parameter")
	}

	if c.Actress.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Actress.AffiliateID")
	}
	if c.Actress.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Actress.ApiID")
	}

	if c.Author.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Author.AffiliateID")
	}
	if c.Author.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Author.ApiID")
	}

	if c.Floor.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Floor.AffiliateID")
	}
	if c.Floor.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Floor.ApiID")
	}

	if c.Genre.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Genre.AffiliateID")
	}
	if c.Genre.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Genre.ApiID")
	}

	if c.Maker.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Maker.AffiliateID")
	}
	if c.Maker.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Maker.ApiID")
	}

	if c.Product.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Product.AffiliateID")
	}
	if c.Product.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Product.ApiID")
	}

	if c.Series.AffiliateID != c.AffiliateID {
		t.Fatalf("Client.AffiliateID is expected to equal Client.Series.AffiliateID")
	}
	if c.Series.ApiID != c.ApiID {
		t.Fatalf("Client.ApiID is expected to equal Client.Series.ApiID")
	}
}
