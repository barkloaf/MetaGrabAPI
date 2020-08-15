package track

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

	writer.Header().Set("Access-Control-Allow-Origin", misc.Config.AccessControl)

	id, exist := request.URL.Query()["id"]
	if !exist {
		http.Error(writer, http.StatusText(400), 400)
		return
	}

	info, err := GetInfo(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	album, err := GetAlbum(client, token, info.Album.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	features, err := GetFeatures(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	analysis, err := GetAnalysis(client, token, id[0])
	if err != nil {
		fmt.Println(err)
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	track := misc.Track{
		Info:     info,
		Album:    album,
		Features: features,
		Analysis: analysis,
	}

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(track)
}
