package main

import (
	"flag"
	"github.com/didip/tollbooth"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var checkURL = flag.String("check-url", "http://www.google.com", "URL to proxy status code from")
	var listenURL = flag.String("listen-url", "http://localhost:3000/status-code", "address to listen and serve upon")
	flag.Parse()

	parsedURL, err := url.Parse(*listenURL)
	if err != nil {
		log.Fatal(err)
	}
	path := "/"
	if parsedURL.Path != "" {
		path = parsedURL.Path
	}
	log.Printf("Reverse proxying status code from %s", *checkURL)
	log.Printf("Listening on %s at %s", parsedURL.Host, path)

	mux := http.NewServeMux()
	mux.Handle(path,
		tollbooth.LimitFuncHandler(
			tollbooth.NewLimiter(1, time.Second),
			func(w http.ResponseWriter, req *http.Request) {
				if req.URL.Path != path {
					http.NotFound(w, req)
					return
				}
				res, err := http.Get(*checkURL)
				if err != nil {
					log.Println(err)
					w.WriteHeader(500)
				} else {
					w.WriteHeader(res.StatusCode)
				}
			},
		),
	)

	http.ListenAndServe(parsedURL.Host, mux)
}
