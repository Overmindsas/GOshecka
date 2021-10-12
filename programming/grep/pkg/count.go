package pkg

import (
	"log"
	"os"
)

func Count(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
