package main

import (
	"fmt"
	"sort"
)

func main() {

	//삭제
	slice := []int{1, 2, 3, 4, 5, 6}
	index := 2

	for i := index + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	//삽입
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice1 = append(slice1, 0)

	//맨 뒤부터 추가하려는 위치까지 값을 하나씩 옮긴다.
	for i := len(slice) - 2; i >= index; i-- {
		slice[i+1] = slice[i]
	}
	slice[index] = 299
	fmt.Println(slice)

	//정렬
	s := []int{323, 5, 11, 1}
	sort.Ints(s)
	fmt.Println(s)
}
