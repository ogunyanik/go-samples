package concurrency

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func feth(url string, wg *sync.WaitGroup) (string, error) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error" + url)
		return "", err
	}

	wg.Done()
	return resp.Status, nil

}

func homePage(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup
	for i := range urls {
		wg.Add(1)
		go feth(urls[i], &wg)
	}

	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

func Handles() {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
