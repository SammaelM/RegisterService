package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mainmod/internal/storage"
	"net/http"
)

func RequestCheckNickname(token, nickname, region string) (status string) {
	//https://eapi.stalcraft.net/region/character/by-name/character/profile

	url := "https://eapi.stalcraft.net/" + region + "/characters"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print("error req on check nickname")
		return "404"
	}

	aut := "Bearer " + token

	log.Print("Aut: ", aut)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", aut)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "404"
	}

	var info []storage.ResponseApiChek

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "404"
	}

	log.Print("Body no js: ", string(body))

	err = json.Unmarshal(body, &info)
	if err != nil {
		return "404"
	}

	log.Print("Body js:", info)

	defer res.Body.Close()

	for _, data := range info {
		if data.Information.Name == nickname {
			return "200"
		}
	}

	return "303"

}
