package channeltypes

import (
	"fmt"
	"sync"
)

//Note:Which direction the arrow pointing gets the value.
//Ex: If Arrow pointing to the channel, this means channel gets the message (Send Only Channel)

func ChannelTypess() {

	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(2)

	go func(ch <-chan string, wg *sync.WaitGroup) { //This channel only recive the message (Receive Only channel)
		fmt.Println(<-ch)
		wg.Done()
	}(ch, &wg)

	go func(ch chan<- string, wg *sync.WaitGroup) { //This channel only sets the message (Sendy Only channel)
		ch <- "exaample"
		wg.Done()
	}(ch, &wg)

	wg.Wait()
}
