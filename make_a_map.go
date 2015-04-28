// make a map in go

package main

import (
	"fmt"
	//	"os"
)

func main() {

	bases := [4]byte{'U', 'C', 'A', 'G'}
	amino_acids := []byte("FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG")

	var trans map[string]byte
	trans = make(map[string]byte)
	for a := 0; a < len(bases); a++ {
		for b := 0; b < len(bases); b++ {
			for c := 0; c < len(bases); c++ {
				aa := []byte{bases[a], bases[b], bases[c]}
				trans[string(aa)] = amino_acids[len(trans)]
			}

		}

	}

	fmt.Println(trans)

	fmt.Println(amino_acids)
	fmt.Println(len(amino_acids))
	fmt.Println(bases)
}
