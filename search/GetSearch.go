package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/barkloaf/MetaGrabAPI/misc"
)

//GetSearch func
func GetSearch(client *http.Client, token misc.Token, query string, limit int) (misc.Search, error) {
	get, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?q="+query+"&type=track&limit="+strconv.Itoa(limit), nil)
	if err != nil {
		return misc.Search{}, err
	}

	get.Header.Set("Authorization", "Bearer "+token.Token)

	response, err := client.Do(get)
	if err != nil {
		return misc.Search{}, err
	}

	var search misc.Search
	err = json.NewDecoder(response.Body).Decode(&search)
	if err != nil {
		fmt.Println(err)
		return misc.Search{}, err
	}

	return search, nil
}
