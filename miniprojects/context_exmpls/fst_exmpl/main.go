package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func doSomeWork(ctx context.Context, resChan chan<- int, numb int) {
	sleepTime := time.Duration(rand.Intn(200)) * time.Millisecond
	timer := time.NewTimer(sleepTime)
	fmt.Println("The worker", numb, "will sleep for", sleepTime)
	select{
	case <-timer.C:
		resChan<- numb
		timer.Stop()
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int)

	for i:=0; i<10; i++ {
		go doSomeWork(ctx, result, i)
	}

	res := <-result
	finish()
	fmt.Println("The", res, "worker has finished")
}