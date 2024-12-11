package main

import (
	"fmt"
	"sync"
)


func main(){
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(){
		for i := 0; i < 10; i++{
			wg.Add(1)
			go work(i, wg)
		}
		wg.Done()
	}()
	wg.Wait()
}

func work(number int, wg *sync.WaitGroup){
	fmt.Println("Done", number)
	defer wg.Done()
}


// func main(){
// 	wg := new(sync.WaitGroup)
// 	var val int = 0
// 	var mu sync.Mutex
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1) // Верно
// 		go func(){
// 			defer wg.Done()
// 			//wg.Add(1) // НЕВЕРНО!!!
// 			mu.Lock()
// 			work(&val)
// 			mu.Unlock()
// 		}()
// 	}
// 	wg.Wait()
// }

// func work(number *int){
// 	fmt.Println("Done", *number)
// 	*number += 1
// }


// func main(){
// 	// synchronise gorutines example
// 	<-funcWithGorut()
// }

// func funcWithGorut() (<-chan struct{}) {
// 	chanRes := make(chan struct{})
// 	go func(){
// 		fmt.Println("Hello")
// 		chanRes <- struct{}{}
// 		//close(chanRes)
// 	}()
// 	time.Sleep(time.Second)
// 	return chanRes
// }


// func main(){
// 	// deadlock example
// 	chan_fst := make(chan int, 3)
// 	chan_scnd := make(chan int, 4)

// 	go generate(chan_fst, 10)
// 	go replecator(chan_fst, chan_scnd)

// 	for it := range chan_scnd{
// 		fmt.Println(it)
// 	}

// }

// func generate(c chan<- int, size int) {
// 	for i := 0; i < size; i++ {
// 		c <- i;
// 	} 
// 	close(c)
// }

// func replecator(ch_r <-chan int, ch_w chan<- int) {
// 	for it := range ch_r{
// 		ch_w <- it * it
// 	}
// 	close(ch_w);
// }

// func task2(channel chan string, str string) {
//   tmp_str := []rune(str)
// 	tmp_str = append(tmp_str, ' ') 
// 	for i := 0; i < 5; i++{
// 		channel <- string(tmp_str)
// 	} 
// }

// func removeDuplicates(inputStream <-chan string, outputStream chan<- string){
//     var prev_val string
//     for cur_val := range inputStream{
//         if prev_val == cur_val{
//             continue
//         } else {
//             outputStream <- cur_val
//         }
//         prev_val = cur_val
//     }
//     close(outputStream)
// }
