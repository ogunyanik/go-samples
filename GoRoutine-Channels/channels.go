// Sends request to url async way with "goroutine"
// Collects the result by "channel"
// Uses waitGroup for wait the "goroutines"

package channels

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func feth(url string, wg *sync.WaitGroup, status chan string) {

	resp, err := http.Get(url) //Send Get request to urls

	if err != nil {
		fmt.Println("Error" + url)
	}

	wg.Done() // after request says to the waitGroup I'am done

	fmt.Println(resp.Status + url)
	status <- url //channel sharing between go routines, and we adds value to channel

}

func run() {

	var wg sync.WaitGroup // wait group created

	c := make(chan string) //channel crated
	c2 := make(chan string)
	var statusCodeArray []string

	for i := range urls {
		wg.Add(2)
		go feth(urls[i], &wg, c)

		responseFromChannel := <-c
		go feth(responseFromChannel, &wg, c2) //send request to url, which is gets from channel1(c)

		statusCodeArray = append(statusCodeArray, responseFromChannel) // <- this operator gets value from channel // <- this operator gets value from channel

	}

	defer close(c) // runs after all line, closes the channel
	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Println(statusCodeArray)
}

func Run() {
	run()
}
