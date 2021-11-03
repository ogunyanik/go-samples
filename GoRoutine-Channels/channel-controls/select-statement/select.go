package selectstatement

import (
	"fmt"
	"sync"
)

func ChannelsControl() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)
	ch2 := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(ch chan<- int, wg *sync.WaitGroup) {
			ch <- 3
			wg.Done()
		}(ch, wg)

		go func(ch2 chan<- int, wg *sync.WaitGroup) {
			ch2 <- 2
			wg.Done()
		}(ch2, wg)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(ch2 <-chan int, ch <-chan int, wg *sync.WaitGroup) {

			select {
			case b := <-ch2:
				fmt.Println("case 1 ")
				fmt.Println(b)

			case b, ok := <-ch:
				if ok {
					fmt.Println("case 2 ")
					fmt.Println(b)
				}

			}

		}(ch2, ch, wg)
		wg.Done()
	}

	wg.Wait()

}
