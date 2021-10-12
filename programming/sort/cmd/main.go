package main

import (
	"flag"
	"mysort/pkg"
)

var (
	sortRev       bool // -r сортировать в обратном порядке
	sortNumb      bool // -n — сортировать по числовому значению
	povElem       bool // -u не выводить повторяющиеся строки
	sortK         bool // -k указание колонки для сортировки
	sortMounth    bool // -M сортировать по названию месяца
	sortHvostProb bool // -b игнорировать хвостовые пробелы
	сheckSort     bool // -c проверять отсортированы ли данные
	sortChislSuff bool // -h сортировать по числовому значению с учётом суффиксов

	fileName    string // -fn
	kolonNumber int    // -kn номер колонки
)

func init() {
	flag.BoolVar(&sortRev, "r", false, "sort reverse")
	flag.BoolVar(&povElem, "u", false, "pov elemet")
	flag.BoolVar(&sortK, "k", false, "sort po kolonke")
	flag.BoolVar(&сheckSort, "c", false, "check sort")
	flag.BoolVar(&sortMounth, "m", false, "sort mouth")
	flag.BoolVar(&sortNumb, "n", false, "sort number")

	flag.StringVar(&fileName, "file", "", "filename")
	flag.IntVar(&kolonNumber, "kn", 0, "kolon number")
}

func main() {
	flag.Parse()
	for true {
		if sortRev == true {
			pkg.SortReverse(fileName + ".txt")
			break
		}
		if sortK == true {
			pkg.SortKolon(fileName+".txt", kolonNumber)
			break
		}
		if povElem == true {
			pkg.PovElem(fileName + ".txt")
			break
		}
		if сheckSort == true {
			pkg.CheckSort(fileName + ".txt")
			break
		}
		if sortMounth == true {
			pkg.SortMouth(fileName + ".txt")
			break
		}
		if sortNumb == true {
			pkg.SortNumb(fileName + ".txt")
		}

	}

}
