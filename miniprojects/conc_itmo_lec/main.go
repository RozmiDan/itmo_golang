package main

import (
	"fmt"
	"sync"
)

type myStructWithMutex struct {
	mut *sync.Mutex
	str string
	counter int
}

func main(){
	a := myStructWithMutex{mut: &sync.Mutex{}, str : "", counter : 0}
	wg := &sync.WaitGroup{}

	for range 10{
		wg.Add(1)
		go func(){
			defer wg.Done()
			a.mut.Lock()
			defer a.mut.Unlock()
			a.counter++
			a.str += "1"
		}()
	}

	wg.Wait()
	fmt.Println(a.counter, a.str)

}