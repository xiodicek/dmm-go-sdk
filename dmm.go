package dmm

import (
    "github.com/DMMcomLabo/dmm-go-sdk/api"
)

type Client struct {
    AffiliateId string
    ApiId       string
    
    // services
    Actress *api.ActressService
    Author  *api.AuthorService
    Floor   *api.FloorService
    Genre   *api.GenreService
    Maker   *api.MakerService
    Product *api.ProductService
    Series  *api.SeriesService
}

func New(affiliate_id, api_id string) *Client {
    c := &Client{
        ApiId:       api_id,
        AffiliateId: affiliate_id,
    }
    c.Actress = api.NewActressService(affiliate_id, api_id)
    c.Author  = api.NewAuthorService(affiliate_id, api_id)
    c.Floor   = api.NewFloorService(affiliate_id, api_id)
    c.Genre   = api.NewGenreService(affiliate_id, api_id)
    c.Maker   = api.NewMakerService(affiliate_id, api_id)
    c.Product = api.NewProductService(affiliate_id, api_id)
    c.Series  = api.NewSeriesService(affiliate_id, api_id)
    return c
}
