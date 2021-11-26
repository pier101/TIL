package main

import "fmt"

/*
	문자열 : 문자의 집합. Go에서는 UTF-8문자코드를 씀.
	유니코드 한 문자를 나타내는데 1~3바이트를 사용
	한글이나 한자 사용에 문제가 없다.
	아스키코드와 1:1매칭이 된다.
	UTF-8은 영문자, 숫자, 특수문자 1바이트 그외의 다른 문자는 2~3바이트를 사용
*/

func main() {
	var str string = "Hello\t'World'\n"  //더블쿼터로 묶으면 특수문자 발동
	var str1 string = `Hello\t'World'\n` //백틱으로 묶으면 특수문자 ㄴㄴ

	fmt.Println(str)
	fmt.Println(str1)

	//문자 하나를 표현할때는 rune타입을 사용한다.
	//go에서는 3바이트 정수타입을 제공하지 않는다.
	//rune은 int32타입의 별칭이다.(이름만 다를뿐 같음)

	var char rune = '한'
	fmt.Printf("%T\n", char) //변수타입(int32라고 나옴)
	fmt.Println(char)        //타입자체가 int32라서 숫자로..
	fmt.Printf("%c\n", char)

	var string1 string = "경일게임"
	var string2 string = "abcd"

	fmt.Println("string1의 길이", len(string1)) //한글은 3바이트
	fmt.Println("string2의 길이", len(string2)) //영문은 1바이트

	//string, rune 슬라이스 타입인 []rune타입은 상호 타입변환이 가능하다.
	//슬라이스는 길이가 변할 수 있는 배열(가변적)
	var string3 string = "Hello World"
	var runes = []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	fmt.Println(string3)
	fmt.Println(string(runes)) //타입 변환을 할 경우 rune배열(rune앞[])의 각 요소에는 문자열의 각 글자가 대입이 된다.
	//이걸 이용해 문자열의 글자수를 알 수 있음

	var string4 string = "Hello 월드"
	var runes1 = []rune(string4)                //[rune]타입으로 변환
	fmt.Println("string4의 길이 : ", len(string4)) //문자열의 바이트 길이가 반환
	fmt.Println("runes1의 길이 : ", len(runes1))   //각 글자들로 이루어진 배열로 반환

}
