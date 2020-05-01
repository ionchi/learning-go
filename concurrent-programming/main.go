package main

import (
	"fmt"
	"learn/go/concurrent"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]concurrent.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
	Concurrent programming with Go
	- concurrency: have multiple task to do, but working on one thing at a time with one worker
	- parallelism: execute multiple tasks simultaneously

	- thread: have own execution stack, fixed stack space (around 1MB), managed by OS
	- goroutines: have own execution stack, variable stack space (starts @2KB), managed by Go runtime [more efficient]
*/

func main() {
	// A WaitGroup waits for a collection of goroutines to finish
	// Note: as a pointer to pass this around the application, don't want to copy it
	wg := &sync.WaitGroup{}

	// A Mutex is a mutual exclusion lock: used to protect a portion of the code so that only one task or only
	// the owner of the mutex lock can access that code
	m := &sync.RWMutex{}

	// Channel are useful for the communication between goroutines
	// Channels are always created bi-directional
	// But passing channels around can be send-only (ch chan<- int) or receive-only (ch <-chan int)
	cacheCh := make(chan concurrent.Book)
	dbCh := make(chan concurrent.Book)

	for i := 0; i < 10; i++ {
		// generate a random id from 1 to 10 [we have 10 book in our "db"]
		id := rnd.Intn(10) + 1
		wg.Add(2)
		// send only channel
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- concurrent.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		// send only channel
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- concurrent.Book) {
			if b, ok := queryDatabase(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// Create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan concurrent.Book) {
			// similar to switch statement, but instead of checking for truthiness with select we gonna try
			// to either send or receive from a channel. If a channel is available to act (has a message)
			// that case will be taken. If both cases are available then will be randomly decided witch case
			// will fire.
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
				//default:
				// use default case for non-blocking select
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (concurrent.Book, bool) {
	// read mutex block
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (concurrent.Book, bool) {
	// to simulate db read
	time.Sleep(100 * time.Millisecond)
	for _, b := range concurrent.Books {
		if b.ID == id {
			// write and read mutex block
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return concurrent.Book{}, false
}
