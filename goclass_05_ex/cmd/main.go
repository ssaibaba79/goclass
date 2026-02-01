package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const ERROR_CODE = -1

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not enought arguments")
		os.Exit(ERROR_CODE)
	}
	searchkey := os.Args[1]
	var count int

	for scan.Scan() {
		if searchkey == scan.Text() {
			count++
		}
	}

	if count > 0 {
		fmt.Fprintf(os.Stdout, "Search key %s found, count :%d \n", searchkey, count)
	} else {
		fmt.Fprintf(os.Stdout, "Search key : %s not found\n", searchkey)
	}

}
