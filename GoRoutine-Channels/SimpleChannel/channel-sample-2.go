//Returns random int for 10 times
// Makes method calls async way
//Waitgroup used for wait all threads

package simplechannel

import (
	"fmt"
	"math/rand"
	"sync"
)

func CalculateValue(values chan int, wg *sync.WaitGroup) {
	wg.Done()
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value

}

func Calculate2() {
	fmt.Println("Go Channel Tutorial")

	var wg sync.WaitGroup

	values := make(chan int)
	defer close(values)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go CalculateValue(values, &wg)

	}

	wg.Wait()
	// defer close(values)
	value := <-values
	fmt.Println(value)
}
