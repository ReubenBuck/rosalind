package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	kk := os.Args[1]
	mm := os.Args[2]
	nn := os.Args[3]

	k1, err := strconv.Atoi(kk)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	m1, err := strconv.Atoi(mm)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	n1, err := strconv.Atoi(nn)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	k := float64(k1)
	m := float64(m1)
	n := float64(n1)

	r := k + m + n

	nom := ((k * k) - k) + (2 * k * m) + (2 * k * n) + (0.75 * ((m * m) - m)) + (m * n)
	dom := (r * r) - r
	fmt.Println(nom / dom)

}
