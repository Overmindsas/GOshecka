package main

import (
	"flag"
	"fmt"
	"mywget/pkg"
)

var (
	url      string
	filePath string
)

func init() {
	flag.StringVar(&url, "u", "", "ulr website for download")
	flag.StringVar(&filePath, "file", "", "filePath for save website code")
}

func main() {
	flag.Parse()
	if url != "" && filePath != "" {
		pkg.Wget(filePath, url)
	} else {
		fmt.Println("uncorrect data or ")
	}
}
