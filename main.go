package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

type NumbersList struct {
	Number []int
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
		log.Printf("Default port: %s", port)
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/", ShowNumbers)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Listening on port: %s", port)
	log.Fatal(server.ListenAndServe())

}

func ShowNumbers(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		numbers := rand.Perm(50)
		list := NumbersList{}

		for _, v := range numbers {

			if len(list.Number) < 6 && v != 0 {
				list.Number = append(list.Number, v)
			}

		}

		json, _ := json.Marshal(list)

		w.Write(json)

		time.Sleep(10 * time.Millisecond)

	}()

	wg.Wait()

}
