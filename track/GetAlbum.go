package track

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/barkloaf/MetaGrabAPI/misc"
)

const (
	dayInFormat  = "2006-01-02"
	dayOutFormat = "2-Jan-2006"

	monthInFormat  = "2006-01"
	monthOutFormat = "Jan-2006"
)

//GetAlbum func
func GetAlbum(client *http.Client, token misc.Token, id string) (misc.Album, error) {
	get, err := http.NewRequest("GET", "https://api.spotify.com/v1/albums/"+id, nil)
	if err != nil {
		return misc.Album{}, err
	}

	get.Header.Set("Authorization", "Bearer "+token.Token)

	response, err := client.Do(get)
	if err != nil {
		return misc.Album{}, err
	}

	var album misc.Album
	err = json.NewDecoder(response.Body).Decode(&album)
	if err != nil {
		return misc.Album{}, err
	}

	switch album.DatePrec {
	case "month":
		date, err := time.Parse(monthInFormat, album.Date)
		if err != nil {
			return misc.Album{}, err
		}

		album.Date = date.Format(monthOutFormat)

	case "day":
		date, err := time.Parse(dayInFormat, album.Date)
		if err != nil {
			return misc.Album{}, err
		}

		album.Date = date.Format(dayOutFormat)
	}

	album.Type = strings.Title(album.Type)

	return album, nil
}
