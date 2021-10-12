package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func SortKolon(fileName string, kolonNumber int) {
	card := make(map[interface{}][]string)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	keysKolon := []string{}
	fileScanner := bufio.NewScanner(file)
	i := 0

	for fileScanner.Scan() {
		keys := []string{}
		words := strings.Split(fileScanner.Text(), " ")
		for _, m := range words {
			keys = append(keys, string(m))
		}

		card[keys[kolonNumber]] = keys
		i++
		keysKolon = append(keysKolon, keys[kolonNumber])
	}

	sort.Strings(keysKolon)
	for _, i := range keysKolon {
		fmt.Println(i, card[i])
	}
}
