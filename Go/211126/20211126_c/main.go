package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type StringHeader struct{
// 	Data uintptr  //문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
// 	Len int		//int타입의 문자열의 길이
// }

func main() {
	str := "안녕 난 홍길동"
	str1 := str
	fmt.Println(str)
	fmt.Println()
	fmt.Println(str1)
	/*
		위 에제 원리
		1.str 문자열을 하나 복사해서 str1이 가리키게한다.
		2.str의 Data와 Len값만 str1에 복사
		>>정답 : 2번 예제 따름
	*/

	// (&)포인터주소값 변경 막고있음
	// 변수주소값으로 형변환하기 위해
	// (*reflect.StringHeader)로 선변환후 위로 변환(강제변환)
	// 문자열 데이터 자체가 복사되는 개념이 아니라
	// 구조체 안의 Data값 주소값(16byte)만 복사된다
	// >>성능상 문제 걱정 덜어도 됨
	//				 (       타입 변환     )
	// StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	// StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	// fmt.Println(StringHeader1) //주소값
	// fmt.Println(StringHeader2)

	var str2 string = "Hello World"
	// str2 = "경일게임아카데미"
	// fmt.Println(str2)
	//str2[3] = 'v' //문자 일부 변경 불가
	// []byte // ->부호가 없는 정수타입의 가변길이 배열(1byte)
	var slice []byte = []byte(str2)
	slice[2] = 't'
	fmt.Println(str2)
	fmt.Printf("%s\n", slice)
	StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str2 : \t%x\n", StringHeader1.Data)
	fmt.Printf("slice : \t%x\n", StringHeader2.Data)

}
