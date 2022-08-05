package main

import (
	"fmt"
	"math"
)

// 이와 동등한 일은 역방향에서 일어날 수 있다
// 값 인수를 사용하는 함수는 특정 유형의 "값"을 인자로 사용해야 한다

// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!

// 값 리시버를 사용하는 메서드는 값이나 포인터를 라사버로 사용한다

// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK

// 이 경우 p.Abs()라는 메서드는 (*p).Abs()로 해석된다

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

}
