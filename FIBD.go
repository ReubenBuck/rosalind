/*
This is to have a go at the fib rabbits if they die
the idea is that they breed at a certin rate then
die at a certain rate

So the simulation works but poorly
it scales quadraticly

is there a way to make it scale linery like in R

the problem is we look through each element of the slice

access a look up table

arrange in sorted order

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read in in string form
	n0 := os.Args[1]
	m0 := os.Args[2]

	//convert to int for both
	n1, err := strconv.Atoi(n0)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	m1, err := strconv.Atoi(m0)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	m := m1 - 1
	n := n1 - 1

	// set up the simulation
	seq1 := []int{0}
	seq2 := []int{0}
	//run it then print the answear
	for i := 0; i < n; i++ {
		// run through each element of the slice
		seq3 := []int{}
		for e := 0; e < len(seq1); e++ {
			if seq1[e] == 0 {
				// change the 0 to a 1
				seq2 = []int{1}
			} else if seq1[e] == m {
				// change the m to a zero
				seq2 = []int{0}
			} else {
				// add 1 to the number and make a zero
				seq2 = []int{seq1[e] + 1, 0}
			}
			seq3 = append(seq3, seq2...)

		}
		seq1 = seq3
		fmt.Println(i, len(seq1))
	}

	//fmt.Println(seq1)
	fmt.Println(len(seq1))

}
