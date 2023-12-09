package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go countToTen(wg, 1)

	wg.Add(1)
	go countToTen(wg, 2)
	
	wg.Wait()

	result1 := countToTenWithChannel(1)
	result2 := countToTenWithChannel(2)

	result := fanIn(result1, result2)

	for {
		msg := <-result
		if msg == "" {
			break
		}

		fmt.Println(msg)
	}
}

func countToTen(wg *sync.WaitGroup, processID int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("process %d counting %d", processID, i))
	}
}

func countToTenWithChannel(processID int) <-chan string {
	result := make(chan string)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			result <- fmt.Sprintf("process %d counting %d and sending to channel", processID, i)
		}
		close(result)
	}()
	return result
}

func fanIn(result1, result2 <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case msg := <-result1:
				result <- msg
			case msg := <-result2:
				result <- msg
			}
		}
	}()
	return result
}