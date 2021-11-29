package main

import "fmt"

func main() {

	/*	ㅁ 문자열 순회
		1.인덱스를 사용해 직접 접근 하는 방법
	*/
	var str string = "Hello 경일"

	//for i := 0; i < len(str); i++ { //문자열 크기만큼 돌린다.
	// fmt.Printf("타입 : %T , 값 : %d , 문자 : %c\n", str[i], str[i], str[i])
	/*
		ㅇ 한글 깨짐 이유
			> type이 uint8(1byte)이기 때문에 한글(3byte)는 못드감
		ㅇ 해결방법
			**[]rune**타입 or **range** 를 이용한다.
			깨짐현상을 막기 위한 **[]rune**으로 타입 변환

		ㅇ **[]rune** 의 단점
			변환을 위해 별도의 배열에 할당 >>불필요한 메모리 사용
			그래서 **range**를 활용
	*/
	//}
	// **[]rune** 방식
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입 : %T , 값 : %d , 문자 : %c\n", arr[i], arr[i], arr[i])
	}

	// **range** 방식
	for _, v := range str { //인덱스 무효화
		fmt.Printf("타입 : %T , 값 : %d , 문자 : %c\n", v, v, v)
	}

	//ㅁ 문자열 연산
	var str1 string = "경일"
	var str2 string = "게임"

	var str3 = str1 + " " + str2 //st1,공백,str2를 잇는다.
	fmt.Println(str3)
	str1 += " " + str2 //str1에다가 " "+str2를 합한 결과를 붙힌다.
	fmt.Println(str3)

	// ㅁ 문자열 비교
	str4 := "hello"
	str5 := "hello0"
	str6 := "hello"
	fmt.Println(str4 == str5)
	fmt.Println(str5 == str6)
	fmt.Println(str4 == str6)
	fmt.Println(str4 != str6)
	fmt.Println(str4 != str5)

	// < , > , <= , >=
	str7 := "AAA"
	str8 := "fdfAAA"
	// 대소비교의 기준
	// 길이 상관없이 문자열 첫글자부터 하나씩 아스키코드값으로 비교
	// 동일할 시 문자열 두번째로 넘어감
	fmt.Println(str7 > str8)
	fmt.Println(str7 < str8)

}
