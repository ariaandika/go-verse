// package main
//
// import (
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// )
//
// func main() {
//   url, _ := url.Parse("http://localhost:8000")
//
//   reverseProxy := httputil.NewSingleHostReverseProxy(url)
//
//   log.Fatal(http.ListenAndServe(":3000", reverseProxy))
// }
