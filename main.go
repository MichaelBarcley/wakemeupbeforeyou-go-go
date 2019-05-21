package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

var accessToken = ""
var tokenCreationDate = time.Now()

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/token/", tokenHandler)
	http.HandleFunc("/series/live", liveSeriesHandler)
	http.HandleFunc("/players/live", livePlayersHandler)
	http.HandleFunc("/teams/live", liveTeamsHandler)
	http.ListenAndServe(":8080", nil)
}

// This function will be removed at release, it is used to get token manually if necessary.
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	fmt.Fprintf(w, string(body))

	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	accessToken = m["access_token"].(string)
	tokenCreationDate = time.Now()
	fmt.Println(accessToken)
}

func getToken() {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	fmt.Println("Your access token for the next hour is: ", m["access_token"])
	accessToken = m["access_token"].(string)
	tokenCreationDate = time.Now()
}

func checkIfTokenIsValid() {
	var timeSinceTokenCreation = time.Now().Sub(tokenCreationDate) / 10e8
	fmt.Println("The age of the token is: ", timeSinceTokenCreation)
	if timeSinceTokenCreation > 3600 || accessToken == "" {
		getToken()
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Abios Gaming test exercise web application!")
}

func liveSeriesHandler(w http.ResponseWriter, r *http.Request) {
	checkIfTokenIsValid()

	baseURL := "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false&is_postponed=false&access_token=" + accessToken
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}

func livePlayersHandler(w http.ResponseWriter, r *http.Request) {
	checkIfTokenIsValid()

	baseURL := "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false&is_postponed=false&access_token=" + accessToken
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	results := gjson.Get(string(body), "data.#.rosters.#.players")

	fmt.Fprintf(w, results.String())
}

func liveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	checkIfTokenIsValid()

	baseURL := "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false&is_postponed=false&access_token=" + accessToken
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	results := gjson.Get(string(body), "data.#.rosters.#.teams")

	fmt.Fprintf(w, results.String())
}
