package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintScore(name string, math int, eng int, science int) {
	total := math + eng + science
	avg := total / 3
	fmt.Println(name, "의 평균 점수", avg, "이다")
}

////////////////////////////////

//멀티반환함수(여러개 리턴값을 가지고 있음)
//ㅁ두가지 방식 -섞어 쓰면 안됨
//매개변수 타입이 같으면 아래와 같이 생략 가능
//1.
func Devide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

//2. 변수명을 이용한 반환
func DevideA(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return //명시적으로 반환할 값을 지정하지 않음
	}
	result = a / b
	success = true
	return
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^/

//피보나치 수열 : 어떤 수열의 항이 앞의 두항의 합과 같은 수열임
//ex)1,2,2,3,5,8,13,21,34,55,89,144,233.......
//재귀

func Fibo(n int) int {
	// 재귀 함수의 탈출 조건 >>없으면 무한 반복! 반드시 있어야함 .없으면 프로그램 뻗어버림
	if n < 2 {
		return n
	}
	//재귀 부분 >>함수안에서 자기 자신 함수를 호출함
	return Fibo(n-2) + Fibo(n-1)
}

//-----------------------메인부분-----------------------
func main() {
	//인수(argument)들이 해당 함수부분 매개변수(parameter)에 복사가되는 개념
	//해당 함수가 종료되는 순간 날라가버림
	c := Add(3, 6) //함수 호출
	fmt.Println(c)

	// PrintScore("홍길동", 80, 80, 100)
	// PrintScore("김동욱", 50, 60, 80)
	// PrintScore("박지성", 70, 50, 31)

	// a, success := Devide(9, 3)
	// fmt.Println(a, success)

	//피보나치 수열의 9번째 값; 34
	fmt.Println(Fibo(9))

	/*---<<<Quest>>>---
	- 함수를 이용해서 평균값을 구해 학점 뽑기
	ex)
	90~100 A
	80~89 B
	70~79 C
	60~69 D
	-------------------*/
	//Quest submit
	avg, grade := score(99, 80)
	fmt.Println("평균점수", avg, "학점:", grade)

}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

//Quest submit
func score(eng, math int) (avg int, grade string) {

	//1번 방식
	if (eng+math)/2 <= 100 && (eng+math)/2 > 90 {
		avg = (eng + math) / 2
		grade = "A"
		return
	} else if (eng+math)/2 <= 90 && (eng+math)/2 > 80 {
		avg = (eng + math) / 2
		grade = "B"
		return
	} else if (eng+math)/2 <= 80 && (eng+math)/2 > 70 {
		avg = (eng + math) / 2
		grade = "C"
		return
	} else if (eng+math)/2 <= 70 && (eng+math)/2 > 60 {
		avg = (eng + math) / 2
		grade = "D"
		return
	} else if (eng+math)/2 <= 60 {
		avg = (eng + math) / 2
		grade = "F"
		return
	} else {
		avg = (eng + math) / 2
		grade = "오류 : 평균값 100초과"
	}
	return

	/*
		2번 방식 (1번과 결과 같음)
		if (eng+math)/2 <= 100 && (eng+math)/2 > 90 {
			return (eng + math) / 2, "A"
		} else if (eng+math)/2 <= 90 && (eng+math)/2 > 80 {
			return (eng + math) / 2, "B"
		} else if (eng+math)/2 <= 80 && (eng+math)/2 > 70 {
			return (eng + math) / 2, "C"
		} else if (eng+math)/2 <= 70 && (eng+math)/2 > 60 {
			return (eng + math) / 2, "D"
		} else if (eng+math)/2 <= 60 {
			return (eng + math) / 2, "F"
		} else {
			return (eng + math) / 2, "오류 : 평균값 100초과"
		}
	*/
}
