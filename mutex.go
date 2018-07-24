package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	m := new(sync.Mutex)
	c := make(chan bool)

	for i := 0; i < 5; i += 1 {
		m.Lock()
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
			m.Unlock()

			c <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-c
	}
}
