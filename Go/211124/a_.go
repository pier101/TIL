package main

import "fmt"

//int 타입을 ColorType로 쓰겠다 별명처럼
type ColorType int

//const 열거형
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "undefined"
	}
}

func GetColor() ColorType {
	return Blue
}

///----------------------------
type Direction int

const (
	East Direction = iota
	South
	North
	West
)

func directionValue(angle Direction) string {
	switch {
	case 0 < angle && angle <= 45:
		return "East"
	case 45 < angle && angle <= 100:
		return "South"
	case 100 < angle && angle <= 200:
		return "North"
	case 200 < angle && angle <= 300:
		return "West"
	default:
		return "undefined"
	}
}
func GetDirection() Direction {
	return 105
}

func main1() {
	fmt.Println("색깔", colorToString(GetColor()))
	fmt.Println("방향", directionValue(GetDirection()))
	/*
		<<misstion>>
		함수 이름은 GetDirection
		함수매개변수는 angle float64받음
		함수 결과는 Direction타입 반환

		angle이 300이상이먄 north return
		angle이 0이상 45 작으면 north
		angle이 45이상 100 작으면 east
		angle이 100이상 200 작으면 east

	*/
}
