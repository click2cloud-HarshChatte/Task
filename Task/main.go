package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Result struct {
	Status string `json:"status"`
}

func Palindrome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	params := mux.Vars(r)
	inputString := params["str"]

	//var reverseString string = ""
	//var len(inputString)

	for i := 0; i < len(inputString)/2; i++ {

		if inputString[i] != inputString[len(inputString)-i-1] {
			result.Status = "Failed"
		} else {

			result.Status = "Success"
		}
	} //reverseString = reverseString + string(inputString[i])

	/*	if strings.ToLower(inputString) == strings.ToLower(reverseString) {
			result.Status = "Success"
		} else {
			result.Status = "Failed"
		}*/
	resultJson, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Palindrome/{str}", Palindrome)
	log.Fatal(http.ListenAndServe(":8020", router))
}
