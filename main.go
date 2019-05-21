package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"./services"
)

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
	services.AccessToken = m["access_token"].(string)
	services.TokenCreationDate = time.Now()
	fmt.Println(services.AccessToken)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Abios Gaming test exercise web application!")
}

func liveSeriesHandler(w http.ResponseWriter, r *http.Request) {
	liveSeries := services.GetLiveAbiosData("series")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, liveSeries)
}

func livePlayersHandler(w http.ResponseWriter, r *http.Request) {
	livePlayers := services.GetLiveAbiosData("players")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, livePlayers)
}

func liveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	liveTeams := services.GetLiveAbiosData("teams")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, liveTeams)
}
