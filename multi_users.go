// Functions that combine and summarize the data of several users
package main

import ()

func MultiUsersRepoCount(users []string, token OAuthToken) int {
	fin := 0
	for _, user := range users {
		fin += UserRepoCount(user, token)
	}
	return fin
}
