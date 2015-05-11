// Functions that combine and summarize the data of several users
package git

// Takes a slice of user handles and returns the total number of repos
// written by them combined
func MultiUserRepoCountTotal(users []string, token OAuthToken) int {
	fin := 0
	for _, user := range users {
		fin += UserRepoCount(user, token)
	}
	return fin
}

// Takes a slice of user handles and returns the total number of repos written
// by them returned as a map where map[username]repo_count
func MultiUserRepoCountMap(users []string, token OAuthToken) map[string]int {
	fin := make(map[string]int)
	for _, user := range users {
		fin[user] = UserRepoCount(user, token)
	}
	return fin
}

// Takes a slice of user handles and returns a map where map[language]bytes_written
// The languages and number of bytes is combined across all the users
func MultiUserLanguageSummary(users []string, token OAuthToken) map[string]int {
	fin := make(map[string]int)
	for _, user := range users {
		langs := UserLanguageSummary(user, token)

		for k, v := range langs {
			fin[k] += v
		}
	}
	return fin
}
