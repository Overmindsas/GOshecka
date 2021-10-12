package pkg

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var monthMap1 = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

var monthMap2 = map[int]string{
	1:  "JAN",
	2:  "FAB",
	3:  "MAR",
	4:  "APR",
	5:  "MAY",
	6:  "JUN",
	7:  "JUL",
	8:  "AUG",
	9:  "SEP",
	10: "OCT",
	11: "NOV",
	12: "DEC",
}

func SortMouth(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	monthSlice := []int{}
	notMothSlice := []string{}
	finalSlice := []string{}
	defer file.Close()
	n := []string{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		n = append(n, fileScanner.Text())
	}

	for _, i := range n {
		if monthMap1[i] == 0 {
			notMothSlice = append(notMothSlice, i)
		} else {
			monthSlice = append(monthSlice, monthMap1[i])
		}
	}

	sort.Ints(monthSlice)
	sort.Strings(notMothSlice)

	for _, i := range monthSlice {
		finalSlice = append(finalSlice, monthMap2[i])
	}

	for _, i := range notMothSlice {
		finalSlice = append(finalSlice, i)
	}

	for _, i := range finalSlice {
		fmt.Println(i)
	}
}

// if contain month ->
