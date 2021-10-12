package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func SortNumb(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	slice := []string{}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}
	sort.Strings(slice)

	for _, i := range slice {
		fmt.Println(i)
	}

}
