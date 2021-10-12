package main

import (
	"fmt"
	"sort"
	"strings"
)

var stirngMap = make(map[string][]string)

func main() {
	slice := []string{
		"пятак",
		"пятка",
		"тяпка",
		"листок",
		"столик",
		"слиток",
		"гоблин",
		"блин",
	}

	Anag(slice)
	for key, value := range stirngMap {
		fmt.Println(key, value)
	}

}

func Anag(slice []string) {
	if len(slice) == 0 {
		fmt.Println("Итог:")
	} else {
		newSlice := []string{}
		sliceToMap := []string{}
		sl := slice[0]
		b := strings.Split(sl, "")
		sort.Strings(b)
		j := strings.Join(b, "")
		for _, i := range slice {
			a := strings.Split(i, "")
			sort.Strings(a)
			m := strings.Join(a, "")
			if j != m {
				newSlice = append(newSlice, i)
			} else {
				sliceToMap = append(sliceToMap, i)
			}
		}
		stirngMap[sl] = sliceToMap
		Anag(newSlice)
	}
}
