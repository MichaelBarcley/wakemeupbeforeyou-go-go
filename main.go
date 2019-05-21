package main

import (
	"fmt"
	"net/http"

	"./services"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/series/live", liveSeriesHandler)
	http.HandleFunc("/players/live", livePlayersHandler)
	http.HandleFunc("/teams/live", liveTeamsHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Abios Gaming test exercise web application!")
}

func liveSeriesHandler(w http.ResponseWriter, r *http.Request) {
	liveSeries := services.ProvideLiveData("series")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, liveSeries)
}

func livePlayersHandler(w http.ResponseWriter, r *http.Request) {
	livePlayers := services.ProvideLiveData("players")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, livePlayers)
}

func liveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	liveTeams := services.ProvideLiveData("teams")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, liveTeams)
}
