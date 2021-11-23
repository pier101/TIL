package main

import (
	"fmt"
	"math"
)

const epsilon = 0.000001 //매우작음

func equal(a, b float64) bool {

	return math.Nextafter(a, b) == b
	// if a > b {
	// 	if a-b <= epsilon {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// } else {
	// 	if b-a <= epsilon {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
}
func main() {

	var a int = 10
	var b int = 20
	var f float32 = 123231255.546
	// fmt.Println(1)

	fmt.Printf("a : %d\\ b : %d\t f: %f\n", a, b, f)

	/*
		% : 서식문자임
		%T : 데이터 타입 출력
		%t : 불리언을 true/false
		%d : 10진수 정수값으로 출력
		%b : 2진수 출력
		%c : 유니코드 문자를 출력
		%x : 16진수로 출력(소)
		%X : 16진수로 출력(대)
		%f : 실수 추력
		%e : 지수형태
		%s : 문자열을 출력
		%p : 메모리 주소값을 출력

		\n:개행
		\t:
		\\:역슬래쉬 자체 출력

	*/

	var str string = "H\nello\t\tWor\\l\"d\""
	fmt.Println(str)

	//Scan()
	var input int32
	var input1 int32
	//입력받는다. 인자로는 해당 변수의 메모리 주소
	fmt.Scan(&input, &input1)

	//비트연산 AND, OR, XOR, 비트클리어
	// & , | , ^, &^(특정 비트를 0으로)
	//시프트연산 <<(왼쪽) ,>>(오른쪽)

	//00000010 <<1

	/*AND 1일때만 1
	A 	B	A&B
	-------------
	0	0	0
	1	0	0
	0	1	0
	1	1	1
	*/
	//10 & 34 =2
	/*
		0000 1010		10
		0010 0010		34
		0000 0010		2
	*/

	/*OR : 하나라도 1이면 1
	0	0	0
	1	0	1
	0	1	1
	1	1	1


	0000 1010		10
	0010 0010		34
	0010 1010		42
	*/
	/*
		A^B
		XOR(A와 B가 다르면 1)
			0	0	0
			1	0	1
			0	1	1
			1	1	0



		0000 1010		10
		0010 0010		34
		0010 1000		40
	*/

	/*
		==, !-, <,>,<=,>=
	*/

	var d float64 = 0.1
	var e float64 = 0.2
	var float float64 = 0.3

	// fmt.Printf("%f+%f==%f : %v\n", fa, fb, fc, fa+fb == fc)
	// fmt.Println(fa + fb)

	fmt.Printf("%0.18f + %0.18f = %0.18f \n", d, e, d+e)
	fmt.Printf("%0.18f ==%0.18f : %v\n", float, d+e, equal(d+e, float))

	// &&,||,!,++,--

	// []
	// .
	// &
	// *
	// ...
	// :

	var test int8 = 30
	test <<= 2
	a += 8
	println(test)

}
