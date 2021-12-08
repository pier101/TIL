package main

import "fmt"

//인터페이스
type Stringer interface {
	String() string
}

//구조체
type Student struct {
	Age int
}

//스튜던트 타입의 String 메서드
func (s *Student) String() string {
	return fmt.Sprintf("나이 : %d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student) //*Student 타입으로 타입변환
	fmt.Printf("나이 : %d\n", s.Age)
}

func main() {
	//ㅁ 인터페이스 타입 변경
	//인터페이스 변수를 다른 구체환된 타입으로 변환할 수 있음
	//인터페이스를 본래 구체화된 타입으로 돌릴때
	// var a interface
	// t := a.(변경하려는 타입)

	//s := &Student{100} //*Student 타입 변수 s
	//PrintAge(s)        //변수 S를 인터페이스 인수로 PrintAge()부름

	// var stringer Stringer
	// stringer.(*Student) //오류: string메서드를 포함하고 있지 않기 때문

	/*
		ㅁ 정리
		1.인터페이스는 메서드 집합
		2.인터페이스에서 정의한 메서드 집합을 가진 모든 타입은 인터페이스로 쓰일 수있다.
		3.모든 타입이 빈 인터페이스 변숫값을 쓸 수 있다.
		4. 인터페이스 변환을 사용하면 인터페이스 변수를 구체화된 타입이나 다른 인터페이스로 변경 가능
	*/
}
