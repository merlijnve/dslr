package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const HUFFLEPUFF = 0
const GRYFFINDOR = 1
const SLYTHERIN = 2
const RAVENCLAW = 3

type Classifier struct {
	House    string
	Feature0 string
	Feature1 string
	data     [][]float64
	t0       float64
	t1       float64
	t2       float64
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func writeThetas(c []Classifier) {
	file, err := json.MarshalIndent(c, "", " ")
	handleError(err, "Error: could not create indent the theta values to json")

	err = ioutil.WriteFile("thetaValues.json", file, 0644)
	handleError(err, "Error: could not create theta values file")
}

func initClassifier() []Classifier {
	classifiers := make([]Classifier, 0)

	classifiers = append(classifiers, Classifier{House: "Hufflepuff", Feature0: "Astronomy", Feature1: "Transfiguration"})
	classifiers = append(classifiers, Classifier{House: "Gryffindor", Feature0: "Charms", Feature1: "Herbology"})
	classifiers = append(classifiers, Classifier{House: "Slytherin", Feature0: "History of Magic", Feature1: "Transfiguration"})
	classifiers = append(classifiers, Classifier{House: "Ravenclaw", Feature0: "Care of Magical Creatures", Feature1: "Charms"})

	return classifiers
}

func main() {
	dataset := readDataset()
	classifiers := initClassifier()

	// i := 0
	for i := range classifiers {
		c := &classifiers[i]
		fmt.Println("Making classifier for", c.House, "using:\n", c.Feature0, "vs", c.Feature1)
		c.data = getDataPair(dataset, *c)
		classifiers[i] = gradientDescent(*c)
	}
	plotScatter(classifiers)
	writeThetas(classifiers)
}
