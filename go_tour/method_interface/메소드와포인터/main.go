package main

import (
	"fmt"
)

// 이전 포인터 리시버와 포인터 함수를 비교하여 포인터 인수의 함수는 인수로 포인터를 사용해야함을 알수있다
// 포인터 리시버가 있는 메소드는 아래와 같이 호출될때 값이나 포인터를 리시버로 받아들인다
// var V vertex
// v.Scale(5)
// p := &v
// p.Scale(10)

// v.Scale(5)에서 v는 포인터가 아니라 값이지만 포인터 리시버가 있는 메서드이기 때문에 자동으로 호출된다
// 즉 Scale 메서드는 포인터 리시버를 가졌기 때문에 편의성 v.Scale(5)라고 작성하지만 Go는 (&v).Scale(5)로 해석한다

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // == (&v).Scale(f)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3} // p자체가 포인터임
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
