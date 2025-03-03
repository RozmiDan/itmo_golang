package main

import (
	"strconv"
	"sync"
)

//type job func(in, out chan interface{})

func ExecutePipeline(functions ...job) {
	for i := range functions {
		
	}
}

func SingleHash(data string, res chan string, quoteChan chan struct{}) <-chan string {
	resultChan := make(chan string)
	go func(){
		defer close(resultChan)
		resultCrc := make(chan string)
		go func(){
			resultCrc<- DataSignerCrc32(data)
			close(resultCrc)
		}()
		quoteChan<- struct{}{}
		resScnd := DataSignerMd5(data)
		<-quoteChan 
		resScnd = DataSignerCrc32(resScnd)
		
		resOfPipe := <-resultCrc + "~" + resScnd
		resultChan <- resOfPipe
	}()
	return resultChan
}

func MultiHash(prevStage <-chan string) <-chan string {
	returnChan := make(chan string)
	
	go func(){
		defer close(returnChan)
		data := <-prevStage
		wg := &sync.WaitGroup{}
		resultOfIter := make([]string, 6)
		for i := 0; i < 6; i++{
			wg.Add(1)
			go func(index int){
				defer wg.Done()
				resultOfIter[index] = DataSignerCrc32(strconv.Itoa(index) + data)
			}(i)
		}
		wg.Wait()
		var resultValue string 
		for i, _ := range resultOfIter{
			resultValue += resultOfIter[i]
		}
		returnChan<- resultValue
	}()
	
	return returnChan
}

func CombineResults(prevStage <-chan string) <-chan string {
	//

}