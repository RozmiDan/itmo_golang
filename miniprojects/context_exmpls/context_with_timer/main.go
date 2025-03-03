package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doSomeWork(ctx context.Context, resChan chan<- int, numb int, wg *sync.WaitGroup) {
	sleepTime := time.Duration(rand.Intn(200)) * time.Millisecond
	timer := time.NewTimer(sleepTime)
	fmt.Println("The worker", numb, "will sleep for", sleepTime)
	for{
		select{
		case <-timer.C:
			resChan<- numb
			timer.Stop()
		case <-ctx.Done():
			defer wg.Done()
			return
		}
	}
	
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	result := make(chan int, 10)
	wg := &sync.WaitGroup{}
	defer cancel()

	for i:=0; i<10; i++ {
		wg.Add(1)
		go doSomeWork(ctx, result, i, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result{
		fmt.Println("The", res, "worker has finished")
	}
}