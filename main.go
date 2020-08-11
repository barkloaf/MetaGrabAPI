package main

import (
	"log"
	"net/http"

	"github.com/barkloaf/SpotGrabAPI/misc"
	"github.com/barkloaf/SpotGrabAPI/search"
	"github.com/barkloaf/SpotGrabAPI/track"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/track", track.Handler)
	mux.HandleFunc("/search", search.Handler)

	log.Fatal(http.ListenAndServe(misc.Config.BindAddress, rateLimit(mux)))
}
