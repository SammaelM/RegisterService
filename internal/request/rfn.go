package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

	var info []ResponseApiChek

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
