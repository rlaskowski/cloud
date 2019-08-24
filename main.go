package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"google.golang.org/appengine"
)

type Numbers struct {
	Number []int `json:"number"`
}

func main() {

	http.HandleFunc("/select", select)
	
	appengine.Main()
        
}

func select(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		numbers := rand.Perm(50)
		list := Numbers{}

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
