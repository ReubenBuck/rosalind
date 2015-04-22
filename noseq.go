/*
Read in line by line and use a finite state machine to analyse the data
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/Users/labadmin/Downloads/rosalind_gc-7.txt")
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}
	defer f.Close()

	var (
		maxGC   = -1.0
		maxName string
		name    string
		seq     []byte
	)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := bytes.TrimSpace(sc.Bytes())
		if len(l) == 0 {
			continue
		}
		switch l[0] {
		case '>':
			if len(seq) != 0 {
				gc := gcOf(seq)
				if gc > maxGC {
					maxGC = gc
					maxName = name
				}
			}
			name = string(l[1:])
			seq = seq[:0]
		default:
			seq = append(seq, l...)
		}
	}
	if sc.Err() != nil {
		log.Fatalf("failed scanning: %v", err)
	}
	if maxGC == -1 {
		log.Println("no sequence")
		os.Exit(0)
	}

	gc := gcOf(seq)
	if gc > maxGC {
		maxGC = gc
		maxName = name
	}

	fmt.Printf("%s %v\n", maxName, maxGC*100)
}

func gcOf(s []byte) float64 {
	var a [128]byte
	for _, b := range s {
		a[b]++
	}
	return (float64(a['C']) + float64(a['G'])) / float64(len(s))
}
