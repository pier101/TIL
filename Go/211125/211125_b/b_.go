package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age   int
	Score float64
}

type Test struct {
	A int8 //1byte
	B int  //8byte
	C int8 //1byte
	D int  //88yte
	E int8 //1byte
}
type Test1 struct {
	A int8 //1byte
	C int8 //1byte
	E int8 //1byte
	B int  //8byte
	D int  //88yte
}

func PrintStudent(st Student) {
	fmt.Printf("나이 : %d, 점수 : %.2f", st.Age, st.Score)
}
func main() {
	var s = Student{15, 20.5}
	// var a int = 10 //인자로 구조체 받기때매 a 못받음
	PrintStudent(s)
	s1 := s          //s 구조체 모든 멤버(필드)들이 s1복사
	PrintStudent(s1) //함수 호출할떄도 구조체가 복사

	fmt.Println(unsafe.Sizeof(s)) //12byte인데 16byte 나옴 >>메모리 정렬떄문임

	var t = Test{1, 2, 3, 4, 5}
	var t1 = Test1{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(t))  //12byte인데 16byte 나옴 >>메모리 정렬떄문임
	fmt.Println(unsafe.Sizeof(t1)) //12byte인데 16byte 나옴 >>메모리 정렬떄문임

}
