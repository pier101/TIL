package main

import (
	"container/ring"
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	// ㅁ ring
	// when?
	// : 실행취소, 리플레이시

	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}
	for k := 0; k < n; k++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}
	fmt.Println()
	for k := 0; k < n; k++ {
		fmt.Printf("%c", r.Value)
		r = r.Prev()
	}
	fmt.Println()
	//-------------------------
	/*
		ㅁ map
		key와 value 쌍으로 이루어져 있다.
		map은 기본 내장 type이라 package 따로 안 불러옴
		map은 처리 속도 빠름
		단점 : 리스트와 배열에 비해 메모리 많이 잡아먹음
	*/

	//m := make(map[type]type)
	m := make(map[string]string) //map 생성
	// 			   key   value

	m["김동욱"] = "대한민국"   //key,value 추가
	m["톰 홀랜드"] = "미국"   //key,value 추가
	m["기무라 타쿠야"] = "일본" //key,value 추가

	m["김동욱"] = "캐나다" //value 변경 가능
	fmt.Printf("김동욱이 사는 곳 : %s\n ", m["김동욱"])
	//--------------------------------------
	mm := make(map[int]Product) //key,value 타입으로 어떤 값도 가능

	mm[10] = Product{"가습기", 500}
	mm[20] = Product{"청소기", 1000}
	mm[13] = Product{"세탁기", 1500}
	mm[33] = Product{"전자렌지", 800}

	//map에서의 range loop는 첫번째 :key,두번째: value
	//(다른데선 range loop는 첫번째 :index,두번째: key)
	for k, v := range mm {
		fmt.Println(k, v)
	}
	ml := make(map[int]int)
	ml[1] = 0
	ml[2] = 2
	ml[3] = 3
	delete(ml, 3)
	delete(ml, 4) // 없는 녀석 지우면
	fmt.Println(ml[3])
	fmt.Println(ml[4]) //0 나옴,디폴트값이 0이기 때문

}
