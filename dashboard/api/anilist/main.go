package main

import (
	"log"
	"net/http"
	"tomtruyen/anilist/internal/anilist"
	"tomtruyen/anilist/internal/config"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	anilistService := anilist.NewService(cfg.AniListClientID, cfg.AniListClientSecret, cfg.AniListUsername)

	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/upcoming", anilistService.GetUpcomingAnimes).Methods("GET")
	router.HandleFunc("/available", anilistService.FetchAvailable).Methods("GET")
	router.HandleFunc("/watching", anilistService.FetchWatching).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
