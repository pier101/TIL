package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 200
}
func changeSlice(slice []int) {
	slice[2] = 200
}

func main() {
	//slice 내부적으로 이렇게 생김
	type SliceHeader struct {
		Data uintptr //실제 배열을 가리키는 포인터
		Len  int     //요소 갯수
		Cap  int     //실제 배열의 길이
	}
	//**사진_slice구조

	//var sclice = make([]int, 3)
	//사진_사진1
	//var sclice1 = make([]int, 3, 5)
	//사진_사진2

	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	changeArr(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)

	slice3 := make([]int, 3, 5) //len:3 cap:5 슬라이스 만듬
	slice4 := append(slice3, 4, 5)
	slice5 := append(slice3, 4, 5, 6)
	slice3[1] = 200
	fmt.Println("슬라이스 3 : ", slice3, len(slice3), cap(slice3))
	fmt.Println("슬라이스 4 : ", slice4, len(slice4), cap(slice4))
	fmt.Println("슬라이스 5 : ", slice5, len(slice5), cap(slice5))

	/*
		만약 빈 공간이 없으면 더 큰 배열을 만듦. 일반적으로 기존 배열의 2배
		->기존 배열의 요소를 새로운 배열에 복사->새로운 배열 맨 뒤에 새 값을 추가
		cap : 새로운 배열의 길이 값
		len : 기존 길이에 추가한 개수만큼 더한 값
		포인터 : 새로운 배열을 가리키는 슬라이스 구조체를 return
	*/

	//슬라이싱 : 배열의 일부를 집어낸다
	//array[startIndex : endIndex]
	//주의할 점 : endIndex - 1 (js의 slice와 비슷하네)
	arr := [5]int{1, 2, 3, 4, 5}
	slice6 := arr[1:3] //슬라이싱
	fmt.Println(slice6, len(slice6), cap(slice6))

	arr[1] = 100
	fmt.Println(slice6, len(slice6), cap(slice6))
	slice6 = append(slice6, 300)
	fmt.Println(slice6, len(slice6), cap(slice6))
	///////////////////////////////////////////

	//슬라이스를 슬라이싱
	//slice7 := []int{1, 2, 3, 4, 5}
	// slice8 := slice[1:2] //[2]

	// //처음부터(startIndex 생략가능)
	// slice8 := slice[0:3] //[1 2 3]

	// //끝까지 (endIndex 생략가능)
	// slice9 := slice7[2:len(slice7)] //[3 4 5]
	// slice9 := slice7[2:]            //[3 4 5]

	// //전체 (주로 배열 전체 가리키는 슬라이스 만들떄 사용)
	// slice10 := slice7[:] //[1 2 3 4 5]

	// //cap크기 조절하기(인덱스 3개로)
	// //slice[startIndex : endIndex : maxIndex]
	// slice11 := slice7[1:3:4] // [2 3] cap:3 (maxIndex - startIndex)
	///////////////////////////////////////////

	slice12 := []int{1, 2, 3, 4, 5}
	slice13 := make([]int, len(slice12)) //slice12 랑 같은 길이의 slice 생성

	//slice12의 모든 값을 복사중
	for i, v := range slice12 {
		slice13[i] = v
	}
	slice12[1] = 200
	fmt.Println(slice12)
	fmt.Println(slice13)

	//copy함수

	slice14 := []int{1, 2, 3, 4, 5}
	slice15 := make([]int, 3, 10) // len :3 cap:10
	slice16 := make([]int, 10)    // len :10 cap:10

	//첫번째인수--복사한 결과를 저장하는 슬라이스 변수
	//두번째인수--복사대상이 되는 슬라이스 변수
	copy1 := copy(slice15, slice14)
	copy2 := copy(slice16, slice14)
	fmt.Println(copy1, slice15)
	fmt.Println(copy2, slice16)

}
