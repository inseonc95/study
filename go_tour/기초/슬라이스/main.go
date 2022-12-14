package main

import "fmt"

// [n]T 타입은 타입이 T인 n (개?) 값들의 배열
// var a [10]int
// 위 표현은 변수 a룰 "10개의 정수들의 배열"로 선언한 것
// 배열의 길이는 그 타입의 일부이기 때문에 조정할 수 없다

// 배열은 고정된 크기를 가지고 있지만 반면에 슬라이스는 배열의 요소들을 동적인 크기로 볼수 있다
// 슬라이스는 배열보다 훨씬 흔하다

// []T 타입은 T 타입을 원소로 가지는 슬라이스
// 슬라이스는 콜론으로 구분된 두개의 인덱스(하한과 상한)을 지정할 수 있다
// a[low:high]
// 이것은 첫번째 요소를 포함하지만 마지막 요소를 제외하는 범위를 선택한다

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// 슬라이스는 어떤 데이터도 저장할 수 없다
// 이것은 단지 기본 배열의 한 영역을 나타낼 뿐이다
// 슬라이스의 요소를 변경하면 기본 배열의 요소가 수정된다
// 동일한 기본 배열을 공유하는 다른 슬라이스는 이러한 변경사항을 볼 수있다

// 슬라이스 리터럴은 길이가 없는 배열 리터럴 이다
// 아래는 배열 리터럴이다
// [3]bool{true, true, false}

// 아래는 위와 동일한 배열이 생성되고, 이를 참조하는 슬라이스가 만들어 진다
// []bool{true, true, false}

// 슬라이스는 _length(길이)와 _capacity(용량)을 둘 다 가지고 있다
// 슬라이스 길이란 슬라이스가 포함하는 요소의 개수
// 슬라이스 용량이란 슬라이스의 첫번째 요소부터 계산하는 기본 배열의 요소 개수이다
// 슬라이스의 길이를 연장하기 위해선 슬라이스에 충분한 용량이 있을때 다시 슬라이싱 할 수 있다
// 예제의 슬라이스 작업중 하나를 슬라이스의 용량을 넘어가도록 변경하여 어떤 일이 발생하는지 확인

// 슬라이스는 내장된 make 함수로 생성할 수 있다
// 이것은 동적 크기의 배열을 생성하는 방법이다
// make 함수는 0으로 이루어진 배열을 할당하고, 해당 배열을 참조하는 슬라이스를 반환한다

// a := make([]int, 5) (int 타입의 배열을 나타내는 슬라이스 a의 길이는 5)
// b := make([]int, 0, 5) (슬라이스 b의 길이는 5, 용량은 5)
// b = b[:cap(b)]   (슬라이스 b의 길이는 5, 용량은 5)
// b = b[1:] (슬라이스 b의 길이는 4, 용량도 4)

// 슬라이스에 새로운 요소를 추가하는것이 일반적이다
// append 결과값은 원래 슬라이스의 모든 요소와 추가로 제공된 값들을 포함하는 슬라이스
// 만약 원래 슬라이스 배열이 너무 작아서 주어진 값들을 모두 추가할수 없는 경우 더 큰 배열이 할당되고,
// 이때 반환된 슬라이스는 새롭게 할당된 배열을 가리킨다
