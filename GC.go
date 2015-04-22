/*
 GC content
 A chance for me to use my newly found regexp stuff

so we want to read in all the headers and store them in a map along with their length and gc contnet

we need to read each time we see this ">"
and the sequence begins on the new line

*/

package main

import (
	"fmt"
	//	"github.com/biogo/biogo/io/seqio/fasta"
	"io/ioutil"
	//	"strings"
)

// this function works well
func gc_counter(s string) float64 {

	bs := []byte(s)
	var a [128]byte
	for i := 0; i < len(bs); i++ {
		a[bs[i]]++
	}
	gc_count := float64(a['G'] + a['C'])
	gc_per := gc_count / float64(len(bs))
	return gc_per

}

// need to code up a reader + writer

func main() {
	b_slice, err := ioutil.ReadFile("test4.fa")
	if err != nil {
		return
	}
	str := string(b_slice)
	fmt.Println(str)
	x := gc_counter(str)
	fmt.Println(x)

	fmt.Println(len(str))

	//fmt.Printf("Fields are: %q", strings.Fields(str))

}
