package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

type User struct {
	Name string
}

func main() {

	http.HandleFunc("/", test)
	appengine.Main()

}

func test(w http.ResponseWriter, r *http.Request) {

	var user = User{
		Name: "Test",
	}

	jsonFile, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonFile)

}
