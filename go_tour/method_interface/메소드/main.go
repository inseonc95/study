package main

import (
	"fmt"
	"math"
)

// go는 클래스를 가지고 있지 않지만 그와 같은 타입의 메소드를 정의할 수 있다
// 메서드는 특별한 reciver 인자가 있는 함수이다

// 그 reciver는 func 키워드와 메서드 이름 사이에 자체 인수 목록에 나타낸다
// 이 예시에서 Abs 메소드는 v라는 이름의 Vertex 유형의 리시버가 존재한다

// 즉 다시 말해 메서드는 리시버라는 인수가 있는 함수이다 !!

// 구조체가 아닌 형식에 대해서도 메소드를 선언할수 있다
// 메소드와 동일한 패키지에 유형이 정의된 "리시버"가 있는 메소드만 선언할 수 있다
// 유형이 다른 패키지에 정의된 리시버로는 메소드를 선언할 수 없다

type Vertex struct {
	X, Y float64
}

func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4} // 25
	fmt.Println(v.AbsMethod())
	fmt.Println(Abs(v))
}
