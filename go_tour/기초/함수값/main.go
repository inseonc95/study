package main

import (
	"fmt"
	"math"
)

// Function Value
// 함수 그 자체도 값이기 때문에 다른 값들과 마찬가지로 전달될 수 있다
// 함수 값은 함수의 인수나 반환 값으로 사용될 수 있다

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
