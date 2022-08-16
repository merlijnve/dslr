package main

import (
	"encoding/json"
	"fmt"
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
	Accuracy float64
	Mean0    float64
	Mean1    float64
	Std0     float64
	Std1     float64
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func writeThetas(c []Classifier) {
	file, err := json.MarshalIndent(c, "", " ")
	handleError(err, "Error: could not indent the theta values to json")
	err = os.WriteFile("./thetaValues.json", file, 0644)
	handleError(err, "Error: could not create theta values file")
}

func initClassifiers() []Classifier {
	classifiers := make([]Classifier, 0)

	classifiers = append(classifiers, Classifier{House: "Hufflepuff", Feature0: "Astronomy", Feature1: "Transfiguration"})
	classifiers = append(classifiers, Classifier{House: "Gryffindor", Feature0: "Charms", Feature1: "Flying"})
	classifiers = append(classifiers, Classifier{House: "Slytherin", Feature0: "Potions", Feature1: "Divination"})
	classifiers = append(classifiers, Classifier{House: "Ravenclaw", Feature0: "Charms", Feature1: "Muggle Studies"})

	return classifiers
}

func main() {
	dataset := readDataset()

	classifiers := initClassifiers()
	for i := range classifiers {
		fmt.Println("1. Making classifier for", classifiers[i].House, "using:\n", classifiers[i].Feature0, "-", classifiers[i].Feature1)
		classifiers[i].data = getDataPairs(dataset, classifiers[i])
		classifiers[i] = standardization(classifiers[i])
		classifiers[i] = gradientDescent(classifiers[i])
		fmt.Print("\n")
	}
	plotScatter(classifiers)
	writeThetas(classifiers)
}
