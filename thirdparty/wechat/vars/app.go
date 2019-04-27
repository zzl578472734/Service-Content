package vars

type WXErrReturn struct {
}

type WXAccessTokenReturn struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expire_in"`
}
