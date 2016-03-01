package api

type MakerService struct {
    ApiId        string
    AffiliateId  string
}

func NewMakerService(affiliateId, apiId string) *MakerService {
    return &MakerService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}