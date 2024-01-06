package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {

	port := ":8000"
	url, err := url.Parse("http://localhost:3000")

	if err != nil {
		log.Fatal(err)
	}

	var key string
	var cert string

	proxyServer := httputil.NewSingleHostReverseProxy(url)
	fileServer := http.FileServer(http.Dir("./static"))

	hosts := map[string]http.Handler{
        "example.com": proxyServer,
        "static.example.com": fileServer,
    }

	log.SetOutput(io.Discard)
	// logger := log.New(os.Stdout, "", 0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target, ok := hosts[r.Host]

		if !ok {
			http.NotFound(w, r)
			return
		}

		target.ServeHTTP(w, r)
	})

	if key != "" && cert != "" {
		log.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
	} else {
		log.Fatal(http.ListenAndServe(port, nil))
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
