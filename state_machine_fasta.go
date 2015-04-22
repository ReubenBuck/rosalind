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
	"sort"
)

type seq struct {
	name string
	seq  []byte
	gc   float64
}

type seqs []seq

func (s seqs) Len() int           { return len(s) }
func (s seqs) Less(i, j int) bool { return s[i].gc > s[j].gc }
func (s seqs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	f, err := os.Open("/Users/labadmin/Downloads/rosalind_gc-6.txt")
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}
	defer f.Close()

	var ss seqs
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := bytes.TrimSpace(sc.Bytes())
		if len(l) == 0 {
			continue
		}
		switch l[0] {
		case '>':
			if len(ss) != 0 {
				ss[len(ss)-1].gc = gc(ss[len(ss)-1].seq)
			}
			ss = append(ss, seq{name: string(l[1:])})
		default:
			ss[len(ss)-1].seq = append(ss[len(ss)-1].seq, l...)
		}
	}
	if sc.Err() != nil {
		log.Fatalf("failed scanning: %v", err)
	}
	if len(ss) == 0 {
		log.Println("no sequence")
		os.Exit(0)
	}

	ss[len(ss)-1].gc = gc(ss[len(ss)-1].seq)
	sort.Sort(ss)

	for i := 0; i < len(ss); i++ {
		fmt.Printf("%s %v\n", ss[i].name, ss[i].gc*100)
	}

}

func gc(s []byte) float64 {
	var a [128]byte
	for _, b := range s {
		a[b]++
	}
	return (float64(a['C']) + float64(a['G'])) / float64(len(s))
}
