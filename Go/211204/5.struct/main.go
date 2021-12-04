package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//struct : object와 비슷하면서 map보다 좀 더 유연함
	favFood := []string{"kimchi", "ramen"}
	wook := person{"wook", 29, favFood}
	fmt.Println(wook)
	fmt.Println(wook.age)
	// 위의 struct는 별로 좋지 않은 코드
	wook1 := person{name: "wook", age: 29, favFood: favFood}
	fmt.Println(wook1)

}
