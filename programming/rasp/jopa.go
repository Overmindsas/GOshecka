package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123441241"
	count := 0
	l := 0
	for _, i := range str {
		numb, err := strconv.Atoi(string(i))
		if err == nil {
			count++
		} else {
			l += numb
		}

	}
	if len(str) == count {
		fmt.Println("uncorrect string")
	}
}
