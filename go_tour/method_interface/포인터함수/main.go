package main

import (
	"fmt"
	"math"
)

// 함수로 재적상된 Abs와 Sacle 함수가 있다
// 만약 Scale의 v인자가 포인터가 아닌 값으로 변경될 경우엔
// main 함수의 Scale 인자를 &v에서 v로 바꿔야 한다

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	Scale(&v, 10) // &v란 v를 가리키는 포인터다
	fmt.Println(v)
	fmt.Println(Abs(v))
}
