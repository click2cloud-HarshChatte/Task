package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Status string `json:"status"`
	Value  int    `json: "value"`
}

func Numbers_Fib(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	params := mux.Vars(r)["number"]

	number, err := strconv.ParseInt(params, 10, 16)
	if err != nil {
		result.Status = "Failed"
		result.Value = 0
		json.NewEncoder(w).Encode(result)
		return
	} else {
		t1 := 0
		t2 := 1
		nextTerm := 0

		for i := 1; i <= int(number); i++ {
			if i == 1 {
				//fmt.Print(" ", t1)
				continue
			}
			if i == 2 {
				//fmt.Print(" ", t2)
				continue
			}
			nextTerm = t1 + t2
			t1 = t2
			t2 = nextTerm

		}

		fmt.Print(nextTerm)

		result.Status = "Success"
		result.Value = nextTerm
		json.NewEncoder(w).Encode(result)
	}
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/Numbers_Fib/{number}", Numbers_Fib)
	log.Fatal(http.ListenAndServe(":8020", router))
}
