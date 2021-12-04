/*
	컴파일을 위해선 main.go로 고정되어야 한다.
	컴파일시 main을 인식하고
	그 안에 package main 과
	func main이 있는지를 확인하기 떄문에
	고정으로 두어야 함
*/

package main

import (
	"fmt" //fomatting을 위한 패키지
)

func main() {
	fmt.Println("Hello") //function을 export 하기 위해 go에서는 대문자로 시작하게 작성한다.
	//something.SayHello() //대문자라 export되어서 import가능
	//something.sayBye()   //소문자라 export안되어서 import 불가능
}
