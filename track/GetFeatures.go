package track

import (
	"encoding/json"
	"net/http"

	"github.com/barkloaf/SpotGrabAPI/misc"
)

//GetFeatures func
func GetFeatures(client *http.Client, token misc.Token, id string) (misc.Features, error) {
	get, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-features/"+id, nil)
	if err != nil {
		return misc.Features{}, err
	}

	get.Header.Set("Authorization", "Bearer "+token.Token)

	response, err := client.Do(get)
	if err != nil {
		return misc.Features{}, err
	}

	var features misc.Features
	err = json.NewDecoder(response.Body).Decode(&features)
	if err != nil {
		return misc.Features{}, err
	}

	return features, nil
}
