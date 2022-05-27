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

type Thetas struct {
	House    string
	Feature0 string
	Feature1 string
	T0       float64
	T1       float64
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func writeThetas(thetas []Thetas) {
	file, err := json.MarshalIndent(thetas, "", " ")
	handleError(err, "Error: could not create indent the theta values to json")

	err = ioutil.WriteFile("thetaValues.json", file, 0644)
	handleError(err, "Error: could not create theta values file")
}

func initThetaHouseAndFeatures() []Thetas {
	thetas := make([]Thetas, 0)

	thetas = append(thetas, Thetas{House: "Hufflepuff", Feature0: "Flying", Feature1: "Herbology"})
	thetas = append(thetas, Thetas{House: "Gryffindor", Feature0: "Muggle Studies", Feature1: "Arithmancy"})
	thetas = append(thetas, Thetas{House: "Slytherin", Feature0: "History of Magic", Feature1: "Transfiguration"})
	thetas = append(thetas, Thetas{House: "Ravenclaw", Feature0: "Care of Magical Creatures", Feature1: "Charms"})

	return thetas
}

func main() {
	dataset := readDataset()
	thetas := initThetaHouseAndFeatures()

	for _, t := range thetas {
		fmt.Println("Making classifier for", t.House, "using:\n", t.Feature0, "vs", t.Feature1)
		data := getDataPair(dataset, t)
		fmt.Println(data)
	}
	writeThetas(thetas)
}
