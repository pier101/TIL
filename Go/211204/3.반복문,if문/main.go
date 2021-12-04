package main

import "fmt"

//loop
func superAdd(numbers ...int) int {
	//range : array에 loop를 적용할 수 있도록 해줌
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//if..
func canIDrink(age int) bool {
	//조건 앞에 변수를 생성할 수도 있다.
	if KoreanAge := age + 2; KoreanAge < 18 {
		return false
	} else {
		return true
	}
}

//switch
func canIDrink1(age int) bool {
	switch KoreanAge := age + 2; KoreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
	// switch {
	// case age < 18:
	// 	return false
	// case age >= 18:
	// 	return true
	// }
	// return false
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(15))
	fmt.Println(canIDrink1(24))
}
