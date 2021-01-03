package main

import "fmt"

func main() {
	// 1. Constants(상수)는 값을 변경할 수 없는 수를 의미한다.
	const name string = "hyunmin"

	// * name 변수의 값을 변경할 수 없으므로 아래 작업은 에러가 발생한다.
	// name = "sangjun"

	// 2. Variables(변수)는 값을 변경할 수 있는 수를 의미한다.
	var name2 string = "hyunmin"
	name2 = "sangjun"

	// 3. 결과 확인
	fmt.Println(name)
	fmt.Println(name2)
}
