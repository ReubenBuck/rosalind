// look at each open readign frame

package main

import (
	"fmt"
	"os"
)

func main() {

	// we want to take our input and find all the open reading frames
	bases := []byte{'T', 'C', 'A', 'G'}
	amino_acids := []byte("FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG")
	trans := translateTable(bases, amino_acids)

	seq1F := []byte(os.Args[1])
	// need to do reverse complement
	seq1R := revC(seq1F)
	fmt.Println(string(seq1F))
	fmt.Println(string(seq1R))

	startF := startFind(seq1F)
	// now we need something to store all of our DNA strings in
	// a slice of slices to store each sequence and translate it
	// remeber to only return if there is a stop codon

	fmt.Println(startF)

	Len := len(seq1F) / 3
	prot := make([]byte, Len)

	for i := 0; i < Len; i++ {
		codon := string(seq1F[(i * 3):((i * 3) + 3)])
		if trans[codon] == '*' {
			break
		}
		prot[i] = trans[codon]
	}

	fmt.Println(string(prot))

	//fmt.Println(len(aminoAcid))

}

func startFind(sequence []byte) []int {
	a := make([]int, 1)
	mot := []byte("ATG")
	for i := 0; i < (len(sequence) - 1); i++ {
		if string(sequence[i:(i+len(mot))]) == string(mot) {
			a = append(a, i+1)
		}
	}
	return (a[1:])
}

// reverse complement package where we use a switch

func revC(sequence []byte) []byte {
	sequenceR := make([]byte, len(sequence))
	for i := 0; i < len(sequence); i++ {
		switch sequence[i] {
		case 'A':
			sequenceR[len(sequenceR)-i-1] = 'T'
		case 'T':
			sequenceR[len(sequenceR)-i-1] = 'A'
		case 'C':
			sequenceR[len(sequenceR)-i-1] = 'G'
		case 'G':
			sequenceR[len(sequenceR)-i-1] = 'C'

		}

	}
	return (sequenceR)
}

func translateTable(bases []byte, amino_acids []byte) map[string]byte {
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
	return (trans)
}
