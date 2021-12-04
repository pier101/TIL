package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//naked return : return할 variable을 굳이 명시하지 않아도 됨
func lenAndUpper1(name string) (length int, uppercase string) {
	//defer : 함수가 끝난 후 무언가를 실행할 수 있음
	defer fmt.Println("i'm done ")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	//const name string = "dongwook"
	//var name string = "dongwook"
	//위랑 같은 말
	//name := "dongwook" //이런식으로 작성 시 go가 알아서 type찾아줌 ,그 type은 변경 못함
	//name = "wook"
	//fmt.Println(name)

	//fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("dondwook")
	fmt.Println(totalLength, upperName)
	totalLength, up := lenAndUpper1("wook")
	fmt.Println(totalLength, up)

	repeatMe("nico", "wook", "soon", "mark")

}
