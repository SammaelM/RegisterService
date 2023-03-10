package storage

type Response struct {
	AccessToken  string
	RefreshToken string
}

type ResponseAPI struct {
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}
