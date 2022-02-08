package misc

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func getEnv(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		err := godotenv.Load()
		if err != nil {
			panic(name + " not set")
		}

		value = os.Getenv(name)
	}

	if value == "" {
		panic(name + " not set")
	}

	return value
}

//Configuration struct
type Configuration struct {
	//ID and Secret are obtained at https://developer.spotify.com/dashboard/
	ID     string // ID
	Secret string // SECRET

	//SearchLimit is an integer that sets the maximum number of search results (type misc.Search) given by the search.GetSearch function
	SearchLimit int // SEARCH_LIMIT

	//RateBucket is the size of the bucket filled at the rate defined by RateLimit, for use for the ratelimit middleware functionality
	RateLimit  time.Duration // RATE_LIMIT
	RateBucket int           // RATE_BUCKET

	//BindAddress is the address in which the API is hosted
	BindAddress string // BIND_ADDRESS

	//AccessControl is a domain that all API access is restricted to in the Access-Control-Allow-Origin header
	AccessControl string // ACCESS_CONTROL
}

//Config Configuration
var Config Configuration

func init() {
	Config.ID = getEnv("ID")
	Config.Secret = getEnv("SECRET")

	sl, err := strconv.Atoi(getEnv("SEARCH_LIMIT"))
	if err != nil {
		panic("Invalid SEARCH_LIMIT")
	}
	Config.SearchLimit = sl

	rl, err := time.ParseDuration(getEnv("RATE_LIMIT"))
	if err != nil {
		panic("Invalid RATE_LIMIT")
	}
	Config.RateLimit = rl

	rb, err := strconv.Atoi(getEnv("RATE_BUCKET"))
	if err != nil {
		panic("Invalid RATE_BUCKET")
	}
	Config.RateBucket = rb

	Config.BindAddress = getEnv("BIND_ADDRESS")

	Config.AccessControl = getEnv("ACCESS_CONTROL")
}
