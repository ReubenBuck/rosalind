// look at each open readign frame

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	// we want to take our input and find all the open reading frames
	bases := []byte{'T', 'C', 'A', 'G'}
	amino_acids := []byte("FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG")
	trans := translateTable(bases, amino_acids)

	// need to read in fasta format

	// this part needs work

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	defer f.Close()
	name := []byte{}
	seq1F := []byte{}
	sc := bufio.NewScanner(f)

	for sc.Scan() {

		// l takes each line of the file and runs it through our state machine

		l := bytes.TrimSpace(sc.Bytes())
		if len(l) == 0 {
			continue
		}
		switch l[0] {
		case '>':
			name = append(name, l...)
		default:
			seq1F = append(seq1F, l...)
		}
	}

	// need to do reverse complement
	seq1R := revC(seq1F)
	//	fmt.Println(string(seq1F))
	//	fmt.Println(string(seq1R))

	startF := startFind(seq1F)
	startR := startFind(seq1R)
	// now we need something to store all of our DNA strings in
	// a slice of slices to store each sequence and translate it
	// remeber to only return if there is a stop codon

	//	fmt.Println(startF)

	// use a for loop to set up all the strings to run throught
	// or run them through the translation and print as we go

	var proteins map[string]int
	proteins = make(map[string]int)

	for s := 0; s < len(startF); s++ {
		seq1 := seq1F[(startF[s] - 1):]
		//	fmt.Println(string(seq1))
		Len := len(seq1) / 3
		prot := make([]byte, Len)
		for i := 0; i < Len; i++ {
			codon := string(seq1[(i * 3):((i * 3) + 3)])
			if trans[codon] == '*' {
				//fmt.Println(string(prot[:i]))
				if proteins[string(prot[:i])] == 0 {
					proteins[string(prot[:i])] = 1
				}
				break
			}
			prot[i] = trans[codon]
		}
	}

	for s := 0; s < len(startR); s++ {
		seq1 := seq1R[(startR[s] - 1):]
		//	fmt.Println(string(seq1))
		Len := len(seq1) / 3
		prot := make([]byte, Len)
		for i := 0; i < Len; i++ {
			codon := string(seq1[(i * 3):((i * 3) + 3)])
			if trans[codon] == '*' {
				//fmt.Println(string(prot[:i]))
				if proteins[string(prot[:i])] == 0 {
					proteins[string(prot[:i])] = 1
				}
				break
			}
			prot[i] = trans[codon]
		}
	}

	// one final problem, stings have to be unique
	// maybe store each string as a name in a map object we can call
	// if we call it and

	// we have to ask if any new string is the name as any old string

	for key := range proteins {
		fmt.Println(key)
	}
}

// find all the start codons in the sting

func startFind(sequence []byte) []int {
	a := make([]int, 1)
	mot := []byte("ATG")
	for i := 0; i < (len(sequence) - (len(mot) - 1)); i++ {
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
