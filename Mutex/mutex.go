package mutex

import (
	"fmt"
	"sync"
)

func read(s *string, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	fmt.Println(*s)
	m.Unlock()
	wg.Done()
}

func write(s *string, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*s = "s"
	fmt.Println(*s)
	m.Unlock()
	wg.Done()
}

func Use() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	s := ""
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go read(&s, wg, m)

		go write(&s, wg, m)

	}
	wg.Wait()
}
