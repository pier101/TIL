package main

import "fmt"

/*
상수 : 변하지 않는 값.
상수로 사용될수 있는 타입
정수, 실수, 문자열, 룬(rune),불리언, 복소수

*/

//상수값에 코드를 부여
const APPLE int = 0
const BANANA int = 1
const GRAPE int = 2

const PI = 3.14
const fPI float64 = 3.14

func PrintFruit(fruit int) {
	if fruit == APPLE {
		fmt.Println("사과는 아침에 먹어야 제맛")
	} else if fruit == BANANA {

		fmt.Println("나는 바나나")
	} else if fruit == GRAPE {

		fmt.Println("나는 포도")
	}
}

func main() {

	// const ConstValue int = 10
	// //fmt.Println(&ConstValue)

	// const PI float64 = 3.141592

	PrintFruit(APPLE)
	PrintFruit(BANANA)
	PrintFruit(GRAPE)

	//iota : 1씩 증가하도록 정의할때는 이녀석을 쓰는게 편리.열거형

	const (
		RED    int = iota //0
		BLACK  int = iota //1
		YELLOW int = iota //2
	)

	const (
		//규칙이 동일하다면 아래와 같이 생략가능
		TIGER int = iota + 1 //0
		RABBIT
		DOG
	)

	fmt.Println(RED)
	fmt.Println(BLACK)
	fmt.Println(YELLOW)

	fmt.Println(TIGER)
	fmt.Println(RABBIT)
	fmt.Println(DOG)

	//타입이 없는 상수
	//var a int = PI * 100 //위에 타입이 없는 상수이기 때문에 오류X

	//var b int = fPI * 100 //타입에러

}
