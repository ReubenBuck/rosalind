// a better way might be to make a big map with all the
// names as the sequences, and the values as their strting position
// that, way we just call the sequences we want and get the positions,
// and sort

// A similar stratergy may be useful for the protein one

package main

import (
	"fmt"
	"os"
)

func main() {
	seq1 := []byte(os.Args[1])
	mot := []byte(os.Args[2])

	for i := 0; i < (len(seq1) - 1); i++ {
		if string(seq1[i:(i+len(mot))]) == string(mot) {
			fmt.Print(i+1, " ")
		}
	}
	fmt.Print("\n")

	fmt.Println(len(seq1))
}
