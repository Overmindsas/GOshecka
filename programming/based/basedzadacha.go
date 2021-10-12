package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("hui: ", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
