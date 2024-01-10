package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var hosts = map[string]interface{}{}

func main() {
    port := ":8000"

    for _, arg := range os.Environ() {
        env := strings.Split(arg, "=")
        conf := strings.Split(env[1], ":")
        if strings.HasPrefix(env[0], "SERVE") {
            hosts[conf[0]] = http.FileServer(http.Dir(conf[1]))
            log.Printf("Serving %s", conf[1])
        }
        if strings.HasPrefix(env[0], "PROXY") {
            hosts[conf[0]] = conf[1]
            log.Printf("Proxying to %s", conf[1])
        }
        if env[0] == "PORT" {
            port = fmt.Sprintf(":%s", env[1])
        }
    }

    proxy := &httputil.ReverseProxy{
        Rewrite: func(p *httputil.ProxyRequest) {
            switch v := hosts[p.In.Host].(type) {
            case string:
                url, _ := url.Parse(fmt.Sprintf("http://localhost:%s", v))
                p.SetURL(url)
                p.Out.Host = p.In.Host
            }
        },
    }

	log.SetOutput(io.Discard)
	// logger := log.New(os.Stdout, "", 0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch v := hosts[r.Host].(type) {
		case string:
            proxy.ServeHTTP(w, r)
		case http.Handler:
			v.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})

    key, keyOk   := os.LookupEnv("KEY");
    cert, certOk := os.LookupEnv("CERT")

    if keyOk && certOk {
        log.Printf("listening with tls in http://localhost%s", port)
		log.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
	} else {
        log.Printf("listening in http://localhost%s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}

