package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type OAuthToken string

func main() {
	file, err := os.Open(fmt.Sprintf("%v/.github_api_key", os.Getenv("HOME")))
	check(err)

	contents, err := ioutil.ReadAll(file)
	check(err)

	token := OAuthToken(contents)
	fmt.Println(token)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Generalized API request function that iterates through subsequent pages
// token:				A string containing a Github OAuth token
// baseRequest:	The API args only, with no leading or trailing slashes
//							eg: "users/octocat/repos"
func APIRequest(token OAuthToken, baseRequest string) []map[string]interface{} {
	page := 0
	done := false
	fin := make([]map[string]interface{}, 0)
	client := &http.Client{Timeout: 5 * time.Second}

	for !done {
		page++

		// Create a request with the OAuth token in the header
		request := fmt.Sprintf("https://api.github.com/%v?page=%v&per_page=%v",
			baseRequest, page, 100)
		req, err := http.NewRequest("GET", request, nil)
		check(err)
		req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

		// Send the request and read the response
		resp, err := client.Do(req)
		check(err)
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		check(err)

		// Parse the response to JSON
		// The API will return either []map[string]interface{} or a single
		// map[string]interface{}, if we get a single item we wrap it in a slice
		// to make the return values consistent across the board
		var js []map[string]interface{}
		err = json.Unmarshal(body, &js)
		if err != nil {
			// If unmarhsaling failed the return value was a single JSON object
			obj := make(map[string]interface{})
			err = json.Unmarshal(body, &obj)
			js = append(js, obj)
		}

		// If that page was less than the page limit we are done
		if len(js) < 100 {
			done = true
		}

		// Append the unmarshaled JSON from this page to the final array
		for _, item := range js {
			fin = append(fin, item)
		}
	}

	return fin
}
