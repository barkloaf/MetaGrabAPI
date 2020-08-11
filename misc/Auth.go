package misc

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
)

//Auth func
func Auth(id string, secret string) (*http.Client, Token, error) {
	client := &http.Client{}

	form := make(url.Values)
	form.Add("grant_type", "client_credentials")

	post, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token?"+form.Encode(), nil)
	if err != nil {
		return client, Token{}, err
	}

	post.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(id+":"+secret)))
	post.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(post)
	if err != nil {
		return client, Token{}, err
	}

	var token Token
	err = json.NewDecoder(response.Body).Decode(&token)
	if err != nil {
		return client, token, err
	}

	return client, token, nil
}
