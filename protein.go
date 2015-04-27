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

	aminoAcid := map[string]byte{
		"UUU": 'F',
		"CUU": 'L',
		"AUU": 'I',
		"GUU": 'V',
		"UUC": 'F',
		"CUC": 'L',
		"AUC": 'I',
		"GUC": 'V',
		"UUA": 'L',
		"CUA": 'L',
		"AUA": 'I',
		"GUA": 'V',
		"UUG": 'L',
		"CUG": 'L',
		"AUG": 'M',
		"GUG": 'V',
		"UCU": 'S',
		"CCU": 'P',
		"ACU": 'T',
		"GCU": 'A',
		"UCC": 'S',
		"CCC": 'P',
		"ACC": 'T',
		"GCC": 'A',
		"UCA": 'S',
		"CCA": 'P',
		"ACA": 'T',
		"GCA": 'A',
		"UCG": 'S',
		"CCG": 'P',
		"ACG": 'T',
		"GCG": 'A',
		"UAU": 'Y',
		"CAU": 'H',
		"AAU": 'N',
		"GAU": 'D',
		"UAC": 'Y',
		"CAC": 'H',
		"AAC": 'N',
		"GAC": 'D',
		"UAA": 'B',
		"CAA": 'Q',
		"AAA": 'K',
		"GAA": 'E',
		"UAG": 'B',
		"CAG": 'Q',
		"AAG": 'K',
		"GAG": 'E',
		"UGU": 'C',
		"CGU": 'R',
		"AGU": 'S',
		"GGU": 'G',
		"UGC": 'C',
		"CGC": 'R',
		"AGC": 'S',
		"GGC": 'G',
		"UGA": 'B',
		"CGA": 'R',
		"AGA": 'R',
		"GGA": 'G',
		"UGG": 'W',
		"CGG": 'R',
		"AGG": 'R',
		"GGG": 'G',
	}

	seq1 := []byte(os.Args[1])
	Len := len(seq1) / 3
	prot := make([]byte, Len)

	for i := 0; i < Len; i++ {
		codon := string(seq1[(i * 3):((i * 3) + 3)])
		if aminoAcid[codon] == 'B' {
			break
		}
		prot[i] = aminoAcid[codon]
	}

	fmt.Println(string(prot))

}
