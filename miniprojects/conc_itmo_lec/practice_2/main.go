package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Workerpool realisation

func Start(dataCh <-chan int, stopCh <-chan struct{}, transform func(int) int, workersCount int) <-chan int {
	resChannel := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()

			for{
				select {
				case val, ok := <-dataCh:
					if !ok {
						return
					}

					select{
					case <-stopCh:
						return
					case resChannel <- transform(val):
					}
					

				case <-stopCh:
					return
				}
			}
		}()
	}

	go func(){
		defer close(resChannel)
		wg.Wait()
	}()

	return resChannel
}

func summ(val int) int {
	return val + 10
}

func generator(count int, data chan<- int) {
	defer close(data)

	for i := 0; i < count; i++ {
		data <- rand.Intn(1000)
	}
}

func main() {
	doneCh := make(chan struct{})
	dataCh := make(chan int)

	ticker := time.NewTicker(time.Millisecond * 2000)
	defer ticker.Stop()

	go generator(10000, dataCh)
	resChan := Start(dataCh, doneCh, summ, 10)

	var flag bool

	for !flag {
		select{
		case val, ok := <-resChan:
			if !ok {
				flag = true
				break
			}
			fmt.Println(val)
		case <- ticker.C:
			close(doneCh)
			flag = true
		}
		
	}
	fmt.Println("Finish programm")
}