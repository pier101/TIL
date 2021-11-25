package main

import "fmt"

func find20(a int) (int, bool) {
	for i := 1; i < 9; i++ {
		if a*i == 20 {
			return i, true
		}
	}
	return 0, false
}

func main2() {
	/*
		1.초기문 생략
		for ;조건; 증감 {
		}

		2.증감 생략
		for 초기문; 조건t; {
		}

		3.조건문만 있는경우
		for ;조건 ; {
		}

		4.무한루프(JS에서 while과 동일)
		for true {  //true생략가능
		}
		ex)while문(go에는 없다)
		while(true){
		}

		**continue 사용
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue
			}
			if i == 6 {
				break
			}
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < i+1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	*/

	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		b = 1
		dan++
		if dan == 10 {
			break
		}
	}
	//------------------------------------------------

	//flag변수
	// var find bool = false
	//flag변수 : for문 구조 깊어질 경우 빠져나가기 위해 씀
	//flag변수 여러개 쓰면 헷갈림
	//flag변수 안 쓰고 처리하는 법 > lable
	/*
		i := 1
		j := 1
		for ; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				if i*j == 45 {
					find = true
					break
				}
			}
			if find {
				break
			}
		}
		fmt.Printf("%d*%d= %d\n", i, j, i*j)

	*/

	// a := 1
	// // b:=1
	// Outerfor: //레이블 정의
	// for ; a <= 9; a++ {
	// 	for b := 1; b <= 9; b++ {
	// 		if a*b == 45 {
	// 			break Outerfor //레이블에 가장 먼저 포함된 for문 까지 종료
	// 		}
	// 	}
	// }
	// fmt.Printf("%d*%d= %d\n", a, b, a*b)
	// //가독성을 해치므로 자주사용 x
	// //-------------------------------------------------------------
	// a := 1
	// // y := 0
	// for ; a <= 9; a++ {
	// 	var find bool
	// 	if y, find = find20(a); find {
	// 		break
	// 	}
	// }
}
