package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/src-d/go-github/github"
	"sync"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

var prPool = sync.Pool{
	New: func() interface{} { return new(github.PullRequestEvent)},
}

func handle(w http.ResponseWriter, r *http.Request) {

	data := prPool.Get().(*github.PullRequestEvent)
	defer prPool.Put(data)

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		log.Printf("Could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "pullrequest id: %d", *data.PullRequest.ID)
}
