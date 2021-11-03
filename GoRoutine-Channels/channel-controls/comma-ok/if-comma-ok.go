package commaok

import (
	"fmt"
	"sync"
)

//those example uses channel with safe way.
// controls channel if is close or got value 
// 


func commaOkControl() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 2
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()

}

//
func forLoopControl() {

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


func selectStatementControl() {

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