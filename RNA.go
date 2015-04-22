// turning DNA into RNA

package main

import (
	"fmt"
	"os"
)

func main() {

	s := []byte(os.Args[1])

	for i, v := range s {
		if v == 'T' {
			s[i] = 'U'
		}
	}

	fmt.Printf("%s\n", s)
}
