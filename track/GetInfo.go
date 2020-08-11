package track

import (
	"encoding/json"
	"net/http"

	"github.com/barkloaf/SpotGrabAPI/misc"
)

//GetInfo func
func GetInfo(client *http.Client, token misc.Token, id string) (misc.Info, error) {
	get, err := http.NewRequest("GET", "https://api.spotify.com/v1/tracks/"+id, nil)
	if err != nil {
		return misc.Info{}, err
	}

	get.Header.Set("Authorization", "Bearer "+token.Token)

	response, err := client.Do(get)
	if err != nil {
		return misc.Info{}, err
	}

	var info misc.Info
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		return misc.Info{}, err
	}

	return info, nil
}
