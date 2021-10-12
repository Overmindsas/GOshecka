package main

import (
	"flag"
	"mygrep/pkg"
)

var (
	after      bool
	before     bool
	context    bool
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	fileName   string
	str        string
)

func init() {
	flag.BoolVar(&count, "c", false, "count strings")
	flag.BoolVar(&after, "A", false, "after")
	flag.BoolVar(&before, "B", false, "before")
	flag.BoolVar(&context, "C", false, "context")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&invert, "v", false, "invert")
	flag.BoolVar(&fixed, "F", false, "fixed")
	flag.BoolVar(&lineNum, "n", false, "print line num")
}

func main() {
	flag.Parse()
	for true {
		if after == true || before == true || context == true {
			if before == true {
				pkg.Before(fileName+".txt", before, str, ignoreCase, count, lineNum)
				break
			}
			if after == true {
				pkg.After(fileName+".txt", after, str, ignoreCase, count, lineNum)
				break
			}
			if context == true {
				pkg.Context(fileName+".txt", context, str, ignoreCase, count, lineNum)
				break
			}
		}
		if invert == true {
			pkg.Invert(fileName+".txt", ignoreCase, str, lineNum, count)
			break
		}
		if count == true {
			pkg.Count(fileName + ".txt")
			break
		}
	}

}
