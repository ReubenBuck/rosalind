// the rabbit problem

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read in in string form
	nn := os.Args[1]
	mm := os.Args[2]

	//convert to int for both
	n, err := strconv.Atoi(nn)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	m, err := strconv.Atoi(mm)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	// set up the slice where the sequence is stored
	var seq []int
	seq = make([]int, nn)
	seq[0] = 1
	seq[1] = 1
	//run it then print the answear
	for i := 0; i < (nn - 2); i++ {
		if i < 
		seq[i+2] = seq[i]*kk + seq[i+1]
	}
	fmt.Println(seq)

}
