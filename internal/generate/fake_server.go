package generate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func DummyHTTP() (s *httptest.Server, err error) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := string(data)
		fmt.Fprintln(w, response)
		w.WriteHeader(http.StatusOK)
	}))

	return server, err
}
