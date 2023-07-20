package routes

import (
	"dep-ex/internal/handlers/index"

	"github.com/gorilla/mux"
)

func Index() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index.Index).Methods("GET")
	//TODO: add GET route with path "/instagram"
	r.HandleFunc("/instagram", index.Instagram).Methods("POST")
	return r
}
