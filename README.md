<p align="center">
  <a href="https://metagrab.barkloaf.com">
    <img src="https://metagrab.barkloaf.com/logo.png" width="300" />
  </a>
</p>

# <p align="center">MetaGrabAPI</p>
<p align="center">An API that grabs Spotify's metadata for any track you'd like!</p>

## Table of Contents
  - [Introduction](#introduction)
  - [Configuration](#configuration)
  - [Endpoints](#endpoints)
    - [track - `http://bindAddress/track?id={id}`](#track---httpbindaddresstrackidid)
      - [Track Example](#track-example)
    - [search - `http://bindAddress/search?{query}`](#search---httpbindaddresssearchquery)
      - [Search Example](#search-example)
  - [Omissions](#omissions)
  - [License Notice](#license-notice)

## Introduction
This is the API used by my other project, [MetaGrab](https://github.com/barkloaf/MetaGrab). It's written in Go and grabs Spotify's metadata about any track on its platform. Metadata is data about other data. Track metadata can include artists, what album the track appeared on, and the qualitative feel of the track.

## Configuration
[config.example.json](https://github.com/barkloaf/MetaGrabAPI/blob/master/config.example.json) is an example for the config that then becomes `config.json`. The config becomes a struct as such:
```go
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
```
## Endpoints
### track - `http://bindAddress/track?id={id}`
The track enpoint provides a JSON of the track associated with the track ID parameter. The track ID is a base-62 identifier assigned to every track by Spotify and can be found at the end of a Spotify URI. The track struct is the combination of information retrieved by multiple Spotify endpoints, all of which are defined as such:
```go
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

	//A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`

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
	//The type of the album: one of "Album", "Single", or "Compilation".
	Type string `json:"album_type"`

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
```
#### Track Example
Here is an example JSON served by this endpoint with the ID `0g5J7GjfRxsunVrrcDLejQ`:
```json
{
    "Info": {
        "album": {
            "id":"2TIcThPVCssqowwlTKYEpv",
            "images": [
                {
                    "url":"https://i.scdn.co/image/ab67616d0000b273c6b83566875ae494bc4997ab"
                },
                {
                    "url":"https://i.scdn.co/image/ab67616d00001e02c6b83566875ae494bc4997ab"
                },
                {
                    "url":"https://i.scdn.co/image/ab67616d00004851c6b83566875ae494bc4997ab"
                }
            ],
            "name":"Mama Told Me"
        },
        "artists": [
            {
                "name":"Alexus",
                "uri":"spotify:artist:73aRlPLCZSv6wlTUUK0aFP"
            }
        ],
        "available_markets": [
            "AD","AE","AL","AR","AT","AU","BA","BE","BG","BH","BO","BR","BY","CA","CH","CL","CO","CR","CY","CZ","DE","DK","DO","DZ","EC","EE","EG","ES","FI","FR","GB","GR","GT","HK","HN","HR","HU","ID","IE","IL","IN","IS","IT","JO","JP","KW","KZ","LB","LI","LT","LU","LV","MA","MC","MD","ME","MK","MT","MX","MY","NI","NL","NO","NZ","OM","PA","PE","PH","PL","PS","PT","PY","QA","RO","RS","RU","SA","SE","SG","SI","SK","SV","TH","TN","TR","TW","UA","US","UY","VN","XK","ZA"
        ],
        "explicit":false,
        "id":"0g5J7GjfRxsunVrrcDLejQ",
        "name":"Mama Told Me - Original Mix",
        "popularity":19,
        "track_number":2,
        "uri":"spotify:track:0g5J7GjfRxsunVrrcDLejQ"
        },

    "Album": {
        "album_type":"Single",
        "id":"2TIcThPVCssqowwlTKYEpv",
        "images": [
            {
                "url":"https://i.scdn.co/image/ab67616d0000b273c6b83566875ae494bc4997ab"
            },
            {
                "url":"https://i.scdn.co/image/ab67616d00001e02c6b83566875ae494bc4997ab"
            },
            {
                "url":"https://i.scdn.co/image/ab67616d00004851c6b83566875ae494bc4997ab"
            }
        ],
        "label":"Interphase Digital",
        "name":"Mama Told Me",
        "release_date":"2-Nov-2012",
        "release_date_precision":"day",
        "uri":"spotify:album:2TIcThPVCssqowwlTKYEpv"
    },
    
    "Features": {
        "duration_ms":352125,
        "key":3,
        "mode":1,
        "time_signature":4,
        "acousticness":0.00918,
        "danceability":0.474,
        "energy":0.764,
        "instrumentalness":0.263,
        "liveness":0.373,
        "loudness":-7.2,
        "speechiness":0.0533,
        "valence":0.532,
        "tempo":178.931
    },
    
    "Analysis": {
        "sections": [
            {
                "start":0,
                "duration":6.48168,
                "confidence":1,
                "loudness":-24.361,
                "tempo":179.689,
                "tempo_confidence":0.387,
                "key":10,
                "key_confidence":0.301,
                "mode":1,
                "mode_confidence":0.472,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":6.48168,
                "duration":16.43191,
                "confidence":0.771,
                "loudness":-17.123,
                "tempo":178.942,
                "tempo_confidence":0.638,
                "key":0,
                "key_confidence":0.648,
                "mode":0,
                "mode_confidence":0.635,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":22.91359,
                "duration":16.75694,
                "confidence":1,
                "loudness":-7.268,
                "tempo":179.042,
                "tempo_confidence":0.614,
                "key":10,
                "key_confidence":0.532,
                "mode":0,
                "mode_confidence":0.499,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":39.67054,
                "duration":25.48386,
                "confidence":0.422,
                "loudness":-6.589,
                "tempo":178.692,
                "tempo_confidence":0.42,
                "key":3,
                "key_confidence":0.383,
                "mode":1,
                "mode_confidence":0.677,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":65.1544,
                "duration":12.49106,
                "confidence":1,
                "loudness":-19.461,
                "tempo":179.19,
                "tempo_confidence":0.272,
                "key":5,
                "key_confidence":0.398,
                "mode":0,
                "mode_confidence":0.386,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":77.64545,
                "duration":31.24873,
                "confidence":1,
                "loudness":-7.087,
                "tempo":177.343,
                "tempo_confidence":0.312,
                "key":3,
                "key_confidence":0.823,
                "mode":1,
                "mode_confidence":0.763,
                "time_signature":4,
                "time_signature_confidence":0.163
            },
            {
                "start":108.89418,
                "duration":67.03196,
                "confidence":0.795,
                "loudness":-5.841,
                "tempo":178.883,
                "tempo_confidence":0.405,
                "key":5,
                "key_confidence":0.686,
                "mode":0,
                "mode_confidence":0.642,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":175.92615,
                "duration":38.89741,
                "confidence":0.515,
                "loudness":-5.077,
                "tempo":178.892,
                "tempo_confidence":0.395,
                "key":3,
                "key_confidence":0.544,
                "mode":1,
                "mode_confidence":0.489,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":214.82355,
                "duration":18.12729,
                "confidence":1,
                "loudness":-15,
                "tempo":179.619,
                "tempo_confidence":0.399,
                "key":3,
                "key_confidence":0.271,
                "mode":1,
                "mode_confidence":0.557,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":232.95084,
                "duration":18.385,
                "confidence":0.848,
                "loudness":-8.892,
                "tempo":179.292,
                "tempo_confidence":0.342,
                "key":3,
                "key_confidence":0.456,
                "mode":1,
                "mode_confidence":0.536,
                "time_signature":4,
                "time_signature_confidence":0.553
            },
            {
                "start":251.33585,
                "duration":74.75952,
                "confidence":0.858,
                "loudness":-6.648,
                "tempo":179.001,
                "tempo_confidence":0.491,
                "key":5,
                "key_confidence":0.65,
                "mode":0,
                "mode_confidence":0.607,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":326.09537,
                "duration":13.76423,
                "confidence":0.569,
                "loudness":-7.039,
                "tempo":179.032,
                "tempo_confidence":0.653,
                "key":3,
                "key_confidence":0.343,
                "mode":1,
                "mode_confidence":0.582,
                "time_signature":4,
                "time_signature_confidence":1
            },
            {
                "start":339.8596,
                "duration":12.26539,
                "confidence":0.318,
                "loudness":-8.866,
                "tempo":180.903,
                "tempo_confidence":0.236,
                "key":3,
                "key_confidence":0.401,
                "mode":1,
                "mode_confidence":0.647,
                "time_signature":4,
                "time_signature_confidence":0.857
            }
        ]
    }
}
```

### search - `http://bindAddress/search?{query}`
The search enpoint provides a JSON of a Spotify search result from the query in the URL. The query matches to tracks by the track name **and** the artist name. The maximum number of tracks returned by this endpoint can be set by modifying the searchLimit in the config. The search result is defined as such:
```go
//Search is information about tracks that match a keyword string by the track name or artist. This is returned by the search.GetSearch function, and the maximum number of internal track structs is provided by the limit parameter in said function
type Search struct {
	//Each track returned by the search
	Tracks struct {
		//An Info struct for each track, as defined above
		Items []Info `json:"items"`
	} `json:"tracks"`
}
```
#### Search Example
Here is an example JSON served by this endpoint with the query `bark` and a maximum of 3 results:
```json
{
    "tracks": {
        "items": [
            {
                "album": {
                    "id":"537qKeG5gbEvKJpQ4Qmszn",
                    "images": [
                        {
                            "url":"https://i.scdn.co/image/ab67616d0000b273813b607b3b7e094c306af4fd"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00001e02813b607b3b7e094c306af4fd"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00004851813b607b3b7e094c306af4fd"
                        }
                    ],
                    "name":"Bark At The Moon (Expanded Edition)"
                },
                "artists": [
                    {
                        "name":"Ozzy Osbourne","uri":"spotify:artist:6ZLTlhejhndI4Rh53vYhrY"
                    }
                ],
                "available_markets": [
                    "AD","AE","AL","AR","AT","AU","BA","BE","BG","BH","BO","BR","BY","CA","CH","CL","CO","CR","CY","CZ","DE","DK","DO","DZ","EC","EE","EG","ES","FI","FR","GB","GR","GT","HK","HN","HR","HU","ID","IE","IL","IN","IS","IT","JO","JP","KW","KZ","LB","LI","LT","LU","LV","MA","MC","MD","ME","MK","MT","MX","MY","NI","NL","NO","NZ","OM","PA","PE","PH","PL","PS","PT","PY","QA","RO","RS","RU","SA","SE","SG","SI","SK","SV","TH","TN","TR","TW","UA","US","UY","VN","XK","ZA"
                ],
                "explicit":false,
                "id":"2E7W1X4maFFcjHrVrFA7Vs",
                "name":"Bark at the Moon",
                "popularity":66,
                "track_number":1,
                "uri":"spotify:track:2E7W1X4maFFcjHrVrFA7Vs"
            },
            {
                "album": {
                    "id":"0pgmZ0HKZYEIr71WHdaMN7",
                    "images": [
                        {
                            "url":"https://i.scdn.co/image/ab67616d0000b273a0d2b590266ae8b25e5ddc4f"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00001e02a0d2b590266ae8b25e5ddc4f"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00004851a0d2b590266ae8b25e5ddc4f"
                        }
                    ],
                    "name":"Barking"
                }, 
                "artists": [
                    {
                        "name":"Ramz",
                        "uri":"spotify:artist:6ywXRaHY7m2DJ0dd7CsLAB"
                    }
                ],
                "available_markets": [
                    "AD","AE","AL","AR","AT","AU","BA","BE","BG","BH","BO","BR","BY","CA","CH","CL","CO","CR","CY","CZ","DE","DK","DO","DZ","EC","EE","EG","ES","FI","FR","GB","GR","GT","HK","HN","HR","HU","ID","IE","IL","IN","IS","IT","JO","JP","KW","KZ","LB","LI","LT","LU","LV","MA","MC","MD","ME","MK","MT","MX","MY","NI","NL","NO","NZ","OM","PA","PE","PH","PL","PS","PT","PY","QA","RO","RS","RU","SA","SE","SG","SI","SK","SV","TH","TN","TR","TW","UA","US","UY","VN","XK","ZA"
                ],
                "explicit":false,
                "id":"2U5cq89GCnsR1yixKkC8d5",
                "name":"Barking",
                "popularity":70,
                "track_number":1,
                    "uri":"spotify:track:2U5cq89GCnsR1yixKkC8d5"
            },
            {
                "album": {
                    "id":"0bJIHF1Or1YBLFBMwv53K2",
                    "images": [
                        {
                            "url":"https://i.scdn.co/image/ab67616d0000b273cfc4b1939aba562fc97159c5"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00001e02cfc4b1939aba562fc97159c5"
                        },
                        {
                            "url":"https://i.scdn.co/image/ab67616d00004851cfc4b1939aba562fc97159c5"
                        }
                    ],
                    "name":"Hotel Diablo"
                },
                "artists": [
                    {
                        "name":"Machine Gun Kelly",
                        "uri":"spotify:artist:6TIYQ3jFPwQSRmorSezPxX"
                    },
                    {
                        "name":"YUNGBLUD",
			            "uri":"spotify:artist:6Ad91Jof8Niiw0lGLLi3NW"
                    },
                    {
                        "name":"Travis Barker",
			            "uri":"spotify:artist:4exLIFE8sISLr28sqG1qNX"
                    }
                ],
                "available_markets": [
                    "AD","AE","AL","AR","AT","AU","BA","BE","BG","BH","BO","BR","CA","CH","CL","CO","CR","CY","CZ","DE","DK","DO","DZ","EC","EE","EG","ES","FI","FR","GB","GR","GT","HK","HN","HR","HU","ID","IE","IL","IN","IS","IT","JO","JP","KW","KZ","LB","LI","LT","LU","LV","MA","MC","MD","ME","MK","MT","MX","MY","NI","NL","NO","NZ","OM","PA","PE","PH","PL","PS","PT","PY","QA","RO","RS","RU","SA","SE","SG","SI","SK","SV","TH","TN","TR","TW","UA","US","UY","VN","XK","ZA"
                ],
                "explicit":true,
                "id":"2gTdDMpNxIRFSiu7HutMCg",
                "name":"I Think I'm OKAY (with YUNGBLUD \u0026 Travis Barker)",
                "popularity":80,
                "track_number":14,
                "uri":"spotify:track:2gTdDMpNxIRFSiu7HutMCg"
            }
        ]
    }
}
```

## Omissions
There are a number of things shipped by the Spotify API that is not covered with this API. This may be because they're not very useful, or they're confusing to most users and therefore not needed for most contexts this API is designed for. Here is a list of them with explanations if needed:
* Info -
  * Artists -
    * `external_urls`
    * `href`
    * `type` - Literally just the string "artist"
  * `disc_number`
  * `external_ids`
  * `external_urls`
  * `href`
  * Track Relinking (comprising of `is_playable`, `linked_from`, and `restrictions`)
  * `preview_url` - Can be derived easily
* Album -
  * `album_type`
  * `available_markets` - Included in the Info struct
  * `copyrights`
  * `external_ids`
  * `external_urls`
  * `genres` - Spotify documents that genre is shipped with albums and artists. This, however, is a lie. Genre is only shipped with artists. For obvious reasons (genre diversity), this is ridiculous. It isn't worth the hassle to grab all the potentially misleading genres from all (multiple!) artists from each track, so it is not covered.
  * `href`
  * `tracks`
* Features -
  * `analysis_url`
  * `type`
* Analysis -
  * `bars`
  * `beats`
  * `segments`
  * `tatums`

## License Notice
Make sure any and all use of this software complies with the [license](https://github.com/barkloaf/MetaGrabAPI/blob/master/LICENSE), the GNU Affero General Public License v3.0
