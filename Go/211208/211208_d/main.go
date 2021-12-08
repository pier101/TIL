package main

import (
	"container/list"
	"fmt"
)

type Element struct {
	Value interface{} //데이터를 저장. 빈 인터페이스기 떔에 어떤 타입이든 저장 가능
	Next  *Element    //다음 요소의 주소를 저장
	Prev  *Element    //이전 요소의 주소를 저장
}

//큐 구조체 정의
type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

//-------------------------------------
type Stack struct {
	v *list.List
}

func (q *Stack) Push(val interface{}) {
	q.v.PushBack(val)
}
func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back != nil {
		return q.v.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	//리스트 : 각 데이터를 담고 있는 요소들을 포인터로 연결한 자료구조

	// v := list.New()       //리스트 생성
	// e4 := v.PushBack(4)   //뒤에 추가
	// e1 := v.PushFront(1)  //앞에 추가
	// v.InsertBefore(3, e4) //e4 앞에 추가
	// v.InsertAfter(2, e1)  //e1 다음에 추가

	// for e := v.Front(); e != nil; e = e.Next() { //정방향 순회
	// 	fmt.Print(e.Value, "->")
	// }
	// fmt.Println()
	// for e := v.Back(); e != nil; e = e.Prev() { //역방향 순회
	// 	fmt.Print(e.Value, "->")
	// }
	//-------------------------------------------
	// queue := NewQueue() //큐 만들기
	// for i := 0; i < 5; i++ {
	// 	queue.Push(i)
	// }
	// v := queue.Pop()
	// for v != nil {
	// 	fmt.Printf("%v ->", v)
	// 	v = queue.Pop()
	// }
	//--------------------------------------------
	stack := NewStack() //스택 만들기
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ->", i)
		stack.Push(i)
	}
	fmt.Println()
	v := stack.Pop()
	for v != nil {
		fmt.Printf("%v ->", v)
		v = stack.Pop()
	}
}

/*
	배열 - 요소 순서대로 나열 (c++에서 vector와 비슷) 중간삽입 비효율
	리스트 -요소 포인터로 연결 (c++에서 list) 중간삽입 용이
	ex)인덱스 활용한 접근 : 배열이 빠름
	ex)삽입/삭제 빈번한 경우 : 리스트가 빠름
	why?
	리스트 1 에서 4로 옮기려면 포인터로 바로 가리킬수 있는게 아닌
	리스트2, 리스트3을 거쳐서 가야하기 떄문에 불편 >>배열 용이

	데이터 지역성 특성 떄문에 data가 작은 경우 배열이 더 편리
*/

//지금 해볼 것 . 리스트 이용해 스택 만들기
