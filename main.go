package main

import (
	"fmt"
	"net/http"

	"./services"
	"./utilities"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/series/live", liveSeriesHandler)
	mux.HandleFunc("/players/live", livePlayersHandler)
	mux.HandleFunc("/teams/live", liveTeamsHandler)
	http.ListenAndServe(":8081", utilities.Limit(mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Abios Gaming test exercise web application!")
}

func liveSeriesHandler(w http.ResponseWriter, r *http.Request) {
	liveSeries := services.ProvideLiveData("series")
	w.Header().Set("Content-Type", "application/json")
	w.Write(liveSeries)
}

func livePlayersHandler(w http.ResponseWriter, r *http.Request) {
	livePlayers := services.ProvideLiveData("players")
	w.Header().Set("Content-Type", "application/json")
	w.Write(livePlayers)
}

func liveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	liveTeams := services.ProvideLiveData("teams")
	w.Header().Set("Content-Type", "application/json")
	w.Write(liveTeams)
}
