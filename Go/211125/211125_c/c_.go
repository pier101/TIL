package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(d Data) { //매개변수로 Data 받음
	d.value = 100
	d.data[100] = 200
}

func PointerChange(d *Data) { //매개변수로 Data포인터를 받는다.
	d.value = 100
	d.data[100] = 200
}

//Actor구조체 생성
type Actor struct {
	Name  string
	HP    int
	Speed float64
}

//값을 초기화해서 반환하는 함수
func NewActor(name string, hp int, speed float64) *Actor {
	var actor = Actor{name, hp, speed}
	return &actor //주소값리턴
}

func main() {
	//포인터
	var p *int //p는 int타입 데이터의 메모리 주소를 가리키는 포인터 변수
	//포인터는 메모리 주소를 값으로 가지기 떄문에 ..숫자..뭐이런거안됨

	var a int = 100
	p = &a //변수 a의 메모리 주소를 포인터 변수 p에 대입

	println("포인터 변수 p의 값(a의 메모리 주소)", p)
	println("p가 가리키는 메모리의 값 : ", *p)
	*p = 300
	println("값 변경후 a의 값 : ", a)

	var x int = 20
	var y int = 10

	var p1 *int = &x //x의 메모리공간
	var p2 *int = &x //x의 메모리공간
	var p3 *int = &y //y의 메모리공간

	fmt.Printf("p1 == p2 :%v\n", p1 == p2)
	fmt.Printf("p2 == p3 :%v\n", p2 == p3)
	//포인터 변수는 초기화를 하지 않으면  default value는 nil
	//0이긴 하지만 유효하지 않는 메모리 주소값, >어떠한 공간도 가리키고 있지 않다.

	var pointer *int
	if pointer != nil {
		//포인터 변수가 nil이 아니라면 > 유효한 메모리 주소를 가리킨다!!
	}

	//그래서 포인터 왜 씀?
	//대입이나 인수 전달시 값을 복사하기 떄문에 성능상에 이슈발생. 다른공간에 복사하기 떄문에
	//변경사항이 적용 X

	var data Data
	ChangeData(data)
	fmt.Printf("value= %d\n", data.value)
	fmt.Printf("data[100]= %d\n", data.data[100])

	/* <<Misstion >>
	함수랑구조체 이용해서 원하는 데이터 출력
	*/

	PointerChange(&data) //data의 변수값이 아니라 data의 메모리 주소를 인수로 전달
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	//포인터는 메모리를 주소를 값으로 갖는다.
	//&를 이용해서 데이터의 메모리 주소를 알수 있다.
	//포인터를 이용하면 메모리 주소값으로 메모리를 컨트롤할수 있다.

	var actor = NewActor("배추도사 만만세", 100, 60.5)
	fmt.Println(actor.Name)
	fmt.Println(actor.Speed)
	fmt.Println(actor.HP)
}
