package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Invert(fileName string, ignoreCase bool, str string, lineNum bool, count bool) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if ignoreCase == true {
		strings.ToLower(str)
		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			trueString := strings.Contains(fileScanner.Text(), str)
			if trueString == true {
				continue
			} else {
				fmt.Println(strings.ToLower(fileScanner.Text()))
			}
		}
	} else {

	}

}
