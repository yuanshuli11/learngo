package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var i int

type foo2Handler struct {
}

func (h *foo2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i++
	cookie := &http.Cookie{
		Name:    "XMEN",
		Value:   "STORM" + strconv.Itoa(i),
		Expires: time.Now().AddDate(1, 0, i),
	}

	http.SetCookie(w, cookie)

	fmt.Println(r.Cookies())
	fmt.Println("fooHandler ServeHTTP")
}
func main() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        new(foo2Handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
