package main

import (
	// "embed"
	"encoding/json"
	"github.com/gorilla/mux"
	// "io/fs"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	// var embeddedFiles embed.FS
	// fsys, err := fs.Sub(embeddedFiles, "web/public")
	// if err != nil {
	// 	panic(err)
	// }

	// router.Handle("/", http.FileServer(http.FS(fsys)))
	// router.Handler("/", http.FileServer(http.Dir("./web/public")))

	router.HandleFunc("/Testing", testing).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func testing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res interface{}
	resJson := `{ "Testing": "Hello World!" }`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
