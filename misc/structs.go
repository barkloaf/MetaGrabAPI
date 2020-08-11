package misc

//Info struct
type Info struct {
	Album struct {
		ID     string `json:"id"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
		Name string `json:"name"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"artists"`
	Explicit   bool   `json:"explicit"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Number     int    `json:"track_number"`
	URI        string `json:"uri"`
}

//Album struct
type Album struct {
	ID     string `json:"id"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
	Label    string `json:"label"`
	Name     string `json:"name"`
	Date     string `json:"release_date"`
	DatePrec string `json:"release_date_precision"`
	URI      string `json:"uri"`
}

//Features struct
type Features struct {
	Duration         int     `json:"duration_ms"`
	Key              int     `json:"key"`
	Mode             int     `json:"mode"`
	TimeSig          int     `json:"time_signature"`
	Acousticness     float64 `json:"acousticness"`
	Danceability     float64 `json:"danceability"`
	Energy           float64 `json:"energy"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Speechiness      float64 `json:"speechiness"`
	Valence          float64 `json:"valence"`
	Tempo            float64 `json:"tempo"`
}

//Analysis struct
type Analysis struct {
	Sections []struct {
		Start       float64 `json:"start"`
		Duration    float64 `json:"duration"`
		Conf        float64 `json:"confidence"`
		Loudness    float64 `json:"loudness"`
		Tempo       float64 `json:"tempo"`
		TempoConf   float64 `json:"tempo_confidence"`
		Key         int     `json:"key"`
		KeyConf     float64 `json:"key_confidence"`
		Mode        int     `json:"mode"`
		ModeConf    float64 `json:"mode_confidence"`
		TimeSig     int     `json:"time_signature"`
		TimeSigConf float64 `json:"time_signature_confidence"`
	} `json:"sections"`
}

//Track struct
type Track struct {
	Info     Info
	Album    Album
	Features Features
	Analysis Analysis
}

//Search struct
type Search struct {
	Tracks struct {
		Items []Info `json:"items"`
	} `json:"tracks"`
}

//Token struct
type Token struct {
	Token  string `json:"access_token"`
	Type   string `json:"token_type"`
	Expiry int    `json:"expires_in"`
}
