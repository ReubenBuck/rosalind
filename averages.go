// averages
// we get six numbers and have to multiply them by 2 * a
// certain probability.
// after that we sum them

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	prob := []float64{2.0, 2.0, 2.0, 1.5, 1.0, 0.0}
	couples := make([]float64, 7)
	couples[0] = 0.0
	for i := 0; i < 6; i++ {
		a := os.Args[i+1]
		a1, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal("did not read data", err)
		}
		couples[i+1] = (float64(a1) * prob[i]) + couples[i]
	}

	fmt.Println(couples[6])

}
