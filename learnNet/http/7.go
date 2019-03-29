package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

//httptest
func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
	//	http.Error(w, "something failed", http.StatusAccepted)
		fmt.Fprintf(w, "%s", `{"tt":"33"}`)
	}
	req, err := http.NewRequest("GET", "http://test-api.zu.ke.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
