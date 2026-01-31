package main

import (
	"fmt"
	"os"
)

func main() {

	a := 10
	b := 5.60
	fp := float64(a)
	i := int(b)
	f := true
	var err error = nil
	p := &a

	fmt.Printf("%8T %[1]v\n", a)
	fmt.Printf("%8T %[1]v\n", b)
	fmt.Printf("%8T %[1]v\n", f)
	fmt.Printf("%8T %[1]v\n", err)
	fmt.Printf("%8T %[1]v\n", p)
	fmt.Printf("%8T %[1]v\n", fp)
	fmt.Printf("%8T %[1]v\n", i)

	var sum float64
	var count int

	for {
		var val float64
		if _, err = fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}

		sum += val
		count++

	}
	if count == 0 {
		fmt.Fprintf(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stdout, "Sum %8f \nAverage:%8f \n", sum, sum/float64(count))

}
