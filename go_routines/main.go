package main

import (
	"fmt"
	"go_routines/data"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start:= time.Now()

	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go readBook(i,wg, m)
	}
	
	wg.Wait()

	duration := time.Since(start).Milliseconds()
	fmt.Printf("%dms\n",duration)
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex){

	data.FinishedBook(id,m)

	delay :=rand.Intn(800)

	time.Sleep(time.Millisecond*time.Duration(delay))

	wg.Done()
}
/*
func showGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Routine #%d with %d ms\n", id,delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}*/
