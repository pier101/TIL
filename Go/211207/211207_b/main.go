package main

import "fmt"

type account struct {
	balance    int
	firstName  string
	secondName string
}

//포인터 메서드
func (a *account) PionterMethod(amount int) {
	a.balance -= amount
}

//값 타입 메서드
func (a1 account) ValueMethod(amount int) {
	a1.balance -= amount
}

func (a2 account) RetrunValue(amount int) account {
	a2.balance -= amount
	return a2
}

func main() {
	/*
		리시버는 두 가지 타입으로 정의할 수 있다.
		1. 값 타입    2. 포인터 타입
	*/

	var A *account = &account{200, "kim", "dongwook"}
	A.PionterMethod(10)
	fmt.Println(A.balance) //190

	A.ValueMethod(20)
	fmt.Println(A.balance) //190

	var B account = A.RetrunValue(30)
	fmt.Println(B.balance) //160

	B.PionterMethod(10)
	fmt.Println(B.balance) //150
	/*
		- 리시버가 있으면 메서드 , 없으면 함수
		- 리시버는 메서드를 호출하는 주체. 메서드는 리시버를 통해서만 호출.
		- 메서드는 리시버에 속한 기능이다.
		- 포인터 메서드는 인스턴스 중심. > 값 변경 가능
		- 값 타입은 값이 모두 복사. 복사하는 양에 따라서 성능이슈가 발생할 수 있다.
	*/
}
