package api

type GenreService struct {
    ApiId        string
    AffiliateId  string
}

func NewGenreService(affiliateId, apiId string) *GenreService {
    return &GenreService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}