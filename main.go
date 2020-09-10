package main

import (
	"log"
	"net/http"

	"github.com/barkloaf/MetaGrabAPI/misc"
	"github.com/barkloaf/MetaGrabAPI/search"
	"github.com/barkloaf/MetaGrabAPI/track"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/track", track.Handler)
	mux.HandleFunc("/search", search.Handler)

	log.Fatal(http.ListenAndServe(misc.Config.BindAddress, rateLimit(mux)))
}
