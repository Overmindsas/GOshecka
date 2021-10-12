package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Context(fileName string, context bool, str string, ignoreCase bool, count bool, lineNum bool) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	sliceCount := []int{}
	sliceText := []string{}
	countNum := 0
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if ignoreCase == true {
			countNum++
			cont := strings.Contains(strings.ToLower(fileScanner.Text()), strings.ToLower(str))
			sliceText = append(sliceText, strings.ToLower(fileScanner.Text()))
			if cont == true {
				sliceCount = append(sliceCount, countNum)
			}
		} else {
			countNum++
			cont := strings.Contains(fileScanner.Text(), str)
			sliceText = append(sliceText, fileScanner.Text())
			if cont == true {
				sliceCount = append(sliceCount, countNum)
			}
		}

	}
	for _, i := range sliceText[0 : sliceCount[0]-1] {
		fmt.Println(i)
	}

	for _, i := range sliceText[sliceCount[0]+1:] {
		fmt.Println(i)
	}
	if count == true {
		fmt.Println(count)
	}
	if lineNum == true {
		fmt.Println(countNum)
	}

}
