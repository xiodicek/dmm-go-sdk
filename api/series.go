package api

type SeriesService struct {
    ApiId        string
    AffiliateId  string
}

func NewSeriesService(affiliateId, apiId string) *SeriesService {
    return &SeriesService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}