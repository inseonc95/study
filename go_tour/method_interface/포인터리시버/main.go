package main

import (
	"fmt"
	"math"
)

// 포인터 리시버로도 메소드를 선언할수 있다
// 이는 리시버 유형이 일부 유형(타입) T에 대한 리터럴 구문  *T를 가진다는 것을 의미한다
// 또한 T자체는 *int와 같은 포인터가 될 수 없다 ????????

// 예를 들어 Scale 메소드는 *Vertex(포인터 리시버)에 정의되어 있다
// 포인터 리시버가 있는 메소드는 Sclae처럼 리시버가 가리키는 값을 수정(변경)할 수 있다 ******* ///

// 메소드는 종종 리시버를 수정해야하기에 포인터 리시버가 값 리시버보다 더 일반적이다

// Scale 메소드에서 포인터 리시버 대신 값 리시버를 사용하면 프로그램이 어떻게 동작할까
// 값 리시버를 사용하면 Scale 메서드가 원래의 Vertex 값의 복사본에서 작동한다
// 이것은 다른 함수 인수와 동일한 방식이다

// Scale 메소드에는 선언된 Vertex 값을 변경하기 위해 포인터 리시버가 있는 것임

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
