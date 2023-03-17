package request

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"mainmod/internal/storage"
	"net/http"
	"net/url"
)

func ParseFlag() (resullt storage.UrlValues) {
	// flag
	flag.Parse()
	if flag.NArg() < 3 {
		log.Fatal("no enough arguments")
	}

	str := &storage.UrlValues{}

	str.Client_id = flag.Arg(0)
	str.Client_secret = flag.Arg(1)
	str.Redirect_uri = flag.Arg(2)

	return *str
}

func Request(code string) (string, string, error) {

	result := ParseFlag()

	data := url.Values{
		"client_id":     {result.Client_id},
		"client_secret": {result.Client_secret},
		"code":          {code},

		"grant_type": {"authorization_code"},

		"redirect_uri": {result.Redirect_uri},
	}

	resp, err := http.PostForm("https://exbo.net/oauth/token", data)
	if err != nil {
		return "", "", err
	}

	body := &storage.ResponseAPI{}

	if err = json.NewDecoder(resp.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return "", "", err
	}

	defer resp.Body.Close()

	return body.Access_token, body.Refresh_token, nil
}
