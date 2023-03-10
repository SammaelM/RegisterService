package request

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type ResponseAPI struct {
	Token_type    string `json:"token_type"`
	Expires_in    int64  `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type UrlValues struct {
	client_id     string
	client_secret string
	redirect_uri  string
}

func ParseFlag() (resullt UrlValues) {
	// flag
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatal("no enough arguments")
	}

	str := &UrlValues{}

	str.client_id = flag.Arg(0)
	str.client_secret = flag.Arg(1)
	str.redirect_uri = flag.Arg(2)

	return *str
}

func Request(code string) (string, string, error) {

	result := ParseFlag()

	data := url.Values{
		"client_id":     {result.client_id},
		"client_secret": {result.client_secret},
		"code":          {code},

		"grant_type": {"authorization_code"},

		"redirect_uri": {result.redirect_uri},
	}

	resp, err := http.PostForm("https://exbo.net/oauth/token", data)
	if err != nil {
		return "", "", err
	}

	body := &ResponseAPI{}

	if err = json.NewDecoder(resp.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return "", "", err
	}

	defer resp.Body.Close()

	return body.Access_token, body.Refresh_token, nil
}
