package main

import "fmt"

/*
	2-55
	3-57
	4-65
	5-16
*/

func main() {

	//pointers
	a := 2
	b := a
	a = 10
	c := &a
	*c = 5
	fmt.Println(a, b, *c)

	//array
	//go에서는 array의 길이를 명시해 주어야 한다
	names := [5]string{"nico", "wook", "soon"}
	names[3] = "good"
	names[4] = "ring"

	fmt.Println(names)

	//배열안을 비워주면 slice type 사용
	names1 := []string{"coco", "kim", "lee"}
	names1 = append(names1, "rara")
	fmt.Println(names1)

	//map
	nico := map[string]string{"name": "nico", "age": "12"}
	//range도 가능
	for key, value := range nico {
		fmt.Println(key, value)
	}
	fmt.Println(nico)
}
