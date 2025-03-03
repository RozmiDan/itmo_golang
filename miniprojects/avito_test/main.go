package main

type user struct {
	balance int64
}

///////////////////Fst task
// func main(){
// 	users := []user {
// 		{balance: 1000},
// 		{balance: 2000},
// 	}

// 	for _, val := range users {
// 		val.balance += 1000
// 	}

// 	for indx, _ := range users {
// 		users[indx].balance += 1000
// 	}

// 	fmt.Println(users)
// }

///////////////////Scnd task
// func main(){

// 	a := [2]int{0,0}

// 	if a == nil { // Not compiled
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

///////////////////3 task
// func main() {
// 	m := map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 	}
// 	for _, v := range m {
// 		fmt.Println(v) // Cant predict answer, bcz map is not sorted
// 	}
// }

///////////////////4 task
// func main(){
// 	b := new(bool)
// 	modify(b)
// 	fmt.Println(b) // will show the nil-pointer
// }

// func modify(b *bool) {
// 	b = nil
// }

///////////////////5 task
// func main(){
// 	var wg sync.WaitGroup
// 	counter := 0
// 	for i:=0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			counter++ // data race
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(counter)
// }

///////////////////6 task
// func main(){
// 	ch := make(chan int)
// 	close(ch)

// 	val := <-ch
// 	fmt.Println(val) // will print 0!!!!!!!!!!!
// }

///////////////////7 task
// func main(){
// 	v := 1
// 	fmt.Println(v, " ")
// 	f(&v)
// 	fmt.Println(v, " ") // 1 3
// }

// func f(v *int) {
// 	*v = *v * 3
// }

///////////////////8 task
// func main(){
// 	inc := func(){
// 		v++          // Undeclared
// 	}

// 	v := 1

// 	fmt.Println(v, " ")
// 	inc()
// 	fmt.Println(v, " ")
// }

///////////////////9 task
// func main(){
// 	ch := make(chan int)
// 	close(ch)
// 	close(ch) // panic()
// }

/////////////////10 task
// func main() {
// 	var ch chan int = nil
// 	close(ch) // panic()
// }

// /////////////11 task
// func main() {
// 	var m = make([]int, 500)
// 	wg := new(sync.WaitGroup)
// 	for i := 0; i < 500; i++ {
// 		wg.Add(1)
// 		go func(i int){
// 			m[i] = i
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Print("f done") // all ok
// }

/////////////////10 task
// func main() {
// 	var ch chan int = nil
// 	close(ch) // panic()
// }