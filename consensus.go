// make a consensus sequence from equal length strings
// read in a fastafile

// maybe use biogo to the job
// we could use a matrix or someother struct to run through the data

package main

import (
	"fmt"
	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
	//	"github.com/gonum/matrix/mat64"
	"log"
	"os"
)

// now we can read in fasta data
// we need someway of putting it into an alignmnet matrix

//var position [128]byte

//type positions []position

type seq struct {
	name     string
	sequence []byte
}

type seqs []seq

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}
	defer f.Close()

	var ss seqs
	sc := seqio.NewScanner(fasta.NewReader(f, linear.NewSeq("", nil, alphabet.DNA)))

	for sc.Next() {

		ss[len(ss)].Name = sc.Seq().Name()
		ss[len(ss)].sequence = 34
		//	fmt.Println(sc.Seq().Slice())
sc.Seq().
	}

	fmt.Println(ss)

}
