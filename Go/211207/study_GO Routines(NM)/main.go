package main

import (
	"fmt"
	"time"
)

//GoRoutines : 여러 함수 동시에
//Channel : goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법
func main() {
	//go 명시해줌으로써 동시 진행되게 해줌
	//GoRoutines은 메인함수가 실행하는 동안만 유효하기 때문에
	//모든 곳에 go 명시해주면 작업 일어나지 않음
	// go sexyCount("wook")
	// go sexyCount("nico")
	// sexyCount("soon")

	c := make(chan string) //1. channel 만듦
	people := [2]string{"nico", "wook"}
	for _, person := range people {
		go isSexy(person, c) //2. channel 인자로 넘김
	}
	// time.Sleep(time.Second * 10)
	fmt.Println("waiting for message...")
	// result := <-c
	// fmt.Println(result)
	// fmt.Println(<-c)

	//people늘어나면 계속 <-c새로 써줘야될까? >>loop로 해결
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

// func sexyCount(person string, c chan bool) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		//time : go 패키지 중 하나
// 		time.Sleep(time.Second)
// 	}
// }

func isSexy(person string, c chan string) { //3. 인자로 받은 channel
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + " : OK" //4.channel에 true값 넘겨줌으로써 파이프라인 형성
}
