package track

import (
	"encoding/json"
	"net/http"

	"github.com/barkloaf/SpotGrabAPI/misc"
)

//GetAnalysis func
func GetAnalysis(client *http.Client, token misc.Token, id string) (misc.Analysis, error) {
	get, err := http.NewRequest("GET", "https://api.spotify.com/v1/audio-analysis/"+id, nil)
	if err != nil {
		return misc.Analysis{}, err
	}

	get.Header.Set("Authorization", "Bearer "+token.Token)

	response, err := client.Do(get)
	if err != nil {
		return misc.Analysis{}, err
	}

	var analysis misc.Analysis
	err = json.NewDecoder(response.Body).Decode(&analysis)
	if err != nil {
		return misc.Analysis{}, err
	}

	return analysis, nil
}
