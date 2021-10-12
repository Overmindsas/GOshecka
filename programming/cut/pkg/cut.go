package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Cut(fileName string, ot int, do int, delimiter string) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	//card := make(map[interface{}][]string)
	kolonStringSlise := []string{}
	fileScanner := bufio.NewScanner(file)
	i := 0
	for fileScanner.Scan() {
		words := strings.Split(strings.TrimSpace(fileScanner.Text()), " ")
		for _, i := range words {
			//m := strings.TrimSpace(string(i))
			kolonStringSlise = append(kolonStringSlise, i)
		}
		m := strings.Join(kolonStringSlise[ot:do], delimiter)
		fmt.Println(m)
		i++
		kolonStringSlise = []string{}
	}

}

func UzDo(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	kolonStringSlise := []string{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		words := strings.Split(strings.TrimSpace(fileScanner.Text()), " ")
		for _, i := range words {
			//m := strings.TrimSpace(string(i))
			kolonStringSlise = append(kolonStringSlise, i)
		}
		break
	}
	return len(kolonStringSlise)
}
