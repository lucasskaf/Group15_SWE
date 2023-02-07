package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/create", createHandler).Methods("POST")
	router.HandleFunc("/read", readHandler).Methods("GET")
	router.HandleFunc("/update", updateHandler).Methods("PUT")
	router.HandleFunc("/delete", deleteHandler).Methods("POST")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalln("Error: Problem with the server", err)
	}
}

type Movie struct {
	Name   string `json:"name"`
	Rating int8   `json:"rating"`
}

var MovieDatabase = make([]Movie, 0)

func createHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var sampleMovie Movie

	err := json.NewDecoder(r.Body).Decode(&sampleMovie)

	if err != nil {
		log.Fatalln("Error: Not able to DECODE the request body into the struct")
	}

	MovieDatabase = append(MovieDatabase, sampleMovie)

	err = json.NewEncoder(w).Encode(&sampleMovie)

	if err != nil {
		log.Fatalln("Error: Not able to ENCODE the request body into the struct")
	}
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"]

	for _, structs := range MovieDatabase {
		if structs.Name == params {
			err := json.NewEncoder(w).Encode(&structs)

			if err != nil {
				log.Fatalln("Error: Encoding the initialized struct not successful")
			}
		}
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sampleMovie Movie

	err := json.NewDecoder(r.Body).Decode(&sampleMovie)

	if err != nil {
		log.Fatalln("Error: Not able to DECODE the request body into the struct")
	}

	for index, structs := range MovieDatabase {
		if structs.Name == sampleMovie.Name {
			MovieDatabase = append(MovieDatabase[:index], MovieDatabase[index+1:]...)
		}
	}

	MovieDatabase = append(MovieDatabase, sampleMovie)

	err = json.NewEncoder(w).Encode(&sampleMovie)

	if err != nil {
		log.Fatalln("Error: Encoding the initialized struct not successful")
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"]
	indexChoice := 0

	for index, structs := range MovieDatabase {
		if structs.Name == params {
			indexChoice = index
		}
	}

	MovieDatabase = append(MovieDatabase[:indexChoice], MovieDatabase[indexChoice+1:]...)
}
