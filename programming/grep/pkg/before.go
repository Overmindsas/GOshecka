package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Before(fileName string, before bool, str string, ignoreCase bool, count bool, lineNum bool) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	slice := []string{}
	firstSovp := 0

	for firstSovp < 1 {
		for fileScanner.Scan() {
			if ignoreCase == true {
				cont := strings.Contains(strings.ToLower(fileScanner.Text()), strings.ToLower(str))
				if cont == true {
					firstSovp++
					break
				} else {
					slice = append(slice, strings.ToLower(fileScanner.Text()))
				}
			} else {
				cont := strings.Contains(fileScanner.Text(), str)
				if cont == true {
					firstSovp++
					break
				} else {
					slice = append(slice, fileScanner.Text())
				}
			}

		}
	}

	for _, i := range slice {
		fmt.Println(i)
	}

	if count == true {
		fmt.Println(count)
	}

}
