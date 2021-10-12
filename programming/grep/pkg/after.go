package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func After(fileName string, ignoreCase bool, str string, lineNum bool, count bool, invert bool) {

	file, err := os.Open(fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	numLine := 0
	newStr := strings.ToLower(str)
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if ignoreCase == true {
			textLower := strings.ToLower(fileScanner.Text())
			//textLower := strings.ToLower(fileScanner.Text())
			trueString := strings.Contains(textLower, newStr)
			if trueString == true {
				numLine++
				continue

			} else {
				if lineNum == true {
					fmt.Println(numLine, textLower)
				} else {
					fmt.Println(numLine, textLower)
				}
			}
			numLine++
		}

	}
	if count == true {
		fmt.Println(numLine)
	}
}
