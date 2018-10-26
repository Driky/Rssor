package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"rssor/atom"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rss/trashscanlation", func(w http.ResponseWriter, r *http.Request) {
		encodeRss(w, r, atom.GetTrashscanlationsLastChapters())
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func encodeRss(w http.ResponseWriter, r *http.Request, rss atom.RSS) {
	w.Write([]byte(atom.Header))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	enc.Encode(rss)
}

func main() {
	handleRequests()
}
