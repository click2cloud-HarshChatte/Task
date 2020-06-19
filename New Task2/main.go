package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Number struct {
	A   []int
	Key int
}
type Result struct {
	Status bool `json:"status"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	n := new(Number)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		json.NewEncoder(w).Encode(err)
	}
	fmt.Println(n.A)
	fmt.Println(n.Key)
	result.Status = false
	for _, item := range n.A {
		if item == n.Key {
			result.Status = true
		}

	}
	//result.Status = false

	resultJson, err := json.Marshal(result)
	if err != nil {
		return
	}
	//w.WriteHeader(http.StatusOK)
	w.Write(resultJson)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Number/Search", Search).Methods("POST")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
