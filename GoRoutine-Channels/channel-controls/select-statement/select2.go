package selectstatement

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func Test() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	channel1 := make(chan int)
	channel2 := make(chan int)

	for i := 0; i < 5; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- int) {
			if b, ok := generateRandomBetween(1, 5); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, channel1)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- int) {
			if b, ok := generateRandomBetween(6, 10); ok {
				m.Lock()
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, channel2)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(channel1, channel2 <-chan int) {
			select {
			case b := <-channel1:
				fmt.Println("from channel1 ")
				fmt.Println(b)
			case b := <-channel2:
				fmt.Println("from channel2")
				fmt.Println(b)

			}
			wg.Done()
		}(channel1, channel2)

	}

	time.Sleep(150 * time.Millisecond)

	wg.Wait()
}

func generateRandomBetween(min, max int) (int, bool) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min, true
}
