package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/motapratik/CallWordCountService/api"
)

func main() {
	//create a new router
	router := mux.NewRouter()
	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/TopTenWordCount", api.CallWordCountService).Methods("POST")
	//start and listen to requests
	http.ListenAndServe(":8090", router)
}
