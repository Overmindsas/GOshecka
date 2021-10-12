package main

import (
	"flag"
	"log"
	"mycut/pkg"
	"strconv"
	"strings"
)

var (
	fields    string
	delimiter string
	separated bool
	ot        int
	do        int

	fileName string
)

func init() {
	flag.StringVar(&fields, "f", "0", "fields")
	flag.StringVar(&delimiter, "d", "	", "delimiter")
	flag.BoolVar(&separated, "s", false, "separated")

	flag.StringVar(&fileName, "file", "", "filename")
}

func main() {
	flag.Parse()
	for true {
		if delimiter != "" {
			fields := []int{0, pkg.UzDo(fileName + ".txt")}
			pkg.Cut(fileName+".txt", fields[0], fields[1], delimiter)
			break
		}
		if fields != "" && delimiter != "" {
			Fild(fields, delimiter)
			break
		}
		if separated == true {
			fields := []int{0, pkg.UzDo(fileName + ".txt")}
			pkg.Cut(fileName+".txt", fields[0], fields[1], delimiter)
			break
		}
	}

}

func Fild(fields, delimiter string) {
	diap := []int{}
	contSymp := strings.Contains(fields, "-")
	if contSymp == true {
		slice := strings.Split(fields, "-")
		for _, i := range slice {
			m, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			diap = append(diap, int(m))
		}
		pkg.Cut(fileName+".txt", diap[0], diap[1], delimiter)

	} else {
		m, err := strconv.Atoi(fields)
		if err != nil {
			log.Fatal(err)
		}
		diap = append(diap, int(m))
		pkg.Cut(fileName+".txt", diap[0], diap[0]+1, delimiter)
	}
}
