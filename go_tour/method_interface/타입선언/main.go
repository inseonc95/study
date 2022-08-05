package main

import "fmt"

// type assertion은 인터페이스 값의 기초적인 콘크리트 값에 대한 접근을 제공한다
// t := i.(T)
// 이는 인터페이스 값 i가 콘크리트 타입 T를 갖고 있으며, 그 기본값인 T값을 변수 t에 할당하고 있다고 선언한다
// 만약 i가 T를 갖지못하면 그 선언은 panic 상태가 된다

// 인터페이스 값이 특정 유형을 보유하고 있는지 여부를 테스트하기 위해
// 타입 선언시 기본값과 선언성공 여부를 보고하는 bool 값을 반환할수있다
// i, ok := i.(T)
// 이 구문과 map에서 읽는 구문간 유사성에 유의

func main() {
	var i interface{} = "hello"

	s := i.(string) // 인터페이스 i가 string 타입을 갖고 있으며, 그 기본값을 변수 s에 할당한다고 선언하는 것
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
