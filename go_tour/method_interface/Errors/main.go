package main

import (
	"fmt"
	"time"
)

// Go는 error 값으로 오류 상태를 표현한다
// error 타입은 fmt.Stringer와 유사한 내장 인터페이스이다

type error interface {
	Error() string
}

// fmt.Stringer와 마찬가지로 fmt 패키지는 값을 출력할때 error 인터페이스를 찾는다
// 함수는 종종 error 값을 반환하며, 호출 코드는 오류가 nil과 같은지 테스트하여 오류를 처리해야한다

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
