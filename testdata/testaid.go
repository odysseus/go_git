// Script to pull down JSON data from the API to use in test cases
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/odysseus/go_git"
)

func main() {
	// Get the API token
	file, err := os.Open(fmt.Sprintf("%v/.github_api_key", os.Getenv("HOME")))
	check(err)
	contents, err := ioutil.ReadAll(file)
	check(err)
	token := git.OAuthToken(contents)
	file.Close()

	// The map stores the query string with the repsonse it receives
	testCases := make(map[string][]map[string]interface{})

	/// Requests ///

	// GET rate_limit
	req := git.NewRequest("rate_limit")
	testCases[req.Query] = git.APIRequest(req, &token)

	/// Users ///

	// GET :user
	req = git.NewRequest("users/odysseus")
	testCases[req.Query] = git.APIRequest(req, &token)

	// GET :user repos
	req = git.NewRequest("users/odysseus/repos")
	testCases[req.Query] = git.APIRequest(req, &token)[0:3]

	/// Repos ///

	// GET /repos/:user/:repo
	req = git.NewRequest("repos/odysseus/go_git")
	testCases[req.Query] = git.APIRequest(req, &token)

	// GET /repos/:user/:repo/languages
	req = git.NewRequest("repos/odysseus/go_git/languages")
	testCases[req.Query] = git.APIRequest(req, &token)

	/// Orgs ///

	// GET orgs/:org
	req = git.NewRequest("orgs/recursecenter")
	testCases[req.Query] = git.APIRequest(req, &token)

	// GET orgs/:org/members
	req = git.NewRequest("orgs/recursecenter/members")
	testCases[req.Query] = git.APIRequest(req, &token)[0:3]

	// Write the API data to a JSON file
	out, err := os.Create("./testdata.json")
	check(err)
	defer out.Close()

	b, err := json.MarshalIndent(testCases, "", "  ")
	check(err)
	out.Write(b)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
