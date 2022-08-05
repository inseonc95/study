package main

import (
	"fmt"
)

// interfaces의 암시적 구현(implements)

// type 구현(implements)은 메소드를 실행함으로서 인터페이스를 구현한다
// 명시적 intent의 선언도, implementation의 키워드도 없다

// 암시적 인터페이스는
// 인터페이스의 정의를 구현으로부터 분리하며
// 이는 사전 정렬 없이 어떠한 패키지에 등장할 수 있다

type I interface {
	M() // 타입도 없다??
}

type T struct {
	S string
}

// this method means type T implements the interface I
// T 타입이 인터페이스 I를 구현한다
// 하지만 이렇다는것을 명시적으로 선언할 필요는 없다

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
