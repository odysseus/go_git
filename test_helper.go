package git

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	testMux    *http.ServeMux
	testServer *httptest.Server
	testData   map[string]interface{}
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	file, err := os.Open("testdata.json")
	check(err)
	contents, err := ioutil.ReadAll(file)
	check(err)
	err = json.Unmarshal(contents, &testData)
	check(err)
}

func teardown() {
	testServer.Close()
}
