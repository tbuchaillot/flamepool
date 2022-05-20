package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello2", hello2)

	fmt.Println("serving at :8989")
	http.ListenAndServe(":8989", nil)

}

type TestStruct struct {
	Exported    string
	notExported bool

	StrField struct {
		Test  string
		Other struct {
			OtherTest  bool
			nonEmbeded string
		}
	}

	JsonExported int `json:"name"`
}

func hello2(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)

	testStruct := TestStruct{
		Exported:    "val1",
		notExported: true,

		StrField: struct {
			Test  string
			Other struct {
				OtherTest  bool
				nonEmbeded string
			}
		}{Test: "test"},
		JsonExported: 5,
	}

	json.NewEncoder(rw).Encode(testStruct)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}
