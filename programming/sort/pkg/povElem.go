package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func PovElem(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	card := make(map[string]int)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		card[string(fileScanner.Text())] += 1
	}
	for word, _ := range card {
		fmt.Println(word)
	}
}
