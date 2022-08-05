package main

import "fmt"

// Function closure

// Go 함수들은 클로저일수도 있다
// 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값이다
// 함수는 참조된 변수에 접근하여 할당될 수있다
// 이러한 의미에서 함수는 변수에 "bound" 된다

// 그 예로 adder 함수는 클로저를 반환하고, 각 클로저는 그 자체으 sum 변수에 바운드 되어있다????

func main() {
	pos, neg := adder(), adder()

	fmt.Println(pos(7))
	fmt.Println(neg(-2 * 2))
}

func adder() func(int) int {
	sum := 10
	return func(x int) int {
		sum += x
		return sum
	}
}

// adder()는 int를 요소로 받아 int를 출력하는 함수?
// adder는 클로저를 반환하고.........
