package search

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/barkloaf/MetaGrabAPI/misc"
)

//Handler func
func Handler(writer http.ResponseWriter, request *http.Request) {
	client, token, err := misc.Auth(misc.Config.ID, misc.Config.Secret)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", misc.Config.AccessControl)

	search, err := GetSearch(client, token, request.URL.RawQuery, misc.Config.SearchLimit)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(search)
}
