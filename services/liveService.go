package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

var liveSeriesData []byte
var lastAbiosQueryDate = time.Now()

// GetAbiosData is responsible for getting the current live /series data from Abios API.
func GetAbiosData() {
	CheckIfTokenIsValid()
	fmt.Println("Refreshing live series data from Abios...")
	baseURL := "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false&is_postponed=false&access_token=" + accessToken
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	liveSeriesData = body
	lastAbiosQueryDate = time.Now()
}

// ProvideLiveData parses the stored live series data and returns the requested live players/teams/series.
func ProvideLiveData(liveType string) []byte {
	var timeSinceLastQuery = time.Now().Sub(lastAbiosQueryDate) / 10e8
	if timeSinceLastQuery > 90 || liveSeriesData == nil {
		GetAbiosData()
	}

	var results []byte
	switch liveType {
	case "players":
		filterLiveSeriesData("data.#.rosters.#.players", &results)
	case "teams":
		filterLiveSeriesData("data.#.rosters.#.teams", &results)
	default:
		results = liveSeriesData
	}

	return results
}

func filterLiveSeriesData(jsonFilter string, output *[]byte) {
	temp := gjson.GetBytes(liveSeriesData, jsonFilter)
	if temp.Index > 0 {
		*output = liveSeriesData[temp.Index : temp.Index+len(temp.Raw)]
	} else {
		*output = []byte(temp.Raw)
	}
}
