package main

import "fmt"

//go 에서는 클래스와 그 안의 멤버함수가 없기때문에
//구조체 밖에서 [메소드]라는 것으로 비슷하게 구현
//리시버 : 어느 구조체에 속하는지 표시하기 위해
//		  메서드가 속하는 타입을 알려주는 녀석

type account struct {
	b int
}

//일반함수 방식
func Function(a *account, amount int) {
	a.b -= amount
}

//메서드 방식
func (a *account) Method(amount int) {
	a.b -= amount
}

//사용자 정의 타입
type myNum int

func (a myNum) add(b int) int {
	return int(a) + b
}

func main() {
	//메서드 선언
	// func (r Rect) cal()int{
	// 	retrun r.width* r.height
	// }

	test := &account{200} //b가 200
	Function(test, 30)

	test.Method(20)
	fmt.Printf("%d \n", test.b)

	var a myNum = 10
	fmt.Println(a.add(20))
	var b int = 20
	fmt.Println(myNum(b).add(100))
}
