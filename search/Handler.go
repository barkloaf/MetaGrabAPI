package search

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/barkloaf/SpotGrabAPI/misc"
)

//Handler func
func Handler(writer http.ResponseWriter, request *http.Request) {
	client, token, err := misc.Auth(misc.Config.ID, misc.Config.Secret)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(401), 401)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")

	search, err := GetSearch(client, token, request.URL.RawQuery, misc.Config.SearchLimit)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(500), 500)
	}

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(search)
}
