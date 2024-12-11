package main

import (
	"fmt"
	"sync"
)

func main(){
	type MyType func(string)(error)
	
	var myfunc = make(chan MyType)

	arrFunc := make([]MyType,0,3)
	arrFunc = append(arrFunc, myFunc1)
	arrFunc = append(arrFunc, myFunc2)
	arrFunc = append(arrFunc, myFunc3)

	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++{
		wg.Add(1)
		go func (i int, channel chan<- MyType) {
			defer wg.Done()
			fmt.Println("go anonym func")
			channel<-arrFunc[i]
		}(i, myfunc)
	}
	
	go func() {
		wg.Wait()
		close(myfunc) 
	}()
	
	fmt.Println("main func")
	for res := range myfunc{
		res("hello from main")
	}
}

func myFunc1(str string)(error){
	fmt.Printf("%s from fst func\n", str)
	return nil
}

func myFunc2(str string)(error){
	fmt.Printf("%s from scnd func\n", str)
	return nil
}

func myFunc3(str string)(error){
	fmt.Printf("%s from thd func\n", str)
	return nil
}

// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
//     resChan := make(chan int)
// 		go func(){
// 			defer close(resChan)
//         select{
//         case arg, ok := <-firstChan:
// 					if !ok {
// 						return
// 					} 
//           resChan<- arg * arg
// 				case arg, ok := <-secondChan:
// 					if !ok {
// 						return
// 					} 
//           resChan<- arg * 3

//         case <-stopChan:
//             return    
//         }
// 		}()
// 		return resChan
// }

// func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
//     var sum int
//     outChannel := make(chan int)
//     go func(){

//         defer close(outChannel)
//         defer func(){
// 					outChannel<- sum
// 				}()

//         for{
//             select{
//                 case val := <-arguments:
//                 sum += val
//                 case <-done:
//                 return 
//             }
//         }
//     }()
//     return outChannel
// }

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(){
		for i:=0; i<n; i++ {
			go func(){
				chan1 := make(chan int)
				chan2 := make(chan int)
				defer close(chan1)
				defer close(chan2)

				val1, ok1 := <-in1
				if ok1 {
						go func() {
								chan1 <- fn(val1)
						}()
				}
						
				val2, ok2 := <-in2
				if ok2 {
						go func() {
								chan2 <- fn(val2)
						}()
				}
				sum := <-chan1 + <-chan2
				out <-sum
			}()
		}
	}()
}





// func (chan struct{}){
	// 	fmt.Println("anonym func")
	// }(make(chan struct{}))