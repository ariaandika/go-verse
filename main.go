package main

import (
	// "fmt"
	// "io"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
  // args := os.Args[1:]
  port := fmt.Sprintf(":%s", getEnv("PORT", "8000"))
  target := getEnv("TARGET", "3000")

  var key string
  var cert string

  for i, arg := range os.Args {
    if arg == "--key" {
      key = os.Args[i + 1]
    } else if arg == "--cert" {
      cert = os.Args[i + 1]
    }
  }

  // originServerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  //
  //   io.Copy(w, r.Body)
  //   // io.Copy(os.Stdout, r.Body)
  //
  //   fmt.Fprintf(w, "Hello, World!")
  // })

  url, err := url.Parse(fmt.Sprintf("http://localhost:%s", target))

  if err != nil {
    fmt.Errorf("Oof url")
    return
  }

  server := httputil.NewSingleHostReverseProxy(url)


  if key != "" && cert != "" {
    log.Fatal(http.ListenAndServeTLS(port, cert, key, server))
  } else {
    log.Fatal(http.ListenAndServe(port, server))
  }
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
