package main

import (
	"fmt"
	"os"
)

//인수 타입 앞에 ... 해당타입 인수를 여러개 받는 가변인수
//가변 인수는 함수 내부에서 해당타입의 슬라이스로 처리
func sum(nums ...int) int { //가변 인수를 받는 함수
	sum := 0
	fmt.Printf("타입 : %T\n", nums)

	for _, v := range nums {
		sum += v
	}
	return sum
}

// func Print(args ...interface{}) string { //모든 타입을 받는 가변인수라는 뜻
// 	for _, arg := range args { //1.모든 인수를 돌면서
// 		switch b := arg.(type) { //2.인수 타입에 따라서
// 		case bool:
// 			val := arg.(bool) //3.변환
// 			//출력...
// 		case int:
// 			val := arg.(int) //3.변환
// 		}
// 	}
// }

//함수 리터럴
type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b //함수 리터럴을 이용해서 더하기 함수 정의하고 리턴
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	//가변인수함수 : 함수인수갯수가 고정적이지 않은 함수
	// ...쓰면 가변인수 처리가능
	fmt.Println()
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4)

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2))
	fmt.Println(sum())

	fmt.Println(1, "ㅇㅇ", 1.55)
	//위처럼 작동할 수 있는 이유? 내부적으로 가변인수를 빈 인터페이스로 받아주고 있음

	//////////////////////////////////////////////////////////////////////
	// 함수종료 전에 실행할 코드 있으면 defer사용
	// defer 명령

	f, err := os.Create("test.txt") //파일 생성
	if err != nil {
		fmt.Println("실패")
		return
	}

	//defer 호출 순서는 역순으로 진행된다.
	defer fmt.Println("호출")    //지연
	defer f.Close()            //지연
	defer fmt.Println("파일 닫음") //지연

	fmt.Println("작성")
	fmt.Fprintln(f, "hello world") //Fprintln : 파일핸들에 text를 쓰는 함수

	////////////////////////////////////////////////////////////////////

	//ㅁ 함수 리터럴 : 이름이 없는 함수(익명함수,람다)
	//                함수명을 쓰지 않고 함수 타입 변숫값으로 대입되는 함숫값
	//				  함수명이 없기 때문에 직업호출 안됨. 함수 타입변수로만 호출가능
	fn := getOperator("*")
	res := fn(3, 4) //함수 타입 변수를 사용해서 함수 호출
	fmt.Println(res)

	a := 0 //함수 외부 변수
	fu := func() {
		//함수 외부변수가 힘스 리터럴 내부상태로 가져옴 (캡쳐 라고 함)
		//함수외부변수를 복사가 아닌 참조 형태로 가져옴
		//※주의할 점

		a += 10
	}
	a++
	fu() //함수 타입 변수로 호출
	fmt.Println(a)

}
