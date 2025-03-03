package main

import (
	"sync"
)

func main() {

}

type Pair struct {
	fstVal  int
	scndVal int
}

func generator(fstCh, scndCh <-chan int, n int) <-chan Pair {
	resultCh := make(chan Pair)

	go func() {
		defer close(resultCh)

		fstSl := make([]int, 0, n)
		scndSl := make([]int, 0, n)
		var flagDone bool
		var dataPtr int

		for !flagDone {
			select {
			case res, ok := <-fstCh:
				if !ok{
					return
				}

				fstSl = append(fstSl, res)

				if dataPtr < min(len(scndSl), len(fstSl)) {
					resultCh<- Pair{
						fstVal: fstSl[dataPtr],
						scndVal: scndSl[dataPtr],
					}
					dataPtr++
				}  

				if len(fstSl) == n && len(scndSl) == n {
					flagDone = true
					return
				}

			case res, ok := <-scndCh:
				if !ok{
					return
				}

				scndSl = append(scndSl, res)
				
				if dataPtr < min(len(scndSl), len(fstSl)) {
					resultCh<- Pair{
						fstVal: fstSl[dataPtr],
						scndVal: scndSl[dataPtr],
					}
					dataPtr++
				}  

				if len(fstSl) == n && len(scndSl) == n {
					flagDone = true
					return
				}
			}
		}
	}()

	return resultCh
}

func merge2Channels(fn func(int) int, in1, in2 <-chan int, out chan<- int, n int) {

	go func() {
		defer close(out)
		wg := &sync.WaitGroup{}

		valCh := generator(in1, in2, n)

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				pairVals := <-valCh


				funcOp := func(value int) <-chan int {
					res := make(chan int, 1)
					go func(){
						defer close(res)
						res <- fn(value)
					}()
					
					return res
				}
				
				out <- <-funcOp(pairVals.fstVal) + <-funcOp(pairVals.scndVal)
			}()
		}

		wg.Wait()
	}()

}