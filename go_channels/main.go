package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {

	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	FakeIDsChan := make(chan string)
	ClosedChan := make(chan int)

	wg.Add(3)
	go generateIDS(wg, IDsChan,ClosedChan)
	go generateFakeIDS(wg, FakeIDsChan,ClosedChan)
	go logIDS(wg, IDsChan, FakeIDsChan, ClosedChan)
	wg.Wait()

}

//send-only
func generateIDS(wg *sync.WaitGroup, idsChan chan<- string, closedChan chan<- int) {
	for i := 0; i < 100; i++ {

		id := uuid.New()
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())

	}

	close(idsChan)
	closedChan <- 1
	wg.Done()
}
func generateFakeIDS(wg *sync.WaitGroup, fakeIdsChan chan<- string, closedChan chan<- int) {
	for i := 0; i < 100; i++ {

		id := uuid.New()
		fakeIdsChan <- fmt.Sprintf("%d. %s", i+1, id.String())

	}

	close(fakeIdsChan)
	closedChan <- 1
	wg.Done()
}

//receive-only
func logIDS(wg *sync.WaitGroup, idsChan <-chan string, fakeIdsChan <-chan string, closedChannels chan int) {

	closedCounter := 0

	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("ID:", id)
			}
		case id, ok := <-fakeIdsChan:
			if ok {
				fmt.Println("FAKE ID:", id)
			}
		case count, ok := <-closedChannels:
			if ok{
				closedCounter+=count
			}
		}

		if closedCounter == 2 {
			close(closedChannels)
			break
		}
	}

	wg.Done()
}
