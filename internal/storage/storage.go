package storage

type Response struct {
	AccessToken  string
	RefreshToken string
	Err          string
}
type UrlValues struct {
	Client_id     string
	Client_secret string
	Redirect_uri  string
}

type ResponseAPI struct {
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}
type ResponseApiChek struct {
	Information struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		CreationTime string `json:"creationTime"`
	}
	Clan struct {
		Id               string `json:"id"`
		Name             string `json:"name"`
		Tag              string `json:"tag"`
		Level            int    `json:"level"`
		LevelPoints      int    `json:"levelPoints"`
		RegistrationTime string `json:"registrationTime"`
		Alliance         string `json:"alliance"`
		Description      string `json:"description"`
		Leader           string `json:"leader"`
		MemberCount      int    `json:"memberCount"`
	}

	Member struct {
		Name     string `json:"name"`
		Rank     string `json:"rank"`
		JoinTime string `json:"joinTime"`
	}
}
