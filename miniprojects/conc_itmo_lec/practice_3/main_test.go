package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func generatorFunc() (fn func(int) int) {
	timeS := rand.Intn(4)
	log.Printf("Created fn with %v seconds lag\n", timeS)
	return func(val int) int{ 
		time.Sleep(time.Duration(timeS)*(time.Second))
		return val + 10 
	} 
}

func generatorValsForChan(count int, fstArr, scndArr *[]int) (<-chan int, <-chan int) {
	fstChan := make(chan int,1)
	scndChan := make(chan int,1)

	go func(){
		defer close(fstChan)
		defer close(scndChan)

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func(){
			defer wg.Done()
			for i := 0; i < count; i++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				randVal := rand.Intn(500)
				*fstArr = append(*fstArr, randVal)
				fstChan <- randVal
			}
		}()

		wg.Add(1)
		go func(){
			defer wg.Done()
			for i := 0; i < count; i++ {
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
				randVal := rand.Intn(500)
				*scndArr = append(*scndArr, randVal)
				scndChan <- randVal
			}
		}()

		wg.Wait()
	}()

	return fstChan, scndChan
}

func TestFunc(t *testing.T) {
	var count = 2
	fstSlice, scndSlice := make([]int, 0, count), make([]int, 0, count)
	fstChan, scndCh := generatorValsForChan(count, &fstSlice, &scndSlice)
	outCh := make(chan int)

	go merge2Channels(generatorFunc(), fstChan, scndCh, outCh, count)
	
	for i := range outCh {
		log.Println(i)
	}

	fmt.Println("Genetated values:")
	fmt.Println(len(fstSlice))
	for i:=0; i<len(fstSlice); i++{
		fmt.Print(fstSlice[i], scndSlice[i], fstSlice[i] + scndSlice[i] + 20, "\n")
	}


	// go func(){
	// 	for i := range fstChan{
	// 		log.Println(i)
	// 	}
	// }()

	// for i := range scndCh{
	// 	log.Println(i)
	// }

	fmt.Println("end")
	// fstFunc := generatorFunc()
	// scndFunc := generatorFunc()
	// res := fstFunc(10)
	// log.Println(res)
	// res = scndFunc(20)
	// log.Println(res)

}