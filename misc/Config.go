package misc

import (
	"encoding/json"
	"os"

	"golang.org/x/time/rate"
)

//Configuration struct
type Configuration struct {
	//ID and Secret are obtained at https://developer.spotify.com/dashboard/
	ID     string `json:"id"`
	Secret string `json:"secret"`

	//SearchLimit is an integer that sets the maximum number of search results (type misc.Search) given by the search.GetSearch function
	SearchLimit int `json:"searchLimit"`

	//RateBucket is the size of the bucket filled at the rate defined by RateLimit, for use for the ratelimit middleware functionality
	RateLimit  rate.Limit `json:"rateLimit"`
	RateBucket int        `json:"rateBucket"`

	//BindAddress is the address in which the API is hosted
	BindAddress string `json:"bindAddress"`

	//AccessControl is a domain that all API access is restricted to in the Access-Control-Allow-Origin header
	AccessControl string `json:"accessControl"`
}

//Config Configuration
var Config Configuration

func init() {
	file, err := os.Open("./config.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	parser := json.NewDecoder(file)
	parser.Decode(&Config)
}
