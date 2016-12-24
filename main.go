package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var url = flag.String("url", "http://www.google.com", "URL to proxy status code from")
	var serveAddress = flag.String("serveaddress", ":3000", "address to listen and serve upon")
	log.Printf("Reverse proxying to %s", *url)
	log.Printf("Listening on %s", *serveAddress)
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Received request URI %s", req.RequestURI)
		newUri := fmt.Sprintf("%s%s", *url, req.RequestURI)
		res, err := http.Get(newUri)
		if err != nil {
			log.Println(err)
			w.WriteHeader(404)
		} else {
			w.WriteHeader(res.StatusCode)
		}
	})

	http.ListenAndServe(*serveAddress, mux)
}
