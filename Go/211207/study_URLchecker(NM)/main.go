package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("요청 실패")

//gorutine 배운후 적용
type result struct {
	url    string
	status string
}

func main() {
	//var results = map[string]string{} //map 초기화하는법1
	var results = make(map[string]string) //map 초기화하는법2
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }
	//fmt.Println(results)
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	//gorutine 배운후 적용
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		allresult := <-c
		results[allresult.url] = allresult.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

//인자 타입에 <- 넣으면 send only ,
func hitURL(url string, c chan<- result) {
	//fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
	}
	c <- result{url: url, status: "OK"}
}
