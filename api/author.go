package api

type AuthorService struct {
    ApiId        string
    AffiliateId  string
}

func NewAuthorService(affiliateId, apiId string) *AuthorService {
    return &AuthorService{
        ApiId:       apiId,
        AffiliateId: affiliateId,
    }
}