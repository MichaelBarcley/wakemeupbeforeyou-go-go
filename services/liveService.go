package services

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// GetLiveAbiosData is responsible for getting the current live /series data from Abios API and parses the information based on the input parameter.
func GetLiveAbiosData(s string) string {
	CheckIfTokenIsValid()
	baseURL := "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false&is_postponed=false&access_token=" + AccessToken
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	switch s {
	case "players":
		results := gjson.Get(string(body), "data.#.rosters.#.players")
		return results.Raw
	case "teams":
		results := gjson.Get(string(body), "data.#.rosters.#.teams")
		return results.Raw
	default:
		return string(body)
	}
}
