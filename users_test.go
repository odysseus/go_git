package git

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	//"github.com/stretchr/testify"
)

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	req := NewRequest("users/odysseus")
	req.BaseURI = testServer.URL

	testMux.HandleFunc(req.String(), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(bytes.NewBufferString("Hello"))
	})

	fmt.Println(req)
}
