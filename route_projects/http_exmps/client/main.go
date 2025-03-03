package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	url = "https://httpbin.org/post"
	contentType = "application/json"
	reqBody = `{"id": 999, "value": "content"}`
)

func main(){
	client := http.Client{Timeout: time.Second}
	req, _ := http.NewRequest(http.MethodGet, url, strings.NewReader(reqBody))

	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}