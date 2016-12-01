// lets make a perceptron

// basicly something we can feed it features and it can train weights

// so feed it a file with tab seperated features

// train it on features and return the weights vector

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// read in data
// make a finite state machine
// we give it a csv file where the last column is the class of the individual

// we want to make an object the can read i

// now we have a function in which we feed it a file name of certain type and it reads it in for us

// unique ids varaible which sets up the targets
// it also checkes if there are more than two classifications

type feature struct {
	id     string
	target int
	values []float64
}

type features []feature

//var unique_id [2]string

func feature_reader(file string) features {

	var unique_id map[string]int
	unique_id = make(map[string]int)
	target_val := [3]int{1, -1}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}
	defer f.Close()

	var feats features
	var val []byte
	var feat feature

	idCount := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := sc.Bytes()
		for i := 0; i < len(l); i++ {
			switch l[i] {
			case ',':
				val_fl, err := strconv.ParseFloat(string(val), 64)
				if err != nil {
					log.Fatalf("parsing floats falied: %v", err)
				}
				feat.values = append(feat.values, val_fl)
				val = nil
				// write into a new section
			default:
				val = append(val, l[i])
			}
		}
		feat.id = string(val)

		feat.target = unique_id[feat.id]
		if feat.target == 0 {
			unique_id[feat.id] = target_val[idCount]
			feat.target = unique_id[feat.id]
			idCount++
			if idCount == 3 {
				log.Fatalf("more than two unique ids: %v", feat.id)
			}
		}

		feats = append(feats, feat)
		val = nil
		feat.values = nil

	}
	//
	return (feats)
}

// all the features are loaded onto

// we also need bias which acts like a weight but is not multiplied by x

// maybe also a targets vector

func main() {

	data1 := feature_reader(os.Args[1])
	var data features

	data_order := rand.Perm(len(data1))
	for i := 0; i < len(data_order); i++ {
		data = append(data, data1[data_order[i]])
	}

	learnRate := 0.000000000001

	weights := make([]float64, len(data[0].values))
	//	weightsUpdate := weights
	bias := 0.0

	var total float64
	var output int

	for i := 0; i < len(data); i++ {
		total = 0.0

		//fmt.Println(weights)

		// multiplying our numbers to gether and getting the sum
		for j := 0; j < len(data[i].values); j++ {
			total = total + (data[i].values[j] * weights[j])
		}

		total = (bias * -1) + total

		if total > 0 {
			output = 1
		} else {
			output = -1
		}

		// here we update the weeights

		for j := 0; j < len(data[i].values); j++ {
			weights[j] = weights[j] + (learnRate * float64(data[i].target-output) * data[i].values[j])
		}
		bias = bias + (learnRate * float64(data[i].target-output))

		fmt.Println(bias, weights, data[i].target, output)

	}

	// sum of x by weigths greter than 0 they get 1
	// sum of x by weigths less than 0 thew get -1

	//	for i := 0; i < len(data); i++ {
	//		fmt.Println(data[i])
	//	}

	// maybe i should permute the order
	// things may come out better then

	//fmt.Println(data)
}
