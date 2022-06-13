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
	T0       float64
	T1       float64
	T2       float64
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func writeThetas(thetas []Classifier) {
	file, err := json.MarshalIndent(thetas, "", " ")
	handleError(err, "Error: could not create indent the theta values to json")

	err = ioutil.WriteFile("thetaValues.json", file, 0644)
	handleError(err, "Error: could not create theta values file")
}

func initThetaHouseAndFeatures() []Classifier {
	thetas := make([]Classifier, 0)

	thetas = append(thetas, Classifier{House: "Hufflepuff", Feature0: "Flying", Feature1: "Herbology"})
	thetas = append(thetas, Classifier{House: "Gryffindor", Feature0: "Muggle Studies", Feature1: "Arithmancy"})
	thetas = append(thetas, Classifier{House: "Slytherin", Feature0: "History of Magic", Feature1: "Transfiguration"})
	thetas = append(thetas, Classifier{House: "Ravenclaw", Feature0: "Care of Magical Creatures", Feature1: "Charms"})

	return thetas
}

func main() {
	dataset := readDataset()
	classifiers := initThetaHouseAndFeatures()

	for i := range classifiers {
		c := &classifiers[i]
		fmt.Println("Making classifier for", c.House, "using:\n", c.Feature0, "vs", c.Feature1)
		c.data = getDataPair(dataset, *c)
	}
	writeThetas(classifiers)
}
