package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", "aaa")
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
	}))
	defer ts.Close()

	const body = "Go is a general-purpose language designed with systems programming in mind."
	ts.URL = "http://10.26.27.128/v1/house/list"

	req, err := http.NewRequest("GET", "http://10.26.27.128/v1/house/list", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Host = "yuan.zu.ke.com"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}