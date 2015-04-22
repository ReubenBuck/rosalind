// calculate the hamming distance

package main

import (
	"fmt"
	"os"
)

func main() {
	seq1 := []byte(os.Args[1])
	seq2 := []byte(os.Args[2])

	noteq := 0
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			noteq++
		}
	}
	fmt.Println(noteq)
}
