package main

import "fmt"

// func multiply(a, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {

// 	return len(name), strings.ToUpper(name)
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// 	fmt.Println(len(words))
// 	fmt.Println(words[1])
// }

// func lenAndUpper(name string) (lenght int, uppercase string) {
// 	defer fmt.Println("I'm done")
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func superAdd(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total

// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Println(numbers[i])
// 	// }
// }

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// switch {
	// case age < 8:
	// 	return false
	// case age > 8:
	// 	return true
	// }
	// return false
}

func main() {
	// var name string = "hyun"
	// name = "lynn"
	// name := false
	// fmt.Println(name)
	// fmt.Println(lenAndUpper("hyunmin"))
	// totalLenght, _ := lenAndUpper("hyunmin")
	// totalLenght, upperName := lenAndUpper("hyunmin")

	// fmt.Println(totalLenght, upperName)

	// repeatMe("nico", "lynn", "dal", "marl")
	// totalLenght, uppercase := lenAndUpper("hyunmin")
	// fmt.Println(totalLenght, uppercase)
	// fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7))
	// fmt.Println(canIDrink(18))
	a := 2
	b := &a
	*b = 30
	fmt.Println(a)
}
