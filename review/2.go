package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

// в начале функция test заекончится на x=1, увеличестя на 1 и вернёт 2
// в функции anotherTest вернёт x=1 и выведет его на экран и только потом увелечит на 1
