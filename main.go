package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	fmt.Fprintf(w, string(body))

	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	tokenCreationDate = time.Now()
}

func getToken() {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	fmt.Println("Your access token for the next hour is: ", m["access_token"])
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
	fmt.Fprintf(w, "smth random")
}

func livePlayersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "smth random")
}

func liveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "smth random")
}
