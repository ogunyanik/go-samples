package rwmutex

import (
	"fmt"
	"sync"
)

func read(s *string, wg *sync.WaitGroup, m *sync.RWMutex) {
	m.RLock()
	fmt.Println(*s)
	m.RUnlock()
	wg.Done()
}

func write(s *string, wg *sync.WaitGroup, m *sync.RWMutex) {
	m.Lock()
	*s = "s"
	fmt.Println(*s)
	m.Unlock()
	wg.Done()
}

func Use() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	s := ""
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go read(&s, wg, m)

		go write(&s, wg, m)

	}
	wg.Wait()
}

//
