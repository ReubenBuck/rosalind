package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type seq struct {
	name string
	seq  []byte
}

type profile struct {
	name string
	freq []int
}

var (
	p         [256]profile
	seqs      []seq
	named     []byte
	consensus []string
)

func bubbleSwap(list []byte) []byte {
	swapped := true
	var listSwap byte
	for i := 0; i < len(named)-1; i++ {
		swapped = false
		if list[i] > list[i+1] {
			listSwap = list[i]
			list[i] = list[i+1]
			list[i+1] = listSwap
			swapped = true
			i = -1
		}
		if swapped == false {
			continue
		}
	}
	return list
}

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}

	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := bytes.TrimSpace(sc.Bytes())
		if len(l) == 0 {
			continue
		}
		switch l[0] {
		case '>':
			seqs = append(seqs, seq{name: string(l[1:])})
		default:
			seqs[len(seqs)-1].seq = append(seqs[len(seqs)-1].seq, l...)
		}
	}

	for j := range seqs[0].seq {
		for i := range seqs {
			if len(p[seqs[i].seq[j]].freq) == 0 {
				p[seqs[i].seq[j]].freq = make([]int, len(seqs[0].seq))
				p[seqs[i].seq[j]].name = string(seqs[i].seq[j])
				named = append(named, seqs[i].seq[j])
			}

			p[seqs[i].seq[j]].freq[j] = p[seqs[i].seq[j]].freq[j] + 1

		}
		max := p[named[0]]
		for i := 0; i < len(named); i++ {
			if max.freq[j] < p[named[i]].freq[j] {
				max = p[named[i]]
			}
		}
		consensus = append(consensus, max.name)

	}

	for j := 0; j < len(seqs[0].seq); j++ {
		fmt.Print(consensus[j])
	}
	fmt.Print("\n")

	named = bubbleSwap(named)
	for i := 0; i < len(named); i++ {
		fmt.Print(p[named[i]].name, ": ")
		for j := 0; j < len(seqs[0].seq); j++ {
			fmt.Print(p[named[i]].freq[j], " ")
		}
		fmt.Print("\n")
	}

}
