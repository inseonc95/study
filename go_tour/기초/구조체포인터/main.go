package main

import "fmt"

// 구조체 포인터를 통해서 구조체 필드를 접근할 수 있다
// (*p).X 로 작성하면 구조체 포인터 p에서 해당 구조체의 X 필드에 접근할 수 있다
// 하지만 이 표기법은 번거로울수 있기 때문에 이 언어는 역참조를 명시할 필요 없이 그냥 p.X로 작성 가능하다

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v // 구조체 v를 바라보는 포인터 p 생성

	p.X = 1e9 // 포인터를 사용하여 구조체 v의 X 필드값 변경
	fmt.Println(v)
}
