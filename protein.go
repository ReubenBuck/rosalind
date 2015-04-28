// protein translation
// how long does it take to index something, vs how long it takes to
// to actually read through

// struct may not be the way to do this a map might be better

// B = Stop

package main

import (
	"fmt"
	"os"
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

	seq1 := []byte(os.Args[1])
	Len := len(seq1) / 3
	prot := make([]byte, Len)

	for i := 0; i < Len; i++ {
		codon := string(seq1[(i * 3):((i * 3) + 3)])
		if trans[codon] == '*' {
			break
		}
		prot[i] = trans[codon]
	}

	fmt.Println(string(prot))

	//fmt.Println(len(aminoAcid))

}
