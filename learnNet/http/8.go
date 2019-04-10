package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.ke.com", nil)
	beginTime := time.Now().UnixNano()
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			endTime := time.Now().UnixNano()
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
			//微秒
			cost := int((endTime-beginTime)/1000)
			fmt.Printf("DNS END Info: %+v μs \n",cost)
		},
		DNSStart: func(dnsStartInfo httptrace.DNSStartInfo) {

			fmt.Printf("DNS Start Info: %+v\n", beginTime)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
