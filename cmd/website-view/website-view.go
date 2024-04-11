package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
	lookup "website-graph/cmd/lookupBackend"
)

var maxDepth int
var ignoreList []string

func getImage(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("remoteUrl") {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Remote URL Missing")
		return
	}

	urlToCheck := r.URL.Query().Get("remoteUrl")
	uri, _ := url.Parse(urlToCheck)

	initialUrlCookie, err := r.Cookie("initialUrl")

	if err != nil {

		initialUrlCookie = &http.Cookie{
			Name:     "initialUrl",
			Value:    uri.Host,
			SameSite: http.SameSiteDefaultMode,
			Expires:  time.Now().Add(time.Minute * 5)}
		http.SetCookie(w, initialUrlCookie)
	}

	if initialUrlCookie.Value != uri.Host {
		w.WriteHeader(http.StatusPartialContent)

	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Add("Content-Type", "image/svg+xml")
	io.WriteString(w, lookup.DoSearch(urlToCheck, initialUrlCookie.Value))

}

func main() {

	lookup.StartUp()
	//lookup.DoSearch()

	mux := http.NewServeMux()
	mux.HandleFunc("/", getImage)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
