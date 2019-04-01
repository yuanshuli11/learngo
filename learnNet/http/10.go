package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

//DumpRequestOut
func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."
	req, err := http.NewRequest("PUT", "http://10.26.27.128:6333", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)

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