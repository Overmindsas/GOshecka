package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(rasp("a4b6y"))
}

func rasp(s string) string {
	newString := ""
	count := 0
	l := 0

	for _, i := range s {
		numb, err := strconv.Atoi(string(i))
		if err == nil {
			count++
		} else {
			l += numb
		}

	}
	if len(s) == count {
		fmt.Println("uncorrect string")
	} else {
		for i := 0; i < len(s); i++ {

			m, err := strconv.Atoi(string(s[i]))
			if err != nil {
				newString += string(s[i])
			} else {
				count++
				v := strings.Repeat(string(s[i-1]), m-1)
				newString += string(v)
			}

		}

	}

	return newString
}

// Проходимя по строке
// Пробуем конвертировать строку в число. Если конвертируется, то повторяем предыдущий элемент строки m-1 раз и добавляем к newString.
// Если не конвертнулось, то добавляем к newString текущий элемент
