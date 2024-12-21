package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const workersCount = 10

func generator(urls []string) <-chan string {

	urlChan := make(chan string)

	go func() {
		defer close(urlChan)
		for i:=0; i<len(urls); i++{
			urlChan <- urls[i]
		}
	}()

	return urlChan
}

func workerFunc(wg *sync.WaitGroup, urlChan <-chan string) {
	defer wg.Done()
	for {
		url, ok := <-urlChan

		if !ok {
			return
		}

		resp, o := http.Get(url)

		if(o != nil){
			fmt.Println("cant connect")
			return
		}
		fmt.Println(resp.Status)
	}
}

func main(){
	urlsPool := []string{
		"https://google.com",
		"https://stepik.org",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://github.com/",
		"https://spb.hh.ru/",
		"https://google.com",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://spb.hh.ru/",
		"https://google.com",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://stepik.org",
		"https://spb.hh.ru/",
		"https://google.com",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://github.com/",
		"https://spb.hh.ru/",
		"https://google.com",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://spb.hh.ru/",
		"https://google.com",
		"https://www.youtube.com",
		"https://my.itmo.ru",
		"https://stepik.org",
	}

	urlChan := generator(urlsPool)

	wg := &sync.WaitGroup{}
	
	timer := time.Now()

	for worker := 0; worker < workersCount; worker++ {
		wg.Add(1)
		go workerFunc(wg, urlChan)
	} 
	
	wg.Wait()
	fmt.Println(time.Since(timer))
	fmt.Println("Programm closed")
}