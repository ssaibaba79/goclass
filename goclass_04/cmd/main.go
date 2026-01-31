package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		s := "Winter is coming!"
		t := s[:6]
		i := s[6:9]
		c := s[10:]
		k := t + i + " not " + c


		fmt.Println(s, t, i, c, k)
	*/

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no enough args!")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)
		fmt.Fprintln(os.Stdout, t)
	}

}
