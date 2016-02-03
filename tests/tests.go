package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	name, err := url.QueryUnescape(r.URL.RequestURI()[1:])
	if err != nil {
		fmt.Fprintf(w, "failed: %v", err)
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
