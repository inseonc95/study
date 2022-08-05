package main

import (
	"fmt"
	"math"
)

// hood 아래에서?? 인터페이스의 값은 콘크리트 타입의 튜플이라고 생각할 수 있다
// (value, type)
// 인터페이스 값은 특정 기초 콘크리트 유형의 가치를 가진다
// 인터페이스 값으로 메소드를 호출하면 기본 형식에 동일한 이름의 메서드가 실행된다

// 인터페이스 자체 내부 콘크리트 값이 0인 경우 그 메소드는 nil 리시버로 호출된다
// 일부 언어에서는 이것이 null 포인터 예외를 발생시키지만,
// Go에서는 nil 리시버로 호출되는 것으로 불리는 매우 좋은 방법을 사용하는 것이 일반적이다
// 아래 예시의 M이라는 방법과 같다
// nil 콘크리트 값을 갖는 인터페이스 값 자체가 nil은 아니라는 점에 주의

// nil 인터페이스 값은 값과 콘크리트 유형을 가지지 않는다
// nil 인터페이스에서 메소드를 호출하는 것은 런타임 에러이다
// 왜냐하면 어떠한 구체적인 메소드를 호출할지 나타내는 인터페이스 튜플 내부의 타입이 업기 때문

// 비어있는 인퍼테이스 값
// 0 메서드를 지칭하는 인터페이스 타입을 empty interface라고 한다
// interface{}
// 빈 인터페이스 타입은 모든 타입의 값을 가질수 있다
// 모든 타입은 최소 0개의 메소드를 구현한다

// 빈 인터페이스는 알수없는 값을 처리하는데에 이용된다
// 예를 들어 fmt.Print는 interface{} 타입의 어떠한 인수라도 취할 수 있다??

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// 비어있는 인터페이스
	var iEmpty interface{}
	describeEmpty(iEmpty)

	iEmpty = 42
	describeEmpty(iEmpty)

	iEmpty = "hello"
	describeEmpty(iEmpty)
}
