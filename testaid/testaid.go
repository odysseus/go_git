// Script to pull down JSON data from the API to use in test cases
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/odysseus/go_git"
)

func main() {
	file, err := os.Open(fmt.Sprintf("%v/.github_api_key", os.Getenv("HOME")))
	check(err)
	contents, err := ioutil.ReadAll(file)
	check(err)
	token := git.OAuthToken(contents)

	fmt.Println(git.RateLimitRemaining(token))

	testCases := make(map[string][]map[string]interface{})

	// Rate Limit JSON
	req := git.NewRequest("rate_limit")
	testCases[req.String(1)] = git.APIRequest(req, token)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
