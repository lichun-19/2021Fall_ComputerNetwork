package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File Not Found\n")
}

func AccessFile(prefix string) http.Handler {
	fs := http.Dir(".")
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		_, err := fs.Open(p)

		if os.IsNotExist(err) {
			NotFound(w, r)
		} else {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			fsh.ServeHTTP(w, r2)
		}
	})
}
func main() {
	fmt.Println("Launching server...")

	http.Handle("/", AccessFile("/"))
	http.ListenAndServeTLS(":12012", "server.cer","server.key", nil)
}
