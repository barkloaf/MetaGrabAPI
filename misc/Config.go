package misc

import (
	"encoding/json"
	"os"

	"golang.org/x/time/rate"
)

//Configuration struct
type Configuration struct {
	ID          string     `json:"id"`
	Secret      string     `json:"secret"`
	SearchLimit int        `json:"searchLimit"`
	RateLimit   rate.Limit `json:"rateLimit"`
	RateBucket  int        `json:"rateBucket"`
	BindAddress string     `json:"bindAddress"`
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
