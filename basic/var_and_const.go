package main

import "fmt"

func main() {
	// 1. Constants(상수)는 값을 변경할 수 없는 수를 의미한다.
	const name string = "hyunmin"

	// * name 변수의 값을 변경할 수 없으므로 아래 작업은 에러가 발생한다.
	// name = "sangjun"

	// 2. Variables(변수)는 값을 변경할 수 있는 수를 의미한다.
	var name2 string = "hyunmin"
	// * func(함수)안에서는 변수를 선언할때 타입을 지정하지 않아도 자동으로 초기화 작업에 의해 타입이 정해진다.
	// * var name2 string = "hyunmin" 와 동일
	// name2 := "hyunmin"
	name2 = "sangjun"

	// 3. 결과 확인
	fmt.Println(name)
	fmt.Println(name2)
}
