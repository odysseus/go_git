package git

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test goes all the way to the API, can be skipped with go test --short
func TestGetUserFull(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping full API test")
	}

	token := getToken()
	handle := "odysseus"
	user := User(handle, token)
	val := user["login"].(string)
	assert.Equal(t, val, handle, "Should be equal")
}

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()
	query := "users/odysseus"

	req := NewRequest(query)
	req.BaseURI = testServer.URL

	testMux.HandleFunc(req.String(), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(bytes.NewBuffer(testData[query]))
	})
}
