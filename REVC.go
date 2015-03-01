// reverse complement
// Do what i did before with replace but see if
// there is an easy way to reverse it

// revesing seems hard and will require me to make myown function

// run through the strung backward and rewrite it

package main

import (
	"fmt"
	"os"
	"strings"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	myString := os.Args[1]
	r := strings.NewReplacer("T", "A", "A", "T", "G", "C", "C", "G")
	thing := r.Replace(myString)
	thing1 := Reverse(thing)

	fmt.Println(thing1)
}
