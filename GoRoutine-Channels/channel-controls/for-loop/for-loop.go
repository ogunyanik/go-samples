package forloop

import (
	"fmt"
	"sync"
)

func ChannelsControl() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		for i := range ch { //loop up to chan count
			fmt.Println(i)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //we must close channel rigt there,
		wg.Done()
	}(ch, wg)

	wg.Wait()

}
