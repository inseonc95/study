package main

import (
	"fmt"
	"math"
)

// interfaces
// interface 타입은 "메소드의 시그니처 집합"으로 정의된다
// 인터페이스 타입의 값은 해당 메소드를 구현하는 모든 값을 보유할 수 있다
// 즉 Abser 타입의 값은.... Abs()메소드를 구현하기 위한 모든 값을 보유할 수 있다

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())

}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
