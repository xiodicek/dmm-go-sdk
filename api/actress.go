package api

type ActressService struct {
    ApiId        string
    AffiliateId  string
}

func NewActressService(affiliateId, apiId string) *ActressService {
    return &ActressService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}