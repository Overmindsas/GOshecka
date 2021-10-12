package pkg

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func CheckSort(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	n := []string{}
	m := []string{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		m = append(m, fileScanner.Text())
		n = append(n, fileScanner.Text())
	}

	sort.Strings(n)
	var counter int
	for i := 0; i < len(m)-1; i++ {

		if n[i] == m[i] {
			counter++
		}
	}
	if counter == len(m)-1 {
		fmt.Println("файлы отсортированы")
	} else {
		fmt.Println("файлы не отсортированны")
	}

}
