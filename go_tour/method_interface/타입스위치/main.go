package main

import "fmt"

// type switch는 여러 타입의 선언을 직렬로 허용하는 구조이다
// 타입 스위치는 일반 스위치문과 같지만, 타입 스위치문의 경우에 "값"이 아닌 "타입"을 명시하며,
// 그 값(타입?)들은 지정된 인터페이스 값에 의해 유지되는 값의 타입과 비교된다

// swich v := i.(type) {
// case T:
// 	// here v has type T
// case S:
// 	// here v has Type S
// default:
// 	// no match; here the same type as i
// }

// 타입 스위치 선언은 타입 선언 i.(T)와 같은 구문을 가지는데, 특정 타입인 T는 type이라는 키워드로 대체된다
// 이 스위치 문장은 인터페이스 값 i가 T 타입인지 S 타입인지 시험한다
// T와 S의 각각의 경우 변수 v는 각각 T과 S형으로, i 타입이 보유한 값을 가지게 된다

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v type is int \n", v)
	case string:
		fmt.Printf("%v type is string \n", v)
	default:
		fmt.Printf("default type is %T \n", v)
	}

}

func main() {
	do(20)
	do("hello")
	do(true)
}
