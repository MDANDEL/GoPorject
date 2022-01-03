package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var urls map[int]string
var addr = "127.0.0.1:8000"

func main() {
	r := mux.NewRouter()
	//   entre accolades = extrait automatiquement        : = caractères autorisés
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	r.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Catégorie: %s\n", queryParams["category"][0])
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Catégorie: %v\n", vars["category"])
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}
