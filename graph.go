// Graph intro

// connecting strings of DNA
// do alignment with a matrix
// maybe read in fasta sequences using our struct

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"github.com/gonum/blas/blas64"
//	"github.com/gonum/matrix/mat64"
)

type seq struct {
	name string
	seq  []byte
}

type seqs []seq

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("reading file %v", err)
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

	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i].name, string(ss[i].seq))
	}

	// Now theres some room to build alignment matricies
	// if i can get them to work :(
	r1 := blas64.General{
		Rows: 2,
		Cols: 2,
		Data: []float64{},
	}
	mat64.





	fmt.Println(r1)

}
