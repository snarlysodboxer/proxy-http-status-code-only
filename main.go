package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var checkURL = flag.String("check-url", "http://www.google.com", "URL to proxy status code from")
	var listenURL = flag.String("listen-url", "http://localhost:3000/status-code", "address to listen and serve upon")
	flag.Parse()

	parsedURL, err := url.Parse(*listenURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reverse proxying status code from %s", *checkURL)
	log.Printf("Listening on %s at %s", parsedURL.Host, parsedURL.Path)

	mux := http.NewServeMux()
	mux.HandleFunc(parsedURL.Path, func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get(*checkURL)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.WriteHeader(res.StatusCode)
		}
	})

	http.ListenAndServe(parsedURL.Host, mux)
}
