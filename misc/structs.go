package misc

//Info is catalog information for a single track
type Info struct {
	//Album on which the track appears. This is considered a "simplified album object", and is mostly used to feed the track.GetAlbum function, which returns a misc.Album
	Album struct {
		//The Spotify ID for the album
		ID string `json:"id"`

		//The cover art for the album in various sizes, widest first
		Images []struct {
			//The source URL of the image
			URL string `json:"url"`
		} `json:"images"`

		//The name of the album. In case of an album takedown, the value may be an empty string
		Name string `json:"name"`
	} `json:"album"`

	//Artists who performed the track. This in considered a "simplified artist object"
	Artists []struct {
		//The name of the artist
		Name string `json:"name"`

		//The Spotify URI for the artist
		URI string `json:"uri"`
	} `json:"artists"`

	//Whether or not the track has explicit lyrics (true = yes it does; false = no it does not OR unknown)
	Explicit bool `json:"explicit"`

	//The Spotify ID for the track
	ID string `json:"id"`

	//The name of the track
	Name string `json:"name"`

	//The popularity of the track. The value will be between 0 and 100, with 100 being the most popular
	//The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are
	//Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. Note that the popularity value may lag actual popularity by a few days: the value is not updated in real time
	Popularity int `json:"popularity"`

	//The number of the track. If an album has several discs, the track number is the number on the specified disc
	Number int `json:"track_number"`

	//The Spotify URI for the track
	URI string `json:"uri"`
}

//Album is catalog information for a single album
type Album struct {
	//The Spotify ID for the album
	ID string `json:"id"`

	//The cover art for the album in various sizes, widest first
	Images []struct {
		//The source URL of the image
		URL string `json:"url"`
	} `json:"images"`

	//The music label for the album
	Label string `json:"label"`

	//The name of the album. In case of an album takedown, the value may be an empty string
	Name string `json:"name"`

	//The date the album was first released, formatted depending on the 3 possible values held by DatePrec
	//"year": yyyy (e.g. "1981")
	//"month": MMM-yyyy (e.g. "Dec-1981")
	//"day": dd-MMM-yyyy (e.g. "15-Dec-1981")
	Date     string `json:"release_date"`
	DatePrec string `json:"release_date_precision"`

	//The Spotify URI for the album
	URI string `json:"uri"`
}

//Features contains audio feature information
type Features struct {
	//The duration of the track in milliseconds
	Duration int `json:"duration_ms"`

	//The estimated overall key of the track. Integers map to pitches using standard Pitch Class notation. (e.g. 0 = C, 1 = C♯/D♭, 2 = D, etc.). If no key was detected, the value is -1
	Key int `json:"key"`

	//Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0
	Mode int `json:"mode"`

	//An estimated overall time signature of a track. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of “3/4”, to “7/4”. Spotify only returns a single integer, which may or may not be very useful
	TimeSig int `json:"time_signature"`

	//A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic
	Acousticness float64 `json:"acousticness"`

	//Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable
	Danceability float64 `json:"danceability"`

	//Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy
	Energy float64 `json:"energy"`

	//Predicts whether a track contains no vocals. “Ooh” and “aah” sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly “vocal”. The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0
	Instrumentalness float64 `json:"instrumentalness"`

	//Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live
	Liveness float64 `json:"liveness"`

	//The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typical range between -60 and 0 db
	Loudness float64 `json:"loudness"`

	//Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks
	Speechiness float64 `json:"speechiness"`

	//A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry)
	Valence float64 `json:"valence"`

	//The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration
	Tempo float64 `json:"tempo"`
}

//Analysis is detailed low-level audio analysis for a single track, describing the track's structure and musical content
type Analysis struct {
	//Sections are defined by large variations in rhythm or timbre
	Sections []struct {

		//The starting point (in seconds) of the time interval
		Start float64 `json:"start"`

		//The duration (in seconds) of the time interval
		Duration float64 `json:"duration"`

		//The confidence, from 0.0 to 1.0, of the reliability of the interval
		Conf float64 `json:"confidence"`

		//The overall loudness of the section in decibels (dB)
		Loudness float64 `json:"loudness"`

		//The overall estimated tempo of the section in beats per minute (BPM)
		//The confidence, from 0.0 to 1.0, is the reliability of the tempo. Some tracks contain tempo changes or sounds which don’t contain tempo (like pure speech) which would correspond to a low value in this field
		Tempo     float64 `json:"tempo"`
		TempoConf float64 `json:"tempo_confidence"`

		//The estimated overall key of the section. The values in this field ranging from 0 to 11 mapping to pitches using standard Pitch Class notation (e.g. 0 = C, 1 = C♯/D♭, 2 = D, etc.). If no key was detected, the value is -1
		//The confidence, from 0.0 to 1.0, is the reliability of the key. Songs with many key changes may correspond to low values in this field
		Key     int     `json:"key"`
		KeyConf float64 `json:"key_confidence"`

		//Indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. This field will contain a 0 for “minor”, a 1 for “major”, or a -1 for no result
		//The confidence, from 0.0 to 1.0, is the reliability of the mode
		Mode     int     `json:"mode"`
		ModeConf float64 `json:"mode_confidence"`

		//An estimated overall time signature of a track. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of “3/4”, to “7/4”. Spotify only returns a single integer, which may or may not be very useful
		//The confidence, from 0.0 to 1.0, is the reliability of the time signature. Sections with time signature changes may correspond to low values in this field
		TimeSig     int     `json:"time_signature"`
		TimeSigConf float64 `json:"time_signature_confidence"`
	} `json:"sections"`
}

//Track struct is a struct that contains all of the above structs and is served at the track endpoint
type Track struct {
	Info     Info
	Album    Album
	Features Features
	Analysis Analysis
}

//Search is information about tracks that match a keyword string by the track name or artist. This is returned by the search.GetSearch function, and the maximum number of internal track structs is provided by the limit parameter in said function
type Search struct {
	//Each track returned by the search
	Tracks struct {
		//An Info struct for each track, as defined above
		Items []Info `json:"items"`
	} `json:"tracks"`
}

//Token struct
type Token struct {
	Token  string `json:"access_token"`
	Type   string `json:"token_type"`
	Expiry int    `json:"expires_in"`
}
