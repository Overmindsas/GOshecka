package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func SortReverse(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	n := []string{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		n = append(n, fileScanner.Text())
	}
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	for _, i := range n {
		fmt.Println(string(i))
	}
}
