package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")

	fmt.Println(i)
}

// выведется [3 5 3 4 6] и [3 2 3]
// при использовании функии append создаётс новый слайс и работа идёт уже с ним
// Так как слайс передаётся по ссылке,  i[0] = "3" изменит заменит элемент в слайсе. После append работа идёт с другим слайсом
//
