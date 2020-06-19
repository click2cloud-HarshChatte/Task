package main

 import (
	"encoding/json"

	"github.com/gorilla/mux"

	"fmt"

	"log"
	"net/http"
)

 type Person struct
 {
	 Firstname string  `json:"firstname"`
	 Lastname string     `json:"lastname"`
 }

 type Fullname struct
 {
	 Fullname string    `json:"fullname"`
 }


 func Concatinate( w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
p:= new(Person)
decoder:= json.NewDecoder(r.Body)
if err:=decoder.Decode(&p); err!=nil {
	json.NewEncoder(w).Encode(err)
}

json.NewEncoder(w).Encode(Fullname{
	Fullname:fmt.Sprintf("%s%s", p.Firstname, p.Lastname)})

//  result:=Result{}
//  result.Fullname=arr.Firstname+arr.Lastname
//  json.NewEncoder(w).Encode(result.Fullname)

 //Body,err := ioutil.ReadAll(r.Body)
 //if err !=nil {
//	 return 
 //}
//  keyVal := make(map[string]string)
//  json.Unmarshal(Body, &keyVal)
// firstname := p["firstname"]
//  lastname := p["lastname"]
//  err := json.NewDecoder(r.Body).Decode(&p)
 //if err !=nil 
 //{
	// http.Error(w, err.Error(), http.StatusBadRequest)
 //}
 }


func main() {
	r:= mux.NewRouter()
//	r.HandleFunc("/person/Conca",Concatinate).Methods("GET")
	r.HandleFunc("/person/Concatinate",Concatinate).Methods("POST")
	err:=http.ListenAndServe(":8080",r)
	if err!=nil {
		log.Fatal("ListenAndServe:",err)
}
}

