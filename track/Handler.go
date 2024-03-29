package track

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/barkloaf/MetaGrabAPI/misc"
)

// Handler func
func Handler(writer http.ResponseWriter, request *http.Request) {
	client, token, err := misc.Auth(misc.Config.ID, misc.Config.Secret)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", misc.Config.AccessControl)

	id, exist := request.URL.Query()["id"]
	if !exist {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	info, err := GetInfo(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if info.ID == "" {
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	album, err := GetAlbum(client, token, info.Album.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	features, err := GetFeatures(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	analysis, err := GetAnalysis(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	track := misc.Track{
		Info:     info,
		Album:    album,
		Features: features,
		Analysis: analysis,
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(track)
}
