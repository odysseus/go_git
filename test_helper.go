package git

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	testMux    *http.ServeMux
	testServer *httptest.Server
	testData   map[string][]byte
)

func getToken() *OAuthToken {
	file, err := os.Open(fmt.Sprintf("%s/.github_api_key", os.Getenv("HOME")))
	check(err)
	contents, err := ioutil.ReadAll(file)
	check(err)
	token := OAuthToken(contents)

	return &token
}

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	file, err := os.Open("./testdata/testdata.json")
	check(err)
	contents, err := ioutil.ReadAll(file)
	check(err)
	err = json.Unmarshal(contents, &testData)
	check(err)
}

func teardown() {
	testServer.Close()
}
