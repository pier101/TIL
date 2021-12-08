package main

import "fmt"

func cap1() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡1")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) // 복사가 아닌 참조로 캡쳐
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
func cap2() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡2")
	for i := 0; i < 3; i++ {
		v := i //v변수에 i값 복사
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
func main() {
	/*
		함수 리터럴 외부 변수를 내부로 가져오는 거를 캡쳐 라고 함
		캡쳐는 값 복사가 아닌 참조형태로 가져옴(변수의 주소를 포인터 값으로 복사한다.)
		※주의할 점

	*/
	cap1()
	cap2()
}
