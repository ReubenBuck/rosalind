package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(x float64) float64 {
	x1 := int(x)
	f := 1
	for i := 1; i < x1+1; i++ {
		f = f * i
	}
	return float64(f)
}

func choice_sample(n float64, N float64) float64 {

	// should work out a better way of canceling this
	// can do this by being smater
	return factorial(N) / (factorial(n) * factorial(N-n))
}

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

	pa := ((choice_sample(1.0, k) * choice_sample(1.0, m+n)) / choice_sample(2.0, r)) + (choice_sample(2.0, k) / choice_sample(2.0, r))

	pb := (choice_sample(2.0, m) / choice_sample(2.0, r)) * 3.0 / 4.0

	pc := (choice_sample(1.0, m) * choice_sample(1.0, n)) / choice_sample(2.0, r) / 2.0

	fmt.Println(pa + pb + pc)

}
