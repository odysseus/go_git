// Functions that deal with individual repos
package git

import (
	"fmt"
)

// Pass-through for GET /repos/:user/:repo
func Repo(user, repo string, token OAuthToken) map[string]interface{} {
	req := NewRequest(fmt.Sprintf("repos/%s/%s", user, repo))
	js := APIRequest(req, token)
	return js[0]
}

// Pass-through for GET /repos/:user/:repo/languages
func RepoLanguages(user, repo string, token OAuthToken) map[string]int {
	fin := make(map[string]int)
	req := NewRequest(fmt.Sprintf("repos/%s/%s/languages", user, repo))
	js := APIRequest(req, token)

	for k, v := range js[0] {
		if val, ok := v.(float64); ok {
			fin[k] += int(val)
		} else {
			panic(fmt.Sprintf("%v Failed to convert", v))
		}
	}

	return fin
}
